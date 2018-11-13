package app

import (
	"os"

	abci "github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "bitbucket.org/shareringvn/cosmos-sdk/baseapp"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"

	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/booking"
	"github.com/sharering/shareledger/x/exchange"
	"github.com/sharering/shareledger/x/pos"
	pKeeper "github.com/sharering/shareledger/x/pos/keeper"
)

const (
	appName = "ShareLedger_v0.0.1"
)

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.shareledgercli")
)

type ShareLedgerApp struct {
	*bapp.BaseApp
	cdc *wire.Codec

	// keys to access the substores
	assetKey   *sdk.KVStoreKey
	bookingKey *sdk.KVStoreKey
	posKey     *sdk.KVStoreKey
	bankKey    *sdk.KVStoreKey
	//accountKey *sdk.KVStoreKey

	//keepers
	bankKeeper     bank.Keeper
	posKeeper      pKeeper.Keeper
	bookingKeeper  booking.Keeper
	assetKeeper    asset.Keeper
	exchangeKeeper exchange.Keeper

	// Manage getting and setting accounts
	accountMapper auth.AccountMapper
}

func NewShareLedgerApp(logger log.Logger, db dbm.DB) *ShareLedgerApp {

	cdc := MakeCodec()

	// Create the base application object.
	baseApp := bapp.NewBaseApp(appName, cdc, logger, db)

	assetKey := sdk.NewKVStoreKey(constants.STORE_ASSET)
	bookingKey := sdk.NewKVStoreKey(constants.STORE_BOOKING)
	//accountKey := sdk.NewKVStoreKey(constants.STORE_BANK)
	authKey := sdk.NewKVStoreKey(constants.STORE_AUTH)
	posKey := sdk.NewKVStoreKey(constants.STORE_POS)
	exchangeKey := sdk.NewKVStoreKey(constants.STORE_EXCHANGE)
	//bankKey := sdk.NewKVStoreKey(constants.STORE_BANK)

	// Mount Store

	baseApp.MountStoresIAVL(authKey, assetKey, bookingKey, posKey, exchangeKey)
	err := baseApp.LoadLatestVersion(authKey)
	if err != nil {
		cmn.Exit(err.Error())
	}

	// accountMapper for Auth Module storing and Bank module
	accountMapper := auth.NewAccountMapper(
		cdc,
		authKey,
		&auth.SHRAccount{},
	)

	// Determine how transactions are decoded.
	//baseApp.SetTxDecoder(types.GetTxDecoder(cdc))

	app := &ShareLedgerApp{
		BaseApp:    baseApp,
		cdc:        cdc,
		assetKey:   assetKey,
		bookingKey: bookingKey,
		posKey:     posKey,
		//accountKey:    accountKey,
		accountMapper: accountMapper,
	}
	app.SetupAsset(assetKey)
	app.SetupBank(accountMapper)
	app.SetupPOS(posKey, accountMapper)
	app.SetupBooking(bookingKey, assetKey, accountMapper)
	app.SetupExchange(exchangeKey)

	app.SetTxDecoder(auth.GetTxDecoder(cdc))
	app.SetAnteHandler(auth.NewAnteHandler(accountMapper))
	app.Router().
		AddRoute(constants.MESSAGE_AUTH, auth.NewHandler(accountMapper))
	app.cdc = auth.RegisterCodec(app.cdc)

	// Set Tx Fee Calculation
	//app.SetFeeHandler(fee.NewFeeHandler(accountMapper))

	// Register InitChain
	logger.Info("Register Init Chainer")
	app.SetInitChainer(app.InitChainer)
	app.SetEndBlocker(EndBlocker(accountMapper, app.posKeeper))
	app.SetBeginBlocker(BeginBlocker)

	return app
}

// InitChainer will set initial balances for accounts as well as initial coin metadata

func (app *ShareLedgerApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {

	stateJSON := req.AppStateBytes
	//fmt.Printf("RequestInitChain.Time: %v\n", req.Time)
	//fmt.Printf("RequestInitChain.ChainId: %v\n", req.ChainId)
	//fmt.Printf("RequestInitChain.ConsensusParams: %v\n", req.ConsensusParams)
	//fmt.Printf("RequestInitChain.Validators: %v\n", req.Validators)
	//fmt.Printf("RequestInitChain.AppStateBytes: %v\n", req.AppStateBytes)

	var genesisState GenesisState
	// fmt.Printf("stateJSON=%s\n", stateJSON)

	err := app.cdc.UnmarshalJSON(stateJSON, &genesisState)
	// fmt.Printf("req=%v\n", genesisState)

	if err != nil {
		panic(err)
	}

	// load the accounts
	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToSHRAccount()
		app.accountMapper.SetAccount(ctx, acc)
	}

	// load the initial POS information
	abciVals, err := pos.InitGenesis(ctx, app.posKeeper, genesisState.StakeData)
	if err != nil {
		panic(err)
	}

	return abci.ResponseInitChain{
		Validators: abciVals, //use the validator defined in stake
	}

}

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) (res abci.ResponseBeginBlock) {

	// Save BlockHeader and Height to Context
	ctx.WithBlockHeader(req.Header).WithBlockHeight(req.Header.Height)

	//fmt.Printf("BeginBlocker: %v\n", req.Header.Proposer)

	return
}

// application updates every end block
func EndBlocker(am auth.AccountMapper, keeper pKeeper.Keeper) sdk.EndBlocker {
	return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

		proposer := ctx.BlockHeader().Proposer

		//	fmt.Printf("Proposer: %v\n", proposer)
		//	fmt.Printf("Proposer PubKey: %v\n", proposer.PubKey)

		var pubKey types.PubKeySecp256k1

		if len(proposer.PubKey.GetData()) > 1 {
			pubKey = types.ConvertToPubKey(proposer.PubKey.GetData())
			// fmt.Printf("Address: %s\n", pubKey.Address())
		} else {
			pubKey = types.NilPubKeySecp256k1()
		}

		validatorUpdates := pos.EndBlocker(ctx, keeper, pubKey)
		// Add these new validators to the addr -> pubkey map.
		return abci.ResponseEndBlock{
			ValidatorUpdates: validatorUpdates,
		}

		// Add these new validators to the addr -> pubkey map.
		return abci.ResponseEndBlock{}
	}
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*types.SHRTx)(nil), nil)
	cdc.RegisterConcrete(types.BasicTx{}, "shareledger/BasicTx", nil)
	cdc.RegisterConcrete(auth.AuthTx{}, "shareledger/AuthTx", nil)
	cdc.RegisterConcrete(types.QueryTx{}, "shareledger/QueryTx", nil)

	cdc.RegisterInterface((*types.SHRSignature)(nil), nil)
	cdc.RegisterConcrete(types.BasicSig{}, "shareledger/BasicSig", nil)
	cdc.RegisterConcrete(auth.AuthSig{}, "shareledger/AuthSig", nil)

	cdc.RegisterInterface((*auth.BaseAccount)(nil), nil)
	cdc.RegisterConcrete(auth.SHRAccount{}, "shareledger/SHRAccount", nil)

	cdc.RegisterInterface((*types.PubKey)(nil), nil)
	cdc.RegisterConcrete(types.PubKeySecp256k1{}, "shareledger/PubSecp256k1", nil)

	cdc.RegisterInterface((*types.Signature)(nil), nil)
	cdc.RegisterConcrete(types.SignatureSecp256k1{}, "shareledger/SigSecp256k1", nil)

	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	return cdc
}

func (app *ShareLedgerApp) SetupBank(am auth.AccountMapper) {
	// Bank module
	// Create a key for accessing the account store.
	app.cdc = bank.RegisterCodec(app.cdc)
	app.bankKeeper = bank.NewKeeper(am /*, cdc*/)
	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(am))

}

func (app *ShareLedgerApp) SetupAsset(assetKey *sdk.KVStoreKey) {

	keeper := asset.NewKeeper(assetKey, app.cdc)

	app.cdc = asset.RegisterCodec(app.cdc)

	app.Router().
		AddRoute("asset", asset.NewHandler(keeper))

	// app.MountStoresIAVL(assetKey)
}

func (app *ShareLedgerApp) SetupBooking(bookingKey *sdk.KVStoreKey,
	assetKey *sdk.KVStoreKey, am auth.AccountMapper) {

	app.cdc = booking.RegisterCodec(app.cdc)

	app.bookingKeeper = booking.NewKeeper(bookingKey,
		assetKey,
		am,
		app.cdc)

	app.Router().
		AddRoute("booking", booking.NewHandler(app.bookingKeeper))

	// app.MountStoresIAVL(bookingKey)

}

func (app *ShareLedgerApp) SetupPOS(posKey *sdk.KVStoreKey,
	am auth.AccountMapper) {
	app.cdc = pos.RegisterCodec(app.cdc)
	app.posKeeper = pKeeper.NewKeeper(posKey, app.bankKeeper, app.cdc)
	app.Router().AddRoute("pos", pos.NewHandler(app.posKeeper))
	app.QueryRouter().
		AddRoute("pos", pos.NewQuerier(app.posKeeper, app.cdc))

	//return k

}

func (app *ShareLedgerApp) SetupExchange(exchangeKey *sdk.KVStoreKey) {
	app.cdc = exchange.RegisterCodec(app.cdc)
	app.exchangeKeeper = exchange.NewKeeper(exchangeKey, app.cdc)
	app.Router().AddRoute("exchangerate", exchange.NewHandler(app.exchangeKeeper))
}
