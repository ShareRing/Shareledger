package exchange

import (
	"fmt"
	"reflect"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/exchange/messages"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case messages.MsgCreate:
			return handleMsgCreate(ctx, k, msg)
		case messages.MsgRetrieve:
			return handleMsgRetrieve(ctx, k, msg)
		case messages.MsgUpdate:
			return handleMsgUpdate(ctx, k, msg)
		case messages.MsgDelete:
			return handleMsgDelete(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgCreate(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgCreate,
) sdk.Result {

	exr, err := k.CreateExchangeRate(ctx, msg)

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

func handleMsgUpdate(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgUpdate,
) sdk.Result {

	exr, err := k.UpdateExchangeRate(ctx, msg)

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

func handleMsgDelete(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgDelete,
) sdk.Result {

	exr, err := k.DeleteExchangeRate(ctx, msg)

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
func handleMsgRetrieve(
	ctx sdk.Context,
	k Keeper,
	msg messages.MsgRetrieve,
) sdk.Result {

	exr, err := k.RetrieveExchangeRate(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}


	return sdk.Result{
		Log:  fmt.Sprintf("%s", exr),
		Tags: msg.Tags(),
	}
}
