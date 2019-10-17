package baseapp

import (
	bapp "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	log "github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
	fee "github.com/sharering/shareledger/x/fee"
)

// BaseApp - wrapper around BaseApp
type BaseApp struct {
	*bapp.BaseApp
	feeHandler fee.FeeHandler
}

func NewBaseApp(
	name string, logger log.Logger, db dbm.DB, txDecoder sdk.TxDecoder, options ...func(*bapp.BaseApp),
) *BaseApp {
	return &BaseApp{
		BaseApp: bapp.NewBaseApp(name, logger, db, txDecoder, options...),
	}
}

func (app *BaseApp) AddRoute(path string, handler sdkTypes.Handler) sdk.Router {
	// Wrap around every handler to ensure Fee is Called
	newHandler := func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		// our handler
		result := handler(ctx, msg)

		// perform Fee Handler
		// Tendermit skip write cache if Result is not OK
		sdkResult, _ := app.FeeHandler()(ctx, result)
		return sdkResult
	}
	return app.Router().AddRoute(path, newHandler)
}

// SetFeeHandler - SetFeeHanlder to BaseApp
func (app *BaseApp) SetFeeHandler(h fee.FeeHandler) {
	app.feeHandler = h
}

//GetFeeHandler
func (app *BaseApp) FeeHandler() fee.FeeHandler {
	return app.feeHandler
}
