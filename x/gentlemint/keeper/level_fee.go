package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// SetLevelFee set a specific levelFee in the store from its index
func (k Keeper) SetLevelFee(ctx sdk.Context, levelFee types.LevelFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LevelFeeKeyPrefix))
	b := k.cdc.MustMarshal(&levelFee)
	store.Set(types.LevelFeeKey(
		levelFee.Level,
	), b)
}

// GetLevelFee returns a levelFee from its index
func (k Keeper) GetLevelFee(
	ctx sdk.Context,
	level string,

) (val types.LevelFee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LevelFeeKeyPrefix))

	b := store.Get(types.LevelFeeKey(
		level,
	))
	if b == nil {
		// Check if it's defined default level.
		d, f := constant.DefaultFeeLevel[constant.DefaultLevel(level)]
		if !f {
			return val, false
		}
		val.Fee = d
		val.Level = level
	} else {
		k.cdc.MustUnmarshal(b, &val)
	}

	return val, true
}

// RemoveLevelFee removes a levelFee from the store
func (k Keeper) RemoveLevelFee(
	ctx sdk.Context,
	level string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LevelFeeKeyPrefix))
	store.Delete(types.LevelFeeKey(
		level,
	))
}

// GetAllLevelFee returns all levelFee
func (k Keeper) GetAllLevelFee(ctx sdk.Context) (list []types.LevelFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LevelFeeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LevelFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
