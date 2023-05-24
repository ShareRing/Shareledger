package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// SetBuilderCount set a specific builderCount in the store from its index
func (k Keeper) SetBuilderCount(ctx sdk.Context, builderCount types.BuilderCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderCountKeyPrefix))
	b := k.cdc.MustMarshal(&builderCount)
	store.Set(types.BuilderCountKey(builderCount.Index), b)
}

// GetBuilderCount returns a builderCount from its index
func (k Keeper) GetBuilderCount(ctx sdk.Context, index string) (val types.BuilderCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderCountKeyPrefix))

	b := store.Get(types.BuilderCountKey(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBuilderCount removes a builderCount from the store
func (k Keeper) RemoveBuilderCount(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderCountKeyPrefix))
	store.Delete(types.BuilderCountKey(index))
}

// GetAllBuilderCount returns all builderCount
func (k Keeper) GetAllBuilderCount(ctx sdk.Context) (list []types.BuilderCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderCountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuilderCount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IncBuilderCount increment builder count by 1
func (k Keeper) IncBuilderCount(ctx sdk.Context, address string) {
	val, found := k.GetBuilderCount(ctx, address)
	if !found {
		val = types.BuilderCount{
			Index: address,
			Count: 0,
		}
	}
	val.Count ++
	k.SetBuilderCount(ctx, val)
}
