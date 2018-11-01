package exchange

import (
	"fmt"
	"reflect"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	msg "github.com/sharering/shareledger/x/exchange/messages"
	"github.com/sharering/shareledger/utils"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msg.MsgCreate:
			return handleFeeMsg(k.CreateExchangeRate, ctx, msg)
		case msg.MsgRetrieve:
			return handleNonFeeMsg(k.RetrieveExchangeRate, ctx, msg)
		case msg.MsgUpdate:
			return handleFeeMsg(k.UpdateExchangeRate, ctx, msg)
		case msg.MsgDelete:
			return handleFeeMsg(k.DeleteExchangeRate, ctx, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleFeeMsg(
	f func(ctx sdk.Context, msg sdk.Msg),
	ctx sdk.Context,
	msg sdk.Msg,

) sdk.Result {

	exr, err := f(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// TODO: MsgFee is based on name of Msg. Currently, Asset and This module ( Exchagne) share the same set of names
	// Create, Delete, Update
	fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", exr),
		Tags:      msg.Tags(),
		FeeAmount: fee,
		FeeDenom:  denom,
	}
}

func handleNonFeeMsg(
	f func(ctx sdk.Context, msg sdk.Msg),
	ctx sdk.Context,
	msg sdk.Msg,

) sdk.Result {

	exr, err := f(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:  fmt.Sprintf("%s", exr),
		Tags: msg.Tags(),
	}
}
