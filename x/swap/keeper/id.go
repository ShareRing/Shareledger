package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"sync"
)

var nextIDMutex sync.Mutex

// NextId return next id and update store for that new id
func (k Keeper) NextId(ctx sdk.Context, iDType string) (id uint64) {
	nextIDMutex.Lock()
	defer nextIDMutex.Unlock()

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdKeyPrefix))
	val := types.Id{
		IDType: iDType,
		Value:  0,
	}
	b := store.Get(types.IdKey(
		iDType,
	))
	if b != nil {
		k.cdc.MustUnmarshal(b, &val)
	}
	val.Value = val.Value + 1
	store.Set(types.IdKey(
		val.IDType,
	), k.cdc.MustMarshal(&val))
	return val.Value
}

// GetId returns a id from its index
func (k Keeper) GetId(
	ctx sdk.Context,
	iDType string,

) (val types.Id, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdKeyPrefix))

	b := store.Get(types.IdKey(
		iDType,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllId returns all id
func (k Keeper) GetAllId(ctx sdk.Context) (list []types.Id) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Id
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
