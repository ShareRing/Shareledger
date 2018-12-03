package app

import (
	"fmt"

	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "bitbucket.org/shareringvn/cosmos-sdk/baseapp"
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/constants"

	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
)

type TestShareLedgerApp struct {
	*bapp.BaseApp
	cdc *wire.Codec

	// keys to access the substores
	// bankKey *sdk.KVStoreKey
	//accountKey *sdk.KVStoreKey

	//keepers
	bankKeeper bank.Keeper

	// Manage getting and setting accounts
	accountMapper auth.AccountMapper
}

func NewTestShareLedgerApp(logger log.Logger, db dbm.DB) *TestShareLedgerApp {

	cdc := MakeCodec()

	// Create the base application object.
	baseApp := bapp.NewBaseApp(appName, cdc, logger, db)

	authKey := sdk.NewKVStoreKey(constants.STORE_AUTH)
	//bankKey := sdk.NewKVStoreKey(constants.STORE_BANK)

	// Mount Store

	baseApp.MountStoresIAVL(authKey)
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

	app := &TestShareLedgerApp{
		BaseApp: baseApp,
		cdc:     cdc,
		//accountKey:    accountKey,
		accountMapper: accountMapper,
	}

	app.SetTxDecoder(auth.GetTxDecoder(cdc))
	app.SetAnteHandler(auth.NewAnteHandler(accountMapper))
	app.Router().
		AddRoute(constants.MESSAGE_AUTH, auth.NewHandler(accountMapper))
	app.cdc = auth.RegisterCodec(app.cdc)

	app.SetupBank(accountMapper)

	// Set Tx Fee Calculation
	// app.SetFeeHandler(fee.NewFeeHandler(accountMapper, exchangeKey))

	// Register InitChain
	// logger.Info("Register Init Chainer")
	// app.SetInitChainer(app.InitChainer)
	// app.SetEndBlocker(EndBlocker(accountMapper, app.posKeeper))
	// app.SetBeginBlocker(BeginBlocker)

	return app
}

func (app *TestShareLedgerApp) SetupBank(am auth.AccountMapper) {
	// Bank module
	// Create a key for accessing the account store.
	app.cdc = bank.RegisterCodec(app.cdc)
	app.bankKeeper = bank.NewKeeper(am /*, cdc*/)
	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(am))
	app.Router().
		AddRoute("test", GetHandler())

}

func GetHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		fmt.Printf("Handler for Msg: %v\n", msg)
		return sdk.Result{
			Log: "Testing",
		}
	}
}
