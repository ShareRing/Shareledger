package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

// GetRequestCount get the total number of request
func (k Keeper) GetRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRequestCount set the total number of request
func (k Keeper) SetRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPendingRequest appends a request in the store with a new id and update the count.
// NewID will be generated
// Return new data
func (k Keeper) AppendPendingRequest(
	ctx sdk.Context,
	request types.Request,
) (types.Request, error) {

	if request.Status != types.SwapStatusPending {
		return request, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request should have status pending")
	}
	// Create the request
	count := k.GetRequestCount(ctx)
	// Set the ID of the appended value
	// We are using 0 as zero value for id.
	request.Id = count + 1
	k.SetRequestCount(ctx, count+1)

	store := k.GetStoreRequestMap(ctx)[request.Status]
	appendedValue := k.cdc.MustMarshal(&request)
	store.Set(GetRequestIDBytes(request.Id), appendedValue)

	return request, nil
}

// ChangeStatusRequest change status of request and move it into respective store
// The status flow should be: pending -> approved|rejected -> processing -> done.
// return error if new status is pending or unsupported status
func (k Keeper) ChangeStatusRequest(ctx sdk.Context, id uint64, status string) (req types.Request, err error) {
	var found bool
	var selectedStatus string
	var currentStatusStore prefix.Store
	switch status {
	case types.SwapStatusApproved, types.SwapStatusReject:
		selectedStatus = types.SwapStatusPending
	case types.SwapStatusProcessing:
		selectedStatus = types.SwapStatusApproved
	case types.SwapStatusDone:
		selectedStatus = types.SwapStatusProcessing
	default:
		return req, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is not support for this function", status)
	}
	currentStatusStore = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey(types.SwapStatusPending)))
	req, found = k.GetRequestFromStore(ctx, currentStatusStore, id)
	if !found {
		return req, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "there is no request with id, %v, and status, %s", id, selectedStatus)
	}
	currentStatusStore.Delete(GetRequestIDBytes(id))
	req.Status = status
	newStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey(req.Status)))
	newStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(&req))
	return req, nil
}

func (k Keeper) RemoveRequest(ctx sdk.Context, id uint64, status string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey(status)))
	store.Delete(GetRequestIDBytes(id))
}

func (k Keeper) GetRequestFromStore(ctx sdk.Context, store sdk.KVStore, id uint64) (val types.Request, found bool) {
	b := store.Get(GetRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetRequest returns a request from its id
func (k Keeper) GetRequest(ctx sdk.Context, id uint64, status string) (val types.Request, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey(status)))
	b := store.Get(GetRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllRequest returns all request
func (k Keeper) GetAllRequest(ctx sdk.Context) (list []types.Request) {
	stores := k.GetStoreRequestMap(ctx)
	for _, store := range stores {
		func(store prefix.Store) {
			iterator := sdk.KVStorePrefixIterator(store, []byte{})
			defer iterator.Close()

			for ; iterator.Valid(); iterator.Next() {
				var val types.Request
				k.cdc.MustUnmarshal(iterator.Value(), &val)
				list = append(list, val)
			}
		}(store)
	}
	return
}

// GetRequestIDBytes returns the byte representation of the ID
func GetRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRequestIDFromBytes returns ID in uint64 format from a byte array
func GetRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
