package auth

import (
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

func NewHandler(am AccountMapper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		ctx.WithEventManager(sdk.NewEventManager())
		constants.LOGGER.Info(
			"Msg for Auth Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		var ret sdk.Result

		switch msg := msg.(type) {
		case MsgNonce:
			ret = handleNonceQuery(ctx, am, msg)
		default:
			errMsg := "Unrecognized Auth Msg type" + reflect.TypeOf(msg).Name()
			ret = sdk.ErrUnknownRequest(errMsg).Result()
		}
		return sdkTypes.Result{
			Result: ret,
		}
	}
}

func handleNonceQuery(ctx sdk.Context, am AccountMapper, msg MsgNonce) sdk.Result {
	nonce, err := am.GetNonce(ctx, msg.Address)
	if err != nil {
		return err.Result()
	}
	event := sdk.NewEvent(
		EventTypeCheckNonce,
		sdk.NewAttribute(AttributeAccountAddress, msg.Address.String()),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    strconv.FormatInt(nonce, 10), // use FormatInt so as to accept int64 type
		Events: ctx.EventManager().Events(),
	}
}
