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
	var requestCount uint64

	if bz != nil {
		requestCount = binary.BigEndian.Uint64(bz)
	}
	if requestCount == 0 {
		requestCount = 1
	}
	return requestCount
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

	var bid uint64
	if batchId != nil {
		bid = *batchId
	}

	switch srcNetwork {
	case types.NetworkNameShareLedger: //cover case swap out request
		return k.changeStatusSwapOut(ctx, requiredStatus, status, ids, bid)
	default: //swap out
		if req.DestNetwork == types.NetworkNameShareLedger {
			return k.changeStatusSwapIn(ctx, requiredStatus, status, ids)
		}
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "not supported yet")
}

func (k Keeper) changeStatusSwapIn(ctx sdk.Context, fromStatus string, toStatus string, ids []uint64) ([]types.Request, error) {
	//approved the swapping requests wil be changed to done
	if toStatus == types.SwapStatusApproved {
		toStatus = types.SwapStatusDone
	}

	fromStatusStore := k.GetStoreRequestMap(ctx)[fromStatus]
	toStatusStore := k.GetStoreRequestMap(ctx)[toStatus]
	requests := k.GetRequestsByIdsFromStore(ctx, fromStatusStore, ids)
	if len(requests) != len(ids) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("transactions don't have same status or not found, with required current status, %s", fromStatus))
	}
	k.Logger(ctx).Debug("changing status swap in", "stt", toStatus, "ids", ids)
	transfers := make(map[string]sdk.DecCoins)
	for i := range requests {
		req := &requests[i]
		if req.SrcNetwork == types.NetworkNameShareLedger || req.DestNetwork != types.NetworkNameShareLedger {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap in with id, %v, has invalid source network, %v, or dest network, %v", req.Id, req.SrcNetwork, req.DestNetwork)
		}
		req.Status = toStatus
		toStatusStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(req))
		if toStatus == types.SwapStatusDone {
			total, found := transfers[req.DestAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*req.Amount)
			transfers[req.DestAddr] = total
		}
	}
	k.Logger(ctx).Debug("transferring shr to account....", "transfers_map", transfers)
	for a, c := range transfers {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bc, err := denom.NormalizeToBaseCoins(c, false)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}
		k.Logger(ctx).Debug("transferring shr to account....", "amount", bc.String(), "address", addr)
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, bc); err != nil {
			return nil, err
		}
	}
	return requests, nil
}

func (k Keeper) changeStatusSwapOut(ctx sdk.Context, fromStatus string, toStatus string, ids []uint64, batchId uint64) ([]types.Request, error) {

	if toStatus == types.SwapStatusApproved && (batchId == 0) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s status requires batch Id", toStatus)
	}

	fromStatusStore := k.GetStoreRequestMap(ctx)[fromStatus]
	toStatusStore := k.GetStoreRequestMap(ctx)[toStatus]
	requests := k.GetRequestsByIdsFromStore(ctx, fromStatusStore, ids)
	if len(requests) != len(ids) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("transactions don't have same status or not found, with required current status, %s", fromStatus))
	}

	destNetwork := requests[0].DestNetwork
	refunds := make(map[string]sdk.DecCoins)
	for i := range requests {
		req := &requests[i]
		if req.SrcNetwork != types.NetworkNameShareLedger {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same source network %v", req.Id, types.NetworkNameShareLedger)
		}
		if req.DestNetwork != destNetwork {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same destination network %v", req.Id, destNetwork)
		}
		fromStatusStore.Delete(GetRequestIDBytes(req.Id))
		req.Status = toStatus
		req.BatchId = batchId
		toStatusStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(req))
		if toStatus == types.SwapStatusRejected {
			total, found := refunds[req.SrcAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*req.Amount).Add(*req.Fee)
			refunds[req.SrcAddr] = total
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

func (k Keeper) ChangeStatusRequestsImprovement(ctx sdk.Context, ids []uint64, status string, batchId *uint64) ([]types.Request, error) {
	if len(ids) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap request transactions' id is empty")
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

	reqs := k.GetRequestsByIdsFromStore(ctx, currentStatusStore, ids)

	firstSrcNet := reqs[0].SrcNetwork
	firstDestNet := reqs[0].DestNetwork

	var isSwapOut bool
	switch firstSrcNet {
	case types.NetworkNameShareLedger:
		//swap out
		isSwapOut = true
		err := k.swapOut(ctx, status, reqs)
		if err != nil {
			return nil, err
		}
	default:
		//swap in
		if firstDestNet == types.NetworkNameShareLedger {
			// change status to done also
			isSwapOut = false
			status = types.SwapStatusDone
			err := k.swapIn(ctx, status, reqs)
			if err != nil {
				return nil, err
			}
		}
	}

	err := k.MoveRequest(ctx, requiredStatus, status, reqs, batchId, isSwapOut)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (k Keeper) swapIn(ctx sdk.Context, stt string, reqs []types.Request) error {
	transfers := make(map[string]sdk.DecCoins)
	for i := range reqs {

		if stt == types.SwapStatusDone {
			total, found := transfers[reqs[i].DestAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*reqs[i].Amount)
			transfers[reqs[i].DestAddr] = total
		}
	}
	for a, c := range transfers {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bc, err := denom.NormalizeToBaseCoins(c, false)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}

		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, bc); err != nil {
			return err
		}
	}
	return nil
}
func (k Keeper) swapOut(ctx sdk.Context, stt string, reqs []types.Request) error {

	refunds := make(map[string]sdk.DecCoins)
	for i := range reqs {
		if stt == types.SwapStatusRejected {
			total, found := refunds[reqs[i].SrcAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*reqs[i].Amount).Add(*reqs[i].Fee)
			refunds[reqs[i].SrcAddr] = total
		}
	}
	for a, c := range refunds {
		add, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
		}
		bc, err := denom.NormalizeToBaseCoins(c, false)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, add, bc); err != nil {
			return err
		}
	}
	return nil
}
func (k Keeper) changeStatusSwapOutImprovement(ctx sdk.Context, fromStatus string, toStatus string, reqs []types.Request, batchId *uint64) ([]types.Request, error) {

	if toStatus == types.SwapStatusApproved && (*batchId == 0) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s status requires batch Id", toStatus)
	}

	err := k.MoveRequest(ctx, fromStatus, toStatus, reqs, batchId, true)
	if err != nil {
		return nil, err
	}

	refunds := make(map[string]sdk.DecCoins)
	for i := range reqs {
		req := reqs[i]
		if toStatus == types.SwapStatusRejected {
			total, found := refunds[req.SrcAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total = total.Add(*req.Amount).Add(*req.Fee)
			refunds[req.SrcAddr] = total
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
	return reqs, nil
}

//MoveRequest move the request to the store base on status
//Delete request form store and add this request to store. This acction also update batch ID if it's existed
//TODO consider the request validation
func (k Keeper) MoveRequest(ctx sdk.Context, fromStt, toStt string, reqs []types.Request, batchID *uint64, isOut bool) error {

	storeMap := k.GetStoreRequestMap(ctx)
	fromStore, found := storeMap[fromStt]
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic, "store not found")
	}
	toStore, found := storeMap[toStt]
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrLogic, "store not found")
	}

	if isOut {
		if batchID == nil || *batchID == 0 {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "batch id not found")
		}
	}

	destNetwork := reqs[0].DestNetwork
	for i := range reqs {
		req := &reqs[i]

		if isOut {
			if req.SrcNetwork != types.NetworkNameShareLedger {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same source network %v", req.Id, types.NetworkNameShareLedger)
			}
		} else {
			//in case swap in need to check all request is same destination network
			if req.DestNetwork != types.NetworkNameShareLedger {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same dest network %v", req.Id, types.NetworkNameShareLedger)
			}
		}
		if req.DestNetwork != destNetwork {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same destination network %v", req.Id, destNetwork)
		}
		fromStore.Delete(GetRequestIDBytes(req.Id))
		req.Status = toStt
		if batchID != nil {
			req.BatchId = *batchID
		}
		toStore.Set(GetRequestIDBytes(req.Id), k.cdc.MustMarshal(req))

	}
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
	k.Logger(ctx).Error("panic in here GetRequestFromStore:211 ")
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

// GetRequestIDFromBytes returns ID in uint64 format from a byte array
func GetRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func ValidateSrcNetwork(ctx sdk.Context, rqs []types.Request) bool {
	if len(rqs) == 0 {
		return false
	}
	src := rqs[0].SrcNetwork

	for i := range rqs {
		if rqs[i].SrcNetwork != src {
			return false
		}
	}
	return true
}
