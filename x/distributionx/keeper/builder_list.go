package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// GetBuilderListCount get the total number of builderList
func (k Keeper) GetBuilderListCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BuilderListCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBuilderListCount set the total number of builderList
func (k Keeper) SetBuilderListCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BuilderListCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBuilderList appends a builderList in the store with a new id and update the count
func (k Keeper) AppendBuilderList(ctx sdk.Context, builderList types.BuilderList) uint64 {
	// Create the builderList
	count := k.GetBuilderListCount(ctx)

	// Set the ID of the appended value
	builderList.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderListKey))
	appendedValue := k.cdc.MustMarshal(&builderList)
	store.Set(GetBuilderListIDBytes(builderList.Id), appendedValue)

	// Update builderList count
	k.SetBuilderListCount(ctx, count+1)

	return count
}

// SetBuilderList set a specific builderList in the store
func (k Keeper) SetBuilderList(ctx sdk.Context, builderList types.BuilderList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderListKey))
	b := k.cdc.MustMarshal(&builderList)
	store.Set(GetBuilderListIDBytes(builderList.Id), b)
}

// GetBuilderList returns a builderList from its id
func (k Keeper) GetBuilderList(ctx sdk.Context, id uint64) (val types.BuilderList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderListKey))
	b := store.Get(GetBuilderListIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBuilderList removes a builderList from the store
func (k Keeper) RemoveBuilderList(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderListKey))
	store.Delete(GetBuilderListIDBytes(id))
}

// GetAllBuilderList returns all builderList
func (k Keeper) GetAllBuilderList(ctx sdk.Context) (list []types.BuilderList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuilderListKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuilderList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBuilderListIDBytes returns the byte representation of the ID
func GetBuilderListIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBuilderListIDFromBytes returns ID in uint64 format from a byte array
func GetBuilderListIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
