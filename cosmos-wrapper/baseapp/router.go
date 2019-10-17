package baseapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/sharering/shareledger/cosmos-wrapper/types"
)

func AddRoute(app *BaseApp, path string, h types.Handler) sdk.Router {

	// Wrap around every handler to ensure Fee is Called
	newHandler := func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		// our handler
		result := h(ctx, msg)

		// perform Fee Handler
		// Tendermit skip write cache if Result is not OK
		sdkResult, _ := app.FeeHandler()(ctx, result)
		return sdkResult
	}
	return app.Router().AddRoute(path, newHandler)
}
