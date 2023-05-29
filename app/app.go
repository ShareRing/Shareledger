//go:build !app_v2

package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	tmjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/spf13/cast"

	"github.com/sharering/shareledger/ante"
	"github.com/sharering/shareledger/app/keepers"
	"github.com/sharering/shareledger/app/params"
	"github.com/sharering/shareledger/app/upgrades"
	v2 "github.com/sharering/shareledger/app/upgrades/v2"
)

const (
	AccountAddressPrefix = "shareledger"
	Name                 = "Shareledger"
)

var (
	_ servertypes.Application = &App{}

	FlagAppOptionSkipCheckVoter = "skip-checking-voter-role"

	DefaultNodeHome string
	Upgrades        = []upgrades.Upgrade{v2.Upgrade}
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp
	keepers.AppKeepers

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	invCheckPeriod    uint

	// the module manager
	ModuleManager *module.Manager
	// simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// New returns a reference to an initialized blockchain app
func New(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig params.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(encodingConfig.TxConfig.TxEncoder())

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
	}

	// Setup keepers
	app.AppKeepers = keepers.NewAppKeeper(
		appCodec,
		bApp,
		encodingConfig.Amino,
		maccPerms,
		app.BlockedModuleAccountAddrs(),
		skipUpgradeHeights,
		homePath,
		invCheckPeriod,
		appOpts,
	)

	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
	app.ModuleManager = module.NewManager(appModules(app, encodingConfig, skipGenesisInvariants)...)

	app.ModuleManager.SetOrderBeginBlockers(orderBeginBlockers()...)
	app.ModuleManager.SetOrderEndBlockers(orderEndBlockers()...)
	app.ModuleManager.SetOrderInitGenesis(orderInitBlockers()...)

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	// app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.ModuleManager.RegisterServices(app.configurator)

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules))

	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(simulationModules(app, encodingConfig, skipGenesisInvariants)...)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(app.AppKeepers.GetKVStoreKey())
	app.MountTransientStores(app.AppKeepers.GetTransientStoreKey())
	app.MountMemoryStores(app.AppKeepers.GetMemoryStoreKey())

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	var roleKeeper ante.RoleKeeper
	roleKeeper = app.ElectoralKeeper

	isSkipCheckVoter := cast.ToBool(appOpts.Get(FlagAppOptionSkipCheckVoter))
	if isSkipCheckVoter {
		or := ante.RoleKeeperWithoutVoter{
			RoleKeeper: roleKeeper,
		}
		roleKeeper = or
	}

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	var bypassMinFeeMsgTypes []string
	bypassMinFeeMsgTypesOptions := appOpts.Get(params.BypassMinFeeMsgTypesKey)
	if bypassMinFeeMsgTypesOptions == nil {
		bypassMinFeeMsgTypes = GetDefaultBypassFeeMessages()
	} else {
		bypassMinFeeMsgTypes = cast.ToStringSlice(bypassMinFeeMsgTypesOptions)
	}

	handlerOpts := ante.HandlerOptions{
		HandlerOptions:       authante.HandlerOptions{AccountKeeper: app.AccountKeeper, BankKeeper: app.BankKeeper, FeegrantKeeper: app.FeeGrantKeeper, SignModeHandler: encodingConfig.TxConfig.SignModeHandler(), SigGasConsumer: authante.DefaultSigVerificationGasConsumer},
		IBCKeeper:            app.IBCKeeper,
		WasmConfig:           &wasmConfig,
		TXCounterStoreKey:    app.AppKeepers.GetKey(wasm.StoreKey),
		GentlemintKeeper:     app.GentleMintKeeper,
		RoleKeeper:           roleKeeper,
		IDKeeper:             app.IDKeeper,
		DistributionxKeeper:  app.DistributionxKeeper,
		WasmKeeper:           app.WasmKeeper,
		BypassMinFeeMsgTypes: bypassMinFeeMsgTypes,
		GlobalFeeSubspace:    app.GetSubspace(gentleminttypes.ModuleName),
	}
	anteHandler, err := ante.NewHandler(handlerOpts)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %w", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.WasmKeeper))
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %w", err))
		}
	}

	app.setupUpgradeHandlers()
	app.setupUpgradeStoreLoaders()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// RegisterNodeService implements types.Application
func (app *App) RegisterNodeService(clientCtx client.Context) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.ModuleManager.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.ModuleManager.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// BlockedModuleAccountAddrs returns all the app's blocked module account
// addresses.
func (app *App) BlockedModuleAccountAddrs() map[string]bool {
	modAccAddrs := app.ModuleAccountAddrs()

	// remove module accounts that are ALLOWED to received funds
	//
	// TODO: Blocked on updating to v0.46.x
	// delete(modAccAddrs, authtypes.NewModuleAddress(grouptypes.ModuleName).String())
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Codec {
	return app.appCodec
}

// SimulationManager implements the SimulationApp interface
func (app *App) SimulationManager() *module.SimulationManager {
	return app.sm
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

func GetDefaultBypassFeeMessages() []string {
	return []string{
		sdk.MsgTypeURL(&ibcchanneltypes.MsgRecvPacket{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgAcknowledgement{}),
		sdk.MsgTypeURL(&ibcclienttypes.MsgUpdateClient{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgTimeout{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgTimeoutOnClose{}),
	}
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *App) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
		}
	}
}

func (app *App) setupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(
				app.ModuleManager,
				app.configurator,
				&app.AppKeepers,
			),
		)
	}
}

// GetMM get module manager
func (app *App) GetMM() *module.Manager {
	return app.ModuleManager
}
