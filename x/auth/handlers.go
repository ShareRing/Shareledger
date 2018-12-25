package auth

import (
	"reflect"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
)

func NewHandler(am AccountMapper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		constants.LOGGER.Info(
			"Msg for Auth Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		switch msg := msg.(type) {
		case MsgNonce:
			return handleNonceQuery(ctx, am, msg)
		default:
			errMsg := "Unrecognized Auth Msg type" + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleNonceQuery(ctx sdk.Context, am AccountMapper, msg MsgNonce) sdk.Result {
	nonce, err := am.GetNonce(ctx, msg.Address)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Log:  strconv.FormatInt(nonce, 10), // use FormatInt so as to accept int64 type
		Tags: msg.Tags(),
	}
}
