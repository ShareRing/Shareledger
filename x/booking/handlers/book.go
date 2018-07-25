package handlers

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/booking"
	"github.com/sharering/shareledger/x/booking/messages"
	"fmt"
	"reflect"
)

func NewHandler(k booking.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
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

func handleBooking(ctx sdk.Context, k booking.Keeper, msg messages.MsgBook) sdk.Result {

	booking, err := k.Book(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}


	return sdk.Result{
		Log:  fmt.Sprintf("%s", booking.String()),
		Tags: msg.Tags(),
	}
}
func handleComplete(ctx sdk.Context, k booking.Keeper, msg messages.MsgComplete) sdk.Result {

	booking, err := k.Complete(ctx, msg)

	if err != nil {
		return sdk.ErrInternal(err.Error()).Result()
	}


	return sdk.Result{
		Log:  fmt.Sprintf("Completed %s", booking.String()),
		Tags: msg.Tags(),
	}
}