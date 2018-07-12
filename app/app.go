package app

import (
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/bank"
	//"github.com/sharering/shareledger/x/bank/handlers"
)

const (
	ShareLedgerApp = "ShareLedger_v0.0.1"
)

func NewShareLedgerApp(logger log.Logger, db dbm.DB) *bapp.BaseApp {

	cdc := bank.MakeCodec()

	// Create the base application object.
	app := bapp.NewBaseApp(ShareLedgerApp, cdc, logger, db)

	// Create a key for accessing the account store.
	keyAccount := sdk.NewKVStoreKey("acc")

	// Determine how transactions are decoded.
	app.SetTxDecoder(bank.TxDecoder)

	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(keyAccount))
		//AddRoute("send", handlers.HandleMsgSend(keyAccount)).
		//AddRoute("check", handlers.HandleMsgCheck(keyAccount)).
		//AddRoute("load", handlers.HandleMsgLoad(keyAccount))

	// Mount stores and load the latest state.
	app.MountStoresIAVL(keyAccount)
	err := app.LoadLatestVersion(keyAccount)
	if err != nil {
		cmn.Exit(err.Error())
	}
	return app
}




