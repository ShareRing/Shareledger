package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	assetTypes "github.com/ShareRing/Shareledger/x/asset/types"
	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Book(goCtx context.Context, msg *types.MsgBook) (*types.MsgBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	oldAsset, found := k.GetAsset(ctx, msg.GetUUID())
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAssetDoesNotExist, msg.GetUUID())
	}
	if err := checkAsset(oldAsset); err != nil {
		return nil, sdkerrors.Wrap(err, msg.GetUUID())
	}

	newBooking := types.NewBookingFromMsgBook(*msg)
	bookID, err := GenBookID(msg)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrUnableToGenerateBookID, bookID)
	}

	newBooking.BookID = bookID
	newBooking.IsCompleted = false
	price := oldAsset.GetRate() * msg.GetDuration()
	priceCoin := sdk.NewCoin("shrp", sdk.NewInt(price))
	booker, _ := sdk.AccAddressFromBech32(msg.GetBooker())

	if err := k.SendCoinsFromAccountToModule(ctx, booker, types.ModuleName, sdk.NewCoins(priceCoin)); err != nil {
		return nil, sdkerrors.Wrapf(err, "cant send coin from %s to %s", msg.GetBooker(), types.ModuleName)
	}

	k.SetAssetStatus(ctx, msg.GetUUID(), false)
	k.SetBooking(ctx, bookID, newBooking)

	event := sdk.NewEvent(
		types.EventTypeBookingStart,
		sdk.NewAttribute(types.AttributeUUID, string(msg.GetUUID())),
		sdk.NewAttribute(types.AttributeBookingID, string(bookID)),
	)
	ctx.EventManager().EmitEvent((event))

	return &types.MsgBookResponse{}, nil
}

func checkAsset(a assetTypes.Asset) error {
	if a.GetRate() <= 0 {
		return types.ErrIllegalAssetRate
	}
	if !a.GetStatus() {
		return types.ErrAssetAlreadyBooked
	}
	return nil
}

// TODO: deterministic problem ?
func GenBookID(inp interface{}) (string, error) {
	h := sha256.New()

	enc, err := json.Marshal(inp)
	if err != nil {
		return "", err
	}

	h.Write(enc)
	hash := h.Sum(nil)

	return hex.EncodeToString(hash)[:20], nil
}