package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// GetBatchCount get the total number of batch
func (k Keeper) GetBatchCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BatchCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetBatchCount set the total number of batch
func (k Keeper) SetBatchCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BatchCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBatch appends a batch in the store with a new id and update the count
func (k Keeper) AppendBatch(
	ctx sdk.Context,
	batch types.Batch,
) uint64 {
	// Create the batch
	count := k.GetBatchCount(ctx)

	// Set the ID of the appended value
	batch.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	appendedValue := k.cdc.MustMarshal(&batch)
	store.Set(GetBatchIDBytes(batch.Id), appendedValue)

	// Update batch count
	count++
	k.SetBatchCount(ctx, count)

	return count
}

// SetBatch set a specific batch in the store
func (k Keeper) SetBatch(ctx sdk.Context, batch types.Batch) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	b := k.cdc.MustMarshal(&batch)
	store.Set(GetBatchIDBytes(batch.Id), b)
}

// GetBatch returns a batch from its id
func (k Keeper) GetBatch(ctx sdk.Context, id uint64) (val types.Batch, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	b := store.Get(GetBatchIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBatch removes a batch from the store
func (k Keeper) RemoveBatch(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	store.Delete(GetBatchIDBytes(id))
}

// GetAllBatch returns all batch
func (k Keeper) GetAllBatch(ctx sdk.Context) (list []types.Batch) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Batch
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBatchesByIDs returns all batch
func (k Keeper) GetBatchesByIDs(ctx sdk.Context, ids []uint64) (list []types.Batch) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	bMap := make(map[uint64]struct{})
	for i := range ids {
		bMap[ids[i]] = struct{}{}
	}

	for ; iterator.Valid(); iterator.Next() {
		var val types.Batch
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if _, f := bMap[val.GetId()]; f {
			list = append(list, val)
		}

	}
	return
}

// GetBatchIDBytes returns the byte representation of the ID
func GetBatchIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBatchIDFromBytes returns ID in uint64 format from a byte array
func GetBatchIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
