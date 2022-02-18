package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

// SetTreasurer set treasurer in the store
func (k Keeper) SetTreasurer(ctx sdk.Context, treasurer types.Treasurer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasurerKey))
	b := k.cdc.MustMarshal(&treasurer)
	store.Set([]byte{0}, b)
}

// GetTreasurer returns treasurer
func (k Keeper) GetTreasurer(ctx sdk.Context) (val types.Treasurer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasurerKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTreasurer removes treasurer from the store
func (k Keeper) RemoveTreasurer(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TreasurerKey))
	store.Delete([]byte{0})
}
