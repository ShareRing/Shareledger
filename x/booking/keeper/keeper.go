package keeper

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/asset"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/booking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

type Keeper struct {
	storeKey     sdk.StoreKey
	cdc          *codec.Codec
	assetKeeper  asset.Keeper
	supplyKeeper supply.Keeper
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, ask asset.Keeper, supplyKeeper supply.Keeper) Keeper {
	return Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		assetKeeper:  ask,
		supplyKeeper: supplyKeeper,
	}
}

func (k Keeper) GetBooking(ctx sdk.Context, bookID string) types.Booking {
	store := ctx.KVStore(k.storeKey)

	if !k.IsBookingPresent(ctx, bookID) {
		return types.NewBooking()
	}

	bz := store.Get([]byte(bookID))

	var result types.Booking

	k.cdc.MustUnmarshalBinaryBare(bz, &result)

	return result
}

func (k Keeper) SetBooking(ctx sdk.Context, bookID string, b types.Booking) {
	if b.Booker.Empty() || len(b.UUID) == 0 || len(b.BookID) == 0 {
		return
	}

	store := ctx.KVStore(k.storeKey)

	store.Set([]byte(bookID), k.cdc.MustMarshalBinaryBare(b))
}

func (k Keeper) SetBookingCompleted(ctx sdk.Context, bookID string) {
	b := k.GetBooking(ctx, bookID)
	b.IsCompleted = true
	k.SetBooking(ctx, bookID, b)
}

func (k Keeper) DeleteBooking(ctx sdk.Context, bookID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(bookID))
}

func (k Keeper) GetBookingsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) IterateBookings(ctx sdk.Context, cb func(b types.Booking) bool) {
	iterator := k.GetBookingsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var booking types.Booking
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &booking)
		if cb(booking) {
			break
		}
	}
}

func (k Keeper) IsBookingPresent(ctx sdk.Context, bookID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(bookID))
}

func (k Keeper) SetAssetStatus(ctx sdk.Context, uuid string, status bool) {
	k.assetKeeper.SetAssetStatus(ctx, uuid, status)
}

func (k Keeper) GetAsset(ctx sdk.Context, uuid string) asset.Asset {
	return k.assetKeeper.GetAsset(ctx, uuid)
}

func (k Keeper) SendCoinsFromModuleToAccount(
	ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins,
) error {
	return k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k Keeper) SendCoinsFromAccountToModule(
	ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins,
) error {
	return k.supplyKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}
