package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// SetSchema set a specific format in the store from its index
func (k Keeper) SetSchema(ctx sdk.Context, format types.Schema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignSchemaKeyPrefix))
	b := k.cdc.MustMarshal(&format)
	store.Set(types.FormatKey(
		format.Network,
	), b)
}

// GetSchema returns a format from its index
func (k Keeper) GetSchema(
	ctx sdk.Context,
	network string,

) (val types.Schema, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignSchemaKeyPrefix))

	b := store.Get(types.FormatKey(
		network,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSchema removes a format from the store
func (k Keeper) RemoveSchema(
	ctx sdk.Context,
	network string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignSchemaKeyPrefix))
	store.Delete(types.FormatKey(
		network,
	))
}

// GetAllSchema returns all format
func (k Keeper) GetAllSchema(ctx sdk.Context) (list []types.Schema) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignSchemaKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Schema
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
