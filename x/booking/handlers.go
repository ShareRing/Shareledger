package booking

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/booking/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

func NewHandler(k Keeper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		ctx.WithEventManager(sdk.NewEventManager())
		constants.LOGGER.Info(
			"Msg for Asset Module",
			"type", reflect.TypeOf(msg),
			"msg", msg,
		)

		var ret sdk.Result

		switch msg := msg.(type) {
		case messages.MsgBook:
			ret = handleBooking(ctx, k, msg)
		case messages.MsgComplete:
			ret = handleComplete(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("Unrecognized trace Msg type: %v", reflect.TypeOf(msg).Name())
			return sdkTypes.NewResult(sdk.ErrUnknownRequest(errMsg).Result())
		}

		if !ret.IsOK() {
			return sdkTypes.NewResult(ret)
		}

		fee, denom := utils.GetMsgFee(msg)

		return sdkTypes.Result{
			Result:    ret,
			FeeDenom:  denom,
			FeeAmount: fee,
		}

	}
}

func handleBooking(ctx sdk.Context, k Keeper, msg messages.MsgBook) sdk.Result {

	booking, err := k.Book(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}

	// fee, denom := utils.GetMsgFee(msg)
	event := sdk.NewEvent(
		EventTypeBookingStart,
		sdk.NewAttribute(AttributeUUID, string(msg.UUID)),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    fmt.Sprintf("%s", booking.String()),
		Events: ctx.EventManager().Events(),
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
	event := sdk.NewEvent(
		EventTypeBookingComplete,
		sdk.NewAttribute(AttributeBookingID, string(msg.BookingID)),
	)
	ctx.EventManager().EmitEvent(event)
	return sdk.Result{
		Log:    booking.String(),
		Events: ctx.EventManager().Events(),
		// FeeAmount: fee,
		// FeeDenom:  denom,
	}
}
