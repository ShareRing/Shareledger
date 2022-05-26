package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	commontypes "github.com/sharering/shareledger/x/common/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// SetRequestedIn set a specific requestedIn in the store from its index
func (k Keeper) SetRequestedIn(ctx sdk.Context, destAddress sdk.Address, txHashes []string) {
	insertingData := types.RequestedIn{
		Address: destAddress.String(),
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestedInKeyPrefix))
	currentData, found := k.GetRequestedIn(ctx, destAddress.String())
	if found {
		insertingData = currentData
	}

	if insertingData.TxHashes == nil {
		insertingData.TxHashes = make(map[string]*commontypes.Empty)
	}
	for _, hash := range txHashes {
		insertingData.TxHashes[hash] = &commontypes.Empty{}
	}
	b := k.cdc.MustMarshal(&insertingData)
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
