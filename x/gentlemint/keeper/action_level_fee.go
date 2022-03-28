package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// SetActionLevelFee set a specific actionLevelFee in the store from its index
func (k Keeper) SetActionLevelFee(ctx sdk.Context, actionLevelFee types.ActionLevelFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionLevelFeeKeyPrefix))
	b := k.cdc.MustMarshal(&actionLevelFee)
	store.Set(types.ActionLevelFeeKey(
		actionLevelFee.Action,
	), b)
}

// GetActionLevelFee returns a actionLevelFee from its index
func (k Keeper) GetActionLevelFee(
	ctx sdk.Context,
	action string,

) (val types.ActionLevelFee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionLevelFeeKeyPrefix))

	b := store.Get(types.ActionLevelFeeKey(
		action,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActionLevelFee removes a actionLevelFee from the store
func (k Keeper) RemoveActionLevelFee(
	ctx sdk.Context,
	action string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionLevelFeeKeyPrefix))
	store.Delete(types.ActionLevelFeeKey(
		action,
	))
}

// GetAllActionLevelFee returns all actionLevelFee
func (k Keeper) GetAllActionLevelFee(ctx sdk.Context) (list []types.ActionLevelFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActionLevelFeeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActionLevelFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
