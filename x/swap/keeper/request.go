package keeper

import (
	"encoding/binary"
	"fmt"
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
	var requestCount uint64

	if bz != nil {
		requestCount = binary.BigEndian.Uint64(bz)
	}

	return requestCount + 1
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
	request.Id = count
	k.SetRequestCount(ctx, count)

	store := k.GetStoreRequestMap(ctx)[request.Status]
	appendedValue := k.cdc.MustMarshal(&request)
	store.Set(GetRequestIDBytes(request.Id), appendedValue)

	return request, nil
}

func (k Keeper) changeStatusSwapOut(ctx sdk.Context, ids []uint64, fromStatus string, toStatus string, batchID *uint64) ([]types.Request, error) {
	reqs, err := k.getRequestsFromIds(ctx, ids, fromStatus)
	if err != nil {
		return nil, err
	}
	destNet := reqs[0].DestNetwork
	for i := range reqs {
		//source network swap out case must is slp3
		if reqs[i].SrcNetwork != types.NetworkNameShareLedger {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same source network %v", reqs[i].Id, types.NetworkNameShareLedger)
		}
		if reqs[i].DestNetwork != destNet {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same dest network %v", reqs[i].Id, destNet)
		}
	}

	if toStatus == types.SwapStatusApproved {
		if batchID == nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "%s required batchID", toStatus)
		}
		b, f := k.GetBatch(ctx, *batchID)
		if !f {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "fail to set batch network")
		}
		b.Network = destNet
		k.SetBatch(ctx, b)
	}

	if err := k.MoveRequest(ctx, fromStatus, toStatus, reqs, batchID, true); err != nil {
		return nil, err
	}

	return reqs, nil
}

func (k Keeper) getRequestsFromIds(ctx sdk.Context, ids []uint64, status string) ([]types.Request, error) {
	currentStatusStore := k.GetStoreRequestMap(ctx)[status]
	reqs := k.GetRequestsByIdsFromStore(ctx, currentStatusStore, ids)
	if len(reqs) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request not found with required status")
	}
	if len(reqs) != len(ids) {
		foundReqs := make(map[uint64]struct{})
		for _, req := range reqs {
			foundReqs[req.Id] = struct{}{}
		}
		notFoundIDs := make([]uint64, 0, len(ids))
		for _, id := range ids {
			if _, found := foundReqs[id]; !found {
				notFoundIDs = append(notFoundIDs, id)
			}
		}
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "there are some not found pending request, %+v", fmt.Sprint(notFoundIDs))
	}
	return reqs, nil
}

func (k Keeper) changeStatusSwapIn(ctx sdk.Context, ids []uint64, fromStatus string, toStatus string) ([]types.Request, error) {
	reqs, err := k.getRequestsFromIds(ctx, ids, fromStatus)
	if err != nil {
		return nil, err
	}

	for i := range reqs {
		if reqs[i].DestNetwork != types.NetworkNameShareLedger {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same destination network %v", reqs[i].Id, types.NetworkNameShareLedger)
		}
	}
	if toStatus == types.SwapStatusApproved {
		// Done after sending coin to destination address in Shareledger

		transfers := make(map[string]sdk.Coins)
		for i := range reqs {
			if toStatus == types.SwapStatusApproved {
				total, found := transfers[reqs[i].DestAddr]
				if !found {
					total = sdk.NewCoins()
				}
				transfers[reqs[i].DestAddr] = total.Add(reqs[i].Amount)
			}
		}
		for dest, t := range transfers {
			destAddr, err := sdk.AccAddressFromBech32(dest)
			if err != nil {
				return nil, err
			}
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, destAddr, t); err != nil {
				return nil, err
			}
		}
		toStatus = types.SwapStatusDone
	}
	if err := k.MoveRequest(ctx, fromStatus, toStatus, reqs, nil, false); err != nil {
		return nil, err
	}

	return reqs, nil
}

// ChangeStatusRequests change status of requests and move it into respective store
// The status flow should be: pending -> approved|rejected -> done.
// return error if new status is pending or unsupported status
func (k Keeper) ChangeStatusRequests(ctx sdk.Context, ids []uint64, status string, batchId *uint64, isSwapOut bool) ([]types.Request, error) {
	if len(ids) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap request transactions' id is empty")
	}

	if isSwapOut && status == types.SwapStatusApproved {
		if batchId == nil || *batchId == 0 {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "approve swap out must provide batch id")
		}
	}

	//validate the request

	var requiredStatus string
	switch status {
	// pending swap can be approved or rejected
	case types.SwapStatusApproved, types.SwapStatusRejected:
		requiredStatus = types.SwapStatusPending
	// approved swap can be change to pending when un-batching action happens
	case types.SwapStatusPending:
		requiredStatus = types.SwapStatusApproved
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is not supported", status)
	}

	if isSwapOut {
		return k.changeStatusSwapOut(ctx, ids, requiredStatus, status, batchId)
	} else {
		return k.changeStatusSwapIn(ctx, ids, requiredStatus, status)
	}
}

func (k Keeper) RemoveRequests(ctx sdk.Context, currentStatus string, reqs []types.Request) error {
	currentStore, found := k.GetStoreRequestMap(ctx)[currentStatus]
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic, "store not found")
	}
	for i := range reqs {
		req := &reqs[i]
		currentStore.Delete(GetRequestIDBytes(req.Id))
	}
	return nil
}

// MoveRequest move the request to the store base on status
// Delete request fromStt store and add this request to store. This action also update batch ID if it's existed
// If toStt == Done || Rejected, the function will only delete from fromStt store.
func (k Keeper) MoveRequest(ctx sdk.Context, fromStt, toStt string, reqs []types.Request, batchID *uint64, isOut bool) error {
	// There is no need to store done/rejected swap requests in keeper

	err := k.RemoveRequests(ctx, fromStt, reqs)
	if err != nil {
		return sdkerrors.Wrapf(err, "fail to remove request from %s store", fromStt)
	}
	storeMap := k.GetStoreRequestMap(ctx)
	var idsStr = make([]string, 0, len(reqs))
	for i := range reqs {
		req := &reqs[i]
		req.Status = toStt
		if batchID != nil {
			req.BatchId = *batchID
		}
		//Just insert to store if
		if toStt != types.SwapStatusDone && toStt != types.SwapStatusRejected {
			toStore, found := storeMap[toStt]
			if !found {
				return sdkerrors.Wrapf(sdkerrors.ErrLogic, "store not found")
			}
			toStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(req))
		}
		idsStr = append(idsStr, fmt.Sprintf("%v", req.Id))

	}
	ctx.EventManager().EmitEvent(
		types.NewChangeRequestStatusesEvent(idsStr, toStt),
	)
	return nil
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
	val = types.Request{}
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
