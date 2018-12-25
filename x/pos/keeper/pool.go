package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// load/save the pool
func (k Keeper) GetPool(ctx sdk.Context) (pool posTypes.Pool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(PoolKey)
	if b == nil {
		panic("Stored pool should not have been nil")
	}
	k.cdc.MustUnmarshalBinary(b, &pool)
	return
}

// set the pool
func (k Keeper) SetPool(ctx sdk.Context, pool posTypes.Pool) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(pool)
	store.Set(PoolKey, b)
}
