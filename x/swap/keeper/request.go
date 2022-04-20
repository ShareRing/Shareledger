package keeper

import (
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
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

// ChangeStatusRequests change status of requests and move it into respective store
// The status flow should be: pending -> approved|rejected -> processing -> done.
// return error if new status is pending or unsupported status
func (k Keeper) ChangeStatusRequests(ctx sdk.Context, ids []uint64, status string, batchId *uint64) ([]types.Request, error) {
	if len(ids) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap request transactions' id is empty")
	}
	if status == types.SwapStatusApproved && (batchId == nil || *batchId == 0) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s status requires batch Id", status)
	}

	var requiredStatus string
	var currentStatusStore prefix.Store
	switch status {
	case types.SwapStatusApproved, types.SwapStatusRejected:
		requiredStatus = types.SwapStatusPending
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is not supported", status)
	}
	currentStatusStore = k.GetStoreRequestMap(ctx)[requiredStatus]
	// Check type swap in out
	req, found := k.GetRequestFromStore(ctx, currentStatusStore, ids[0])
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap request was not found with status, %s, and id, %v", requiredStatus, ids[0])
	}
	srcNetwork := req.SrcNetwork
	requests := k.GetRequestsByIdsFromStore(ctx, currentStatusStore, ids)
	if len(requests) != len(ids) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("transactions don't have same status or not found, with required current status, %s", requiredStatus))
	}
	var bid uint64
	if batchId != nil {
		bid = *batchId
	}

	switch srcNetwork {
	case types.NetworkNameShareLedger: //cover case swap out request
		return k.changeStatusSwapOut(ctx, requiredStatus, status, ids, bid)
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "not supported yet")
}

func (k Keeper) changeStatusSwapOut(ctx sdk.Context, fromStatus string, toStatus string, ids []uint64, batchId uint64) ([]types.Request, error) {
	if batchId == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "batch id is required")
	}
	fromStatusStore := k.GetStoreRequestMap(ctx)[fromStatus]
	toStatusStore := k.GetStoreRequestMap(ctx)[toStatus]
	requests := k.GetRequestsByIdsFromStore(ctx, fromStatusStore, ids)
	if len(requests) != len(ids) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("transactions don't have same status or not found, with required current status, %s", fromStatus))
	}

	refunds := make(map[string]sdk.DecCoins)
	for i := range requests {
		req := &requests[i]
		if req.SrcNetwork != types.NetworkNameShareLedger {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same source network %v", req.Id, types.NetworkNameShareLedger)
		}
		fromStatusStore.Delete(GetRequestIDBytes(req.Id))
		req.Status = toStatus
		req.BatchId = batchId
		toStatusStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(req))
		if toStatus == types.SwapStatusRejected {
			total, found := refunds[req.DestAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*req.Amount).Add(*req.Fee)
			refunds[req.DestAddr] = total
		}
	}

	for a, c := range refunds {
		add, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
		}
		bc, err := denom.NormalizeToBaseCoins(c, false)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, add, bc); err != nil {
			return nil, err
		}
	}
	return requests, nil
}

//func (k Keeper) RemoveRequest(ctx sdk.Context, id uint64, status string) {
//	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey(status)))
//	store.Delete(GetRequestIDBytes(id))
//}

func (k Keeper) GetRequestFromStore(ctx sdk.Context, store sdk.KVStore, id uint64) (val types.Request, found bool) {
	b := store.Get(GetRequestIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetRequestsByIdsFromStore get al requests by ids
func (k Keeper) GetRequestsByIdsFromStore(ctx sdk.Context, store sdk.KVStore, ids []uint64) []types.Request {
	requests := make([]types.Request, 0, len(ids))
	mapIds := make(map[uint64]struct{})
	for _, id := range ids {
		mapIds[id] = struct{}{}
	}
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if _, found := mapIds[val.Id]; found {
			requests = append(requests, val)
		}
	}
	return requests
}

// GetRequest returns a request from its id
func (k Keeper) GetRequest(ctx sdk.Context, id uint64) (val types.Request, found bool) {
	stores := k.GetStoreRequestMap(ctx)
	for _, store := range stores {
		b := store.Get(GetRequestIDBytes(id))
		if b == nil {
			continue
		}
		k.cdc.MustUnmarshal(b, &val)
		found = true
		break
	}
	return
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
