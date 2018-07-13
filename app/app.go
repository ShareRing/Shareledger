package app

import (
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/asset/handlers"
)

const (
	ShareLedgerApp = "ShareLedger_v0.0.1"
)

func NewShareLedgerApp(logger log.Logger, db dbm.DB) *bapp.BaseApp {

	cdc := MakeCodec()

	// Create the base application object.
	app := bapp.NewBaseApp(ShareLedgerApp, cdc, logger, db)


	SetupAsset(app, cdc)
	SetupBank(app, cdc)

	// Determine how transactions are decoded.
	app.SetTxDecoder(types.GetTxDecoder(cdc))

	return app
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	return cdc
}

func SetupBank(app *bapp.BaseApp, cdc *wire.Codec) {
	// Bank module
	// Create a key for accessing the account store.
	keyAccount := sdk.NewKVStoreKey("acc")
	cdc = bank.RegisterCodec(cdc)

	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(keyAccount))

	// Mount stores and load the latest state.
	app.MountStoresIAVL(keyAccount)
	err := app.LoadLatestVersion(keyAccount)
	if err != nil {
		cmn.Exit(err.Error())
	}
}

func SetupAsset(app *bapp.BaseApp, cdc *wire.Codec) {
	keyAsset := sdk.NewKVStoreKey("asset")

	keeper := asset.NewKeeper(keyAsset, cdc)

	cdc = asset.RegisterCodec(cdc)


	app.Router().
		AddRoute("asset", handlers.NewHandler(keeper))

	app.MountStoresIAVL(keyAsset)
	err := app.LoadLatestVersion(keyAsset)
	if err != nil {
		cmn.Exit(err.Error())
	}
}
