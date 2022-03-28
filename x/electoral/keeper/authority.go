package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

// SetAuthority set authority in the store
func (k Keeper) SetAuthority(ctx sdk.Context, authority types.Authority) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorityKey))
	b := k.cdc.MustMarshal(&authority)
	store.Set([]byte{0}, b)
}

// GetAuthority returns authority
func (k Keeper) GetAuthority(ctx sdk.Context) (val types.Authority, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorityKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuthority removes authority from the store
func (k Keeper) RemoveAuthority(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorityKey))
	store.Delete([]byte{0})
}
