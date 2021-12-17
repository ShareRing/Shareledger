package simapp

import (
	shareApp "github.com/ShareRing/Shareledger/app"
	"github.com/ShareRing/Shareledger/docs"
	assetmodule "github.com/ShareRing/Shareledger/x/asset"
	assetmodulekeeper "github.com/ShareRing/Shareledger/x/asset/keeper"
	assetmoduletypes "github.com/ShareRing/Shareledger/x/asset/types"
	bookingmodule "github.com/ShareRing/Shareledger/x/booking"
	bookingmodulekeeper "github.com/ShareRing/Shareledger/x/booking/keeper"
	bookingmoduletypes "github.com/ShareRing/Shareledger/x/booking/types"
	documentmodule "github.com/ShareRing/Shareledger/x/document"
	documentmodulekeeper "github.com/ShareRing/Shareledger/x/document/keeper"
	documentmoduletypes "github.com/ShareRing/Shareledger/x/document/types"
	"github.com/ShareRing/Shareledger/x/electoral"
	gentlemintmodule "github.com/ShareRing/Shareledger/x/gentlemint"
	gentlemintmodulekeeper "github.com/ShareRing/Shareledger/x/gentlemint/keeper"
	gentlemintmoduletypes "github.com/ShareRing/Shareledger/x/gentlemint/types"
	idmodule "github.com/ShareRing/Shareledger/x/id"
	idmodulekeeper "github.com/ShareRing/Shareledger/x/id/keeper"
	idmoduletypes "github.com/ShareRing/Shareledger/x/id/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/modules/core"
	ibcclient "github.com/cosmos/ibc-go/modules/core/02-client"
	ibcporttypes "github.com/cosmos/ibc-go/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/modules/core/keeper"
	"github.com/spf13/cast"
	"github.com/tendermint/spm/openapiconsole"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/ShareRing/Shareledger/app"
)

const appName = "SimApp"

type SimApp struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	DocumentKeeper documentmodulekeeper.Keeper

	IdKeeper idmodulekeeper.Keeper

	AssetKeeper assetmodulekeeper.Keeper

	BookingKeeper bookingmodulekeeper.Keeper

	GentleMintKeeper gentlemintmodulekeeper.Keeper


	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	// the module manager
	mm *module.Manager
}

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler
	// this line is used by starport scaffolding # stargate/app/govProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		// this line is used by starport scaffolding # stargate/app/govProposalHandler
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()...),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		documentmodule.AppModuleBasic{},
		idmodule.AppModuleBasic{},
		assetmodule.AppModuleBasic{},
		bookingmodule.AppModuleBasic{},
		gentlemintmodule.AppModuleBasic{},
		electoral.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:       nil,
		distrtypes.ModuleName:            nil,
		minttypes.ModuleName:             {authtypes.Minter},
		stakingtypes.BondedPoolName:      {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:   {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:              {authtypes.Burner},
		ibctransfertypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		bookingmoduletypes.ModuleName:    nil,
		gentlemintmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}
)

var (
	_ servertypes.Application = (*SimApp)(nil)
)

// New creates application instance with in-memory database and disabled logging.
func New(dir string) SimApp {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	a := NewSimApp(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
		simapp.EmptyAppOptions{})
	// InitChain updates deliverState which is required when app.NewContext is called
	a.InitChain(abci.RequestInitChain{
		ConsensusParams: defaultConsensusParams,
		AppStateBytes:   []byte("{}"),
	})
	return a
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+appName)
}

//NewSimApp returns a reference to an initialized Gaia.
func NewSimApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) SimApp {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		documentmoduletypes.StoreKey,
		idmoduletypes.StoreKey,
		assetmoduletypes.StoreKey,
		bookingmoduletypes.StoreKey,
		gentlemintmoduletypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey, "testingkey")

	sApp := SimApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	sApp.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(sApp.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	sApp.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := sApp.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := sApp.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	sApp.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], sApp.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	sApp.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], sApp.AccountKeeper, sApp.GetSubspace(banktypes.ModuleName), sApp.ModuleAccountAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], sApp.AccountKeeper, sApp.BankKeeper, sApp.GetSubspace(stakingtypes.ModuleName),
	)
	sApp.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], sApp.GetSubspace(minttypes.ModuleName), &stakingKeeper,
		sApp.AccountKeeper, sApp.BankKeeper, authtypes.FeeCollectorName,
	)
	sApp.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], sApp.GetSubspace(distrtypes.ModuleName), sApp.AccountKeeper, sApp.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, sApp.ModuleAccountAddrs(),
	)
	sApp.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, sApp.GetSubspace(slashingtypes.ModuleName),
	)
	sApp.CrisisKeeper = crisiskeeper.NewKeeper(
		sApp.GetSubspace(crisistypes.ModuleName), invCheckPeriod, sApp.BankKeeper, authtypes.FeeCollectorName,
	)

	sApp.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], sApp.AccountKeeper)
	sApp.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, sApp.BaseApp)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	sApp.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(sApp.DistrKeeper.Hooks(), sApp.SlashingKeeper.Hooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	sApp.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], sApp.GetSubspace(ibchost.ModuleName), sApp.StakingKeeper, sApp.UpgradeKeeper, scopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(sApp.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(sApp.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(sApp.UpgradeKeeper)).
		AddRoute(ibchost.RouterKey, ibcclient.NewClientProposalHandler(sApp.IBCKeeper.ClientKeeper))

	// Create Transfer Keepers
	sApp.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], sApp.GetSubspace(ibctransfertypes.ModuleName),
		sApp.IBCKeeper.ChannelKeeper, &sApp.IBCKeeper.PortKeeper,
		sApp.AccountKeeper, sApp.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(sApp.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &sApp.StakingKeeper, sApp.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	sApp.EvidenceKeeper = *evidenceKeeper

	sApp.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], sApp.GetSubspace(govtypes.ModuleName), sApp.AccountKeeper, sApp.BankKeeper,
		&stakingKeeper, govRouter,
	)

	sApp.DocumentKeeper = *documentmodulekeeper.NewKeeper(
		appCodec,
		keys[documentmoduletypes.StoreKey],
		keys[documentmoduletypes.MemStoreKey],
	)
	documentModule := documentmodule.NewAppModule(appCodec, sApp.DocumentKeeper)

	sApp.GentleMintKeeper = *gentlemintmodulekeeper.NewKeeper(
		appCodec,
		keys[gentlemintmoduletypes.StoreKey],
		keys[gentlemintmoduletypes.MemStoreKey],

		sApp.BankKeeper,
		sApp.AccountKeeper,
	)
	gentlemintModule := gentlemintmodule.NewAppModule(appCodec, sApp.GentleMintKeeper)

	sApp.IdKeeper = *idmodulekeeper.NewKeeper(
		appCodec,
		keys[idmoduletypes.StoreKey],
		keys[idmoduletypes.MemStoreKey],
		sApp.GentleMintKeeper,
	)
	idModule := idmodule.NewAppModule(appCodec, sApp.IdKeeper)

	sApp.AssetKeeper = *assetmodulekeeper.NewKeeper(
		appCodec,
		keys[assetmoduletypes.StoreKey],
		keys[assetmoduletypes.MemStoreKey],
	)
	assetModule := assetmodule.NewAppModule(appCodec, sApp.AssetKeeper)

	sApp.BookingKeeper = *bookingmodulekeeper.NewKeeper(
		appCodec,
		keys[bookingmoduletypes.StoreKey],
		keys[bookingmoduletypes.MemStoreKey],
		sApp.AssetKeeper,
		sApp.BankKeeper,
	)
	bookingModule := bookingmodule.NewAppModule(appCodec, sApp.BookingKeeper)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	// this line is used by starport scaffolding # ibc/app/router
	sApp.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	sApp.mm = module.NewManager(
		genutil.NewAppModule(
			sApp.AccountKeeper, sApp.StakingKeeper, sApp.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, sApp.AccountKeeper, nil),
		vesting.NewAppModule(sApp.AccountKeeper, sApp.BankKeeper),
		bank.NewAppModule(appCodec, sApp.BankKeeper, sApp.AccountKeeper),
		capability.NewAppModule(appCodec, *sApp.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, sApp.AccountKeeper, sApp.BankKeeper, sApp.FeeGrantKeeper, sApp.interfaceRegistry),
		crisis.NewAppModule(&sApp.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, sApp.GovKeeper, sApp.AccountKeeper, sApp.BankKeeper),
		mint.NewAppModule(appCodec, sApp.MintKeeper, sApp.AccountKeeper),
		slashing.NewAppModule(appCodec, sApp.SlashingKeeper, sApp.AccountKeeper, sApp.BankKeeper, sApp.StakingKeeper),
		distr.NewAppModule(appCodec, sApp.DistrKeeper, sApp.AccountKeeper, sApp.BankKeeper, sApp.StakingKeeper),
		staking.NewAppModule(appCodec, sApp.StakingKeeper, sApp.AccountKeeper, sApp.BankKeeper),
		upgrade.NewAppModule(sApp.UpgradeKeeper),
		evidence.NewAppModule(sApp.EvidenceKeeper),
		ibc.NewAppModule(sApp.IBCKeeper),
		params.NewAppModule(sApp.ParamsKeeper),
		transferModule,
		documentModule,
		idModule,
		assetModule,
		bookingModule,
		gentlemintModule,
		// this line is used by starport scaffolding # stargate/app/appModule

	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	sApp.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName, capabilitytypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
		evidencetypes.ModuleName, stakingtypes.ModuleName, ibchost.ModuleName,
		feegrant.ModuleName,
	)

	sApp.mm.SetOrderEndBlockers(crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	sApp.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		gentlemintmoduletypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		documentmoduletypes.ModuleName,
		idmoduletypes.ModuleName,
		assetmoduletypes.ModuleName,
		bookingmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis

	)

	sApp.mm.RegisterInvariants(&sApp.CrisisKeeper)
	sApp.mm.RegisterRoutes(sApp.Router(), sApp.QueryRouter(), encodingConfig.Amino)
	sApp.mm.RegisterServices(module.NewConfigurator(sApp.appCodec, sApp.MsgServiceRouter(), sApp.GRPCQueryRouter()))

	// initialize stores
	sApp.MountKVStores(keys)
	sApp.MountTransientStores(tkeys)
	sApp.MountMemoryStores(memKeys)

	// initialize BaseApp
	sApp.SetInitChainer(sApp.InitChainer)
	sApp.SetBeginBlocker(sApp.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   sApp.AccountKeeper,
			BankKeeper:      sApp.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  sApp.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	sApp.SetAnteHandler(anteHandler)
	sApp.SetEndBlocker(sApp.EndBlocker)

	if loadLatest {
		if err := sApp.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	sApp.ScopedIBCKeeper = scopedIBCKeeper
	sApp.ScopedTransferKeeper = scopedTransferKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return sApp
}

// Name returns the name of the App
func (app SimApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState shareApp.GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimApp) ModuleAccountAddrs() map[string]bool {
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
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *SimApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register app's OpenAPI routes.
	apiSvr.Router.Handle("/static/openapi.yml", http.FileServer(http.FS(docs.Docs)))
	apiSvr.Router.HandleFunc("/", openapiconsole.Handler(appName, "/static/openapi.yml"))
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app SimApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(documentmoduletypes.ModuleName)
	paramsKeeper.Subspace(idmoduletypes.ModuleName)
	paramsKeeper.Subspace(assetmoduletypes.ModuleName)
	paramsKeeper.Subspace(bookingmoduletypes.ModuleName)
	paramsKeeper.Subspace(gentlemintmoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace

	return paramsKeeper
}
