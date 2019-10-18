package bank

import (
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank/handlers"
	"github.com/sharering/shareledger/x/bank/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

func NewHandler(am auth.AccountMapper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		ctx.WithEventManager(sdk.NewEventManager())
		constants.LOGGER.Info(
			"Msg for Bank Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		switch msg := msg.(type) {
		// case messages.MsgCheck:
		// return handlers.HandleMsgCheck(am)(ctx, msg)
		case messages.MsgLoad:
			return handlers.HandleMsgLoad(am)(ctx, msg)
		case messages.MsgSend:
			return handlers.HandleMsgSend(am)(ctx, msg)
		case messages.MsgBurn:
			return handlers.HandleMsgBurn(am)(ctx, msg)
		default:
			errMsg := "Unrecognized bank Msg type" + reflect.TypeOf(msg).Name()
			return sdkTypes.NewResult(sdk.ErrUnknownRequest(errMsg).Result())
		}
	}
}
