package booking

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	// "github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/booking/messages"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {

		constants.LOGGER.Info(
			"Msg for Asset Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		switch msg := msg.(type) {
		case messages.MsgBook:
			return handleBooking(ctx, k, msg)
		case messages.MsgComplete:
			return handleComplete(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleBooking(ctx sdk.Context, k Keeper, msg messages.MsgBook) sdk.Result {

	booking, err := k.Book(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("%s", booking.String()),
		Tags:      msg.Tags(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}

func handleComplete(ctx sdk.Context, k Keeper, msg messages.MsgComplete) sdk.Result {

	booking, err := k.Complete(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)

	return sdk.Result{
		Log:       fmt.Sprintf("Completed %s", booking.String()),
		Tags:      msg.Tags(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}
