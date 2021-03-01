package booking

import (
	"fmt"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgBook:
			return handleMsgBook(ctx, keeper, msg)
		case MsgBookComplete:
			return handleMsgComplete(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized booking Msg type: %v", msg.Type()))
		}
	}
}

func handleMsgBook(ctx sdk.Context, keeper Keeper, msg MsgBook) (*sdk.Result, error) {
	oldAsset := keeper.GetAsset(ctx, msg.UUID)
	if err := checkAsset(oldAsset); err != nil {
		return nil, err
	}
	newBooking := types.NewBookingFromMsgBook(msg)
	bookID, err := GenBookID(msg)
	if err != nil {
		return nil, types.ErrUnableToGenerateBookID
	}
	newBooking.BookID = bookID
	newBooking.IsCompleted = false
	price := oldAsset.Rate * msg.Duration
	priceCoin := sdk.NewCoin("shrp", sdk.NewInt(price))
	if err := keeper.SendCoinsFromAccountToModule(ctx, msg.Booker, ModuleName, sdk.NewCoins(priceCoin)); err != nil {
		return nil, err
	}
	keeper.SetAssetStatus(ctx, msg.UUID, false)
	keeper.SetBooking(ctx, bookID, newBooking)
	log, err := newBooking.GetString()
	if err != nil {
		return nil, err
	}
	event := sdk.NewEvent(
		EventTypeBookingStart,
		sdk.NewAttribute(AttributeUUID, string(msg.UUID)),
		sdk.NewAttribute(AttributeBookingID, string(bookID)),
	)
	ctx.EventManager().EmitEvent((event))
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgComplete(ctx sdk.Context, keeper Keeper, msg MsgBookComplete) (*sdk.Result, error) {
	oldBooking := keeper.GetBooking(ctx, msg.BookID)
	if err := checkBooking(oldBooking); err != nil {
		return nil, err
	}
	if !oldBooking.Booker.Equals(msg.Booker) {
		return nil, types.ErrNotBookerOfAsset
	}
	oldAsset := keeper.GetAsset(ctx, oldBooking.UUID)
	if oldAsset.Creator.Empty() {
		return nil, types.ErrAssetDoesNotExist
	}
	if oldAsset.Status {
		return nil, types.ErrAssetNotBooked
	}
	if oldAsset.UUID != oldBooking.UUID {
		return nil, types.ErrUUIDMismatch
	}
	price := oldAsset.Rate * oldBooking.Duration
	priceCoin := sdk.NewCoin("shrp", sdk.NewInt(price))
	if err := keeper.SendCoinsFromModuleToAccount(ctx, ModuleName, oldAsset.Creator, sdk.NewCoins(priceCoin)); err != nil {
		return nil, err
	}
	keeper.SetAssetStatus(ctx, oldAsset.UUID, true)
	keeper.SetBookingCompleted(ctx, oldBooking.BookID)
	newBooking := keeper.GetBooking(ctx, msg.BookID)
	log, err := newBooking.GetString()
	if err != nil {
		return nil, err
	}
	event := sdk.NewEvent(
		EventTypeBookingComplete,
		sdk.NewAttribute(AttributeBookingID, string(msg.BookID)),
	)
	ctx.EventManager().EmitEvent(event)
	return &sdk.Result{
		Log:    log,
		Events: ctx.EventManager().Events(),
	}, nil
}

func checkAsset(a asset.Asset) error {
	if a.Creator.Empty() {
		return types.ErrAssetDoesNotExist
	}
	if a.Rate <= 0 {
		return types.ErrIllegalAssetRate
	}
	if !a.Status {
		return types.ErrAssetAlreadyBooked
	}
	return nil
}

func checkBooking(b types.Booking) error {
	if len(b.BookID) == 0 || b.Booker.Empty() || b.Duration <= 0 {
		return types.ErrInvalidBooking
	}
	if b.IsCompleted {
		return types.ErrBookingIsCompleted
	}
	return nil
}
