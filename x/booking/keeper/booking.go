package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/booking/types"
)

func (k Keeper) GetBooking(ctx sdk.Context, bookID string) (types.Booking, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookIDPrefix))

	if !k.IsBookingExist(ctx, bookID) {
		return types.Booking{}, false
	}

	bz := store.Get([]byte(bookID))

	var result types.Booking

	k.cdc.MustUnmarshalLengthPrefixed(bz, &result)

	return result, true
}

func (k Keeper) SetBooking(ctx sdk.Context, bookID string, booking types.Booking) {
	// TODO: should return error
	if len(booking.GetBooker()) == 0 || len(booking.GetUUID()) == 0 || len(booking.GetBookID()) == 0 {
		return
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookIDPrefix))

	store.Set([]byte(bookID), k.cdc.MustMarshalLengthPrefixed(&booking))
}

func (k Keeper) SetBookingCompleted(ctx sdk.Context, bookID string) {
	b, found := k.GetBooking(ctx, bookID)
	if !found {
		return
	}

	b.IsCompleted = true
	k.SetBooking(ctx, bookID, b)
}

func (k Keeper) DeleteBooking(ctx sdk.Context, bookID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookIDPrefix))
	store.Delete([]byte(bookID))
}

func (k Keeper) GetBookingsIterator(ctx sdk.Context) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookIDPrefix))
}

func (k Keeper) IterateBookings(ctx sdk.Context, cb func(b types.Booking) bool) {
	iterator := k.GetBookingsIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var booking types.Booking
		k.cdc.MustUnmarshalLengthPrefixed(iterator.Value(), &booking)
		if cb(booking) {
			break
		}
	}
}

func (k Keeper) IsBookingExist(ctx sdk.Context, bookID string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookIDPrefix))
	return store.Has([]byte(bookID))
}
