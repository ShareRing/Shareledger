package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// SetFormat set a specific format in the store from its index
func (k Keeper) SetFormat(ctx sdk.Context, format types.SignSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormatKeyPrefix))
	b := k.cdc.MustMarshal(&format)
	store.Set(types.FormatKey(
		format.Network,
	), b)
}

// GetFormat returns a format from its index
func (k Keeper) GetFormat(
	ctx sdk.Context,
	network string,

) (val types.SignSchema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormatKeyPrefix))

	b := store.Get(types.FormatKey(
		network,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFormat removes a format from the store
func (k Keeper) RemoveFormat(
	ctx sdk.Context,
	network string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormatKeyPrefix))
	store.Delete(types.FormatKey(
		network,
	))
}

// GetAllFormat returns all format
func (k Keeper) GetAllFormat(ctx sdk.Context) (list []types.SignSchema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FormatKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignSchema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
