package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"sort"
)

// SetRequestedIn set a specific requestedIn in the store from its index
func (k Keeper) SetRequestedIn(ctx sdk.Context, destAddress sdk.Address, txHashes []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	currentData, found := k.GetRequestedIn(ctx, destAddress.String())

	if !found {
		currentData.Address = destAddress.String()
		currentData.TxHashes = make([]string, 0, len(txHashes))
	}
	if !sort.StringsAreSorted(currentData.TxHashes) {
		sort.Strings(currentData.TxHashes)
	}

	for _, txHash := range txHashes {
		// should store as sorted slice of strings  ascending order.
		insertIndex := sort.SearchStrings(currentData.TxHashes, txHash)
		currentData.TxHashes = append(currentData.TxHashes, "")
		copy(currentData.TxHashes[insertIndex+1:], currentData.TxHashes[insertIndex:])
		currentData.TxHashes[insertIndex] = txHash
	}

	b := k.cdc.MustMarshal(&currentData)
	store.Set(types.RequestedInKey(
		destAddress.String(),
	), b)
}

// GetRequestedIn returns a requestedIn from its index
func (k Keeper) GetRequestedIn(
	ctx sdk.Context,
	address string,

) (val types.RequestedIn, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))

	b := store.Get(types.RequestedInKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveRequestedIn removes a requestedIn from the store
func (k Keeper) RemoveRequestedIn(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	store.Delete(types.RequestedInKey(
		address,
	))
}

// GetAllRequestedIn returns all requestedIn
func (k Keeper) GetAllRequestedIn(ctx sdk.Context) (list []types.RequestedIn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RequestedIn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
