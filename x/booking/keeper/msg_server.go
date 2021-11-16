package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	assetTypes "github.com/ShareRing/Shareledger/x/asset/types"
	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Book(goCtx context.Context, msg *types.MsgBook) (*types.MsgBookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oldAsset := k.GetAsset(ctx, msg.UUID)
	if err := checkAsset(oldAsset); err != nil {
		return nil, err
	}
	newBooking := types.NewBookingFromMsgBook(*msg)
	bookID, err := GenBookID(msg)
	if err != nil {
		return nil, types.ErrUnableToGenerateBookID
	}
	newBooking.BookID = bookID
	newBooking.IsCompleted = false
	price := oldAsset.Rate * msg.Duration
	priceCoin := sdk.NewCoin("shrp", sdk.NewInt(price))
	booker, _ := sdk.AccAddressFromBech32(msg.Booker)

	if err := k.SendCoinsFromAccountToModule(ctx, booker, types.ModuleName, sdk.NewCoins(priceCoin)); err != nil {
		return nil, err
	}

	k.SetAssetStatus(ctx, msg.UUID, false)
	k.SetBooking(ctx, bookID, newBooking)
	// log, err := newBooking.GetString()
	// if err != nil {
	// 	return nil, err
	// }
	event := sdk.NewEvent(
		types.EventTypeBookingStart,
		sdk.NewAttribute(types.AttributeUUID, string(msg.UUID)),
		sdk.NewAttribute(types.AttributeBookingID, string(bookID)),
	)
	ctx.EventManager().EmitEvent((event))

	return &types.MsgBookResponse{}, nil
}

func (k msgServer) Complete(goCtx context.Context, msg *types.MsgComplete) (*types.MsgCompleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oldBooking := k.GetBooking(ctx, msg.BookID)
	if err := checkBooking(oldBooking); err != nil {
		return nil, err
	}
	if oldBooking.Booker != msg.Booker {
		return nil, types.ErrNotBookerOfAsset
	}
	oldAsset := k.GetAsset(ctx, oldBooking.UUID)

	if len(oldAsset.Creator) == 0 {
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

	creator, _ := sdk.AccAddressFromBech32(oldAsset.Creator)

	if err := k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(priceCoin)); err != nil {
		return nil, err
	}

	k.SetAssetStatus(ctx, oldAsset.UUID, true)
	k.SetBookingCompleted(ctx, oldBooking.BookID)
	// newBooking := k.GetBooking(ctx, msg.BookID)
	// log, err := newBooking.GetString()

	// if err != nil {
	// 	return nil, err
	// }

	event := sdk.NewEvent(
		types.EventTypeBookingComplete,
		sdk.NewAttribute(types.AttributeBookingID, msg.BookID),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgCompleteResponse{}, nil
}

func checkAsset(a assetTypes.Asset) error {
	if len(a.Creator) == 0 {
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
	if len(b.BookID) == 0 || len(b.Booker) == 0 || b.Duration <= 0 {
		return types.ErrInvalidBooking
	}
	if b.IsCompleted {
		return types.ErrBookingIsCompleted
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
