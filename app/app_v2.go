//go:build app_v2

package app

import (
	_ "embed"
	"io"
	"os"
	"path/filepath"

	"cosmossdk.io/depinject"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	groupmodule "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"
	nftmodule "github.com/cosmos/cosmos-sdk/x/nft/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	ibcfeekeeper "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"

	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	assetmodule "github.com/sharering/shareledger/x/asset"
	assetmodulekeeper "github.com/sharering/shareledger/x/asset/keeper"
	bookingmodule "github.com/sharering/shareledger/x/booking"
	bookingmodulekeeper "github.com/sharering/shareledger/x/booking/keeper"
	"github.com/sharering/shareledger/x/distributionx"
	distributionxkeeper "github.com/sharering/shareledger/x/distributionx/keeper"
	documentmodule "github.com/sharering/shareledger/x/document"
	documentmodulekeeper "github.com/sharering/shareledger/x/document/keeper"
	electoralmodule "github.com/sharering/shareledger/x/electoral"
	electoralmodulekeeper "github.com/sharering/shareledger/x/electoral/keeper"
	gentlemintmodule "github.com/sharering/shareledger/x/gentlemint"
	gentlemintmodulekeeper "github.com/sharering/shareledger/x/gentlemint/keeper"
	idmodule "github.com/sharering/shareledger/x/id"
	idmodulekeeper "github.com/sharering/shareledger/x/id/keeper"
	swapmodule "github.com/sharering/shareledger/x/swap"
	swapmodulekeeper "github.com/sharering/shareledger/x/swap/keeper"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		groupmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		nftmodule.AppModuleBasic{},
		consensus.AppModuleBasic{},

		// shareledger modules
		documentmodule.AppModuleBasic{},
		idmodule.AppModuleBasic{},
		assetmodule.AppModuleBasic{},
		bookingmodule.AppModuleBasic{},
		gentlemintmodule.AppModuleBasic{},
		electoralmodule.AppModuleBasic{},
		swapmodule.AppModuleBasic{},
		distributionx.AppModuleBasic{},

		// extends modules
		wasm.AppModuleBasic{},
		ica.AppModuleBasic{},
	)
)

var (
	_ runtime.AppI            = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             *govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	GroupKeeper           groupkeeper.Keeper
	NFTKeeper             nftkeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper

	IBCKeeper           *ibckeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	WasmKeeper          wasm.Keeper

	// shareledger keeper
	DocumentKeeper      documentmodulekeeper.Keeper
	IDKeeper            idmodulekeeper.Keeper
	AssetKeeper         assetmodulekeeper.Keeper
	BookingKeeper       bookingmodulekeeper.Keeper
	GentleMintKeeper    gentlemintmodulekeeper.Keeper
	ElectoralKeeper     electoralmodulekeeper.Keeper
	SwapKeeper          swapmodulekeeper.Keeper
	ContractKeeper      *wasmkeeper.PermissionedKeeper
	DistributionxKeeper distributionxkeeper.Keeper

	// simulation manager
	sm *module.SimulationManager
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".Shareledger")
}

func New(logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	var (
		shareledgerApp = &App{}
		appBuilder     *runtime.AppBuilder
		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(AppConfig, depinject.Supply(appOpts))
	)

	if err := depinject.Inject(appConfig,
		&appBuilder,
		&shareledgerApp.appCodec,
		&shareledgerApp.legacyAmino,
		&shareledgerApp.txConfig,
		&shareledgerApp.interfaceRegistry,
		&shareledgerApp.AccountKeeper,
		&shareledgerApp.BankKeeper,
		&shareledgerApp.CapabilityKeeper,
		&shareledgerApp.StakingKeeper,
		&shareledgerApp.SlashingKeeper,
		&shareledgerApp.MintKeeper,
		&shareledgerApp.DistrKeeper,
		&shareledgerApp.GovKeeper,
		&shareledgerApp.CrisisKeeper,
		&shareledgerApp.UpgradeKeeper,
		&shareledgerApp.ParamsKeeper,
		&shareledgerApp.AuthzKeeper,
		&shareledgerApp.EvidenceKeeper,
		&shareledgerApp.FeeGrantKeeper,
		&shareledgerApp.GroupKeeper,
		&shareledgerApp.NFTKeeper,
		&shareledgerApp.ConsensusParamsKeeper,

		// &shareledgerApp.IBCKeeper,
		// &shareledgerApp.IBCFeeKeeper,
		// &shareledgerApp.ICAControllerKeeper,
		// &shareledgerApp.ICAHostKeeper,
		// &shareledgerApp.TransferKeeper,
		// &shareledgerApp.WasmKeeper,
		// &shareledgerApp.ContractKeeper,

		// shareledger keeper
		&shareledgerApp.DocumentKeeper,
		&shareledgerApp.IDKeeper,
		&shareledgerApp.AssetKeeper,
		&shareledgerApp.BookingKeeper,
		&shareledgerApp.GentleMintKeeper,
		&shareledgerApp.ElectoralKeeper,
		&shareledgerApp.SwapKeeper,
		&shareledgerApp.DistributionxKeeper,
	); err != nil {
		panic(err)
	}

	shareledgerApp.App = appBuilder.Build(logger, db, traceStore, baseAppOptions...)

	if _, _, err := streaming.LoadStreamingServices(shareledgerApp.App.BaseApp, appOpts, shareledgerApp.appCodec, logger, shareledgerApp.kvStoreKeys()); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	shareledgerApp.ModuleManager.RegisterInvariants(shareledgerApp.CrisisKeeper)
	// TODO: handle upgrade here

	if err := shareledgerApp.Load(loadLatest); err != nil {
		panic(err)
	}

	return shareledgerApp
}

func (a *App) Name() string {
	return a.App.Name()
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.legacyAmino
}

func (a *App) AppCodec() codec.Codec {
	return a.appCodec
}

func (a *App) InterfaceRegistry() codectypes.InterfaceRegistry {
	return a.interfaceRegistry
}

func (a *App) TxConfig() client.TxConfig {
	return a.txConfig
}

// NOTE: This is solely to be used for testing purposes.
func (a *App) GetKey(storeKey string) *storetypes.KVStoreKey {
	sk := a.UnsafeFindStoreKey(storeKey)
	kvStoreKey, ok := sk.(*storetypes.KVStoreKey)
	if !ok {
		return nil
	}
	return kvStoreKey
}

func (a *App) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range a.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}
	return keys
}

// NOTE: This is solely to be used for testing purposes.
func (a *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := a.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (a *App) SimulationManager() *module.SimulationManager {
	return a.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (a *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	a.App.RegisterAPIRoutes(apiSvr, apiConfig)
}

// NOTE: This is solely to be used for testing purposes.
func GetMaccPerms() map[string][]string {
	dup := make(map[string][]string)
	for _, perms := range moduleAccPerms {
		dup[perms.Account] = perms.Permissions
	}

	return dup
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)

	if len(blockAccAddrs) > 0 {
		for _, addr := range blockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}

	return result
}

// GetDefaultBypassFeeMessages used by globalfees
func GetDefaultBypassFeeMessages() []string {
	return []string{
		sdk.MsgTypeURL(&ibcchanneltypes.MsgRecvPacket{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgAcknowledgement{}),
		sdk.MsgTypeURL(&ibcclienttypes.MsgUpdateClient{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgTimeout{}),
		sdk.MsgTypeURL(&ibcchanneltypes.MsgTimeoutOnClose{}),
	}
}
