package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/booking/types"
)

func (k msgServer) Complete(goCtx context.Context, msg *types.MsgComplete) (*types.MsgCompleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	oldBooking, found := k.GetBooking(ctx, msg.GetBookID())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrBookingDoesNotExist, msg.GetBookID())
	}
	if err := checkBooking(oldBooking); err != nil {
		return nil, sdkerrors.Wrap(err, msg.GetBookID())
	}

	if oldBooking.GetBooker() != msg.GetBooker() {
		return nil, sdkerrors.Wrap(types.ErrNotBookerOfAsset, msg.GetBooker())
	}

	oldAsset, found := k.GetAsset(ctx, oldBooking.GetUUID())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAssetDoesNotExist, oldBooking.GetUUID())
	}

	if oldAsset.GetStatus() {
		return nil, sdkerrors.Wrap(types.ErrAssetNotBooked, oldAsset.GetUUID())
	}

	price := oldAsset.GetRate() * oldBooking.GetDuration()
	priceCoin := sdk.NewCoin("shrp", sdk.NewInt(price))

	creator, err := sdk.AccAddressFromBech32(oldAsset.GetCreator())
	if err != nil {
		return nil, sdkerrors.Wrap(err, oldAsset.GetCreator())
	}

	if err := k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(priceCoin)); err != nil {
		return nil, sdkerrors.Wrapf(err, "cant send coins from %s to %s", types.ModuleName, oldAsset.GetCreator())
	}

	k.SetAssetStatus(ctx, oldAsset.GetUUID(), true)
	k.SetBookingCompleted(ctx, oldBooking.GetBookID())

	event := sdk.NewEvent(
		types.EventTypeBookingComplete,
		sdk.NewAttribute(types.AttributeBookingID, msg.BookID),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgCompleteResponse{}, nil
}

func checkBooking(b types.Booking) error {
	if len(b.GetBooker()) == 0 || b.GetDuration() <= 0 {
		return types.ErrInvalidBooking
	}
	if b.IsCompleted {
		return types.ErrBookingIsCompleted
	}
	return nil
}
