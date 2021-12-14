package keeper

import (
	"github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccState set a specific accState in the store from its index
func (k Keeper) SetAccState(ctx sdk.Context, accState types.AccState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	b := k.cdc.MustMarshal(&accState)
	store.Set(types.AccStateKey(
		types.IndexKeyAccState(accState.Key),
	), b)
}

// GetAccState returns a accState from its index
func (k Keeper) GetAccState(
	ctx sdk.Context,
	key types.IndexKeyAccState,

) (val types.AccState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))

	b := store.Get(types.AccStateKey(
		key,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccState removes a accState from the store
func (k Keeper) RemoveAccState(
	ctx sdk.Context,
	key types.IndexKeyAccState,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	store.Delete(types.AccStateKey(
		key,
	))
}

// GetAllAccState returns all accState
func (k Keeper) GetAllAccState(ctx sdk.Context) (list []types.AccState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccStateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer func() {
		if err := iterator.Close(); err != nil {
			ctx.Logger().Error(err.Error())
		}
	}()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
