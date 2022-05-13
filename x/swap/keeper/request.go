package keeper

import (
	"encoding/binary"
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

// ChangeStatusRequests change status of requests and move it into respective store
// The status flow should be: pending -> approved|rejected -> processing -> done.
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
	var currentStatusStore prefix.Store
	switch status {
	case types.SwapStatusApproved, types.SwapStatusRejected:
		requiredStatus = types.SwapStatusPending
	case types.BatchStatusPending: // the request just gone to pending when this status was approved. In case we cancel the swap batch
		requiredStatus = types.SwapStatusApproved
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is not supported", status)
	}

	currentStatusStore = k.GetStoreRequestMap(ctx)[requiredStatus]

	reqs := k.GetRequestsByIdsFromStore(ctx, currentStatusStore, ids)
	if len(reqs) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request not found")
	}
	destNet := reqs[0].DestNetwork
	for i := range reqs {
		//source network swap out case must is slp3
		if isSwapOut {
			if reqs[i].SrcNetwork != types.NetworkNameShareLedger {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same source network %v", reqs[i].Id, types.NetworkNameShareLedger)
			}
			if reqs[i].DestNetwork != destNet {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same dest network %v", reqs[i].Id, destNet)
			}
		} else {
			if reqs[i].DestNetwork != types.NetworkNameShareLedger {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "swap transaction with %v does not have same destination network %v", reqs[i].Id, types.NetworkNameShareLedger)
			}
		}

	}

	if isSwapOut && status == types.SwapStatusApproved {
		b, f := k.GetBatch(ctx, *batchId)
		if !f {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "fail to set batch network")
		}
		b.Network = destNet
		k.SetBatch(ctx, b)
	}

	if !isSwapOut && status == types.SwapStatusApproved {
		status = types.SwapStatusDone
	}
	err := k.MoveRequest(ctx, requiredStatus, status, reqs, batchId, isSwapOut)
	if err != nil {
		return nil, err
	}

	bankTransferCoins := make(map[string]sdk.DecCoins)

	if isSwapOut {
		bankTransferCoins = k.swapOut(ctx, status, reqs)
	} else {
		bankTransferCoins = k.swapIn(ctx, status, reqs)
	}

	err = k.SendCoinToAddr(ctx, bankTransferCoins)
	if err != nil {
		return nil, err
	}

	return reqs, nil
}

func (k Keeper) swapIn(_ sdk.Context, stt string, reqs []types.Request) map[string]sdk.DecCoins {

	transfers := make(map[string]sdk.DecCoins)
	for i := range reqs {
		if stt == types.SwapStatusDone {
			total, found := transfers[reqs[i].DestAddr]
			if !found {
				total = sdk.NewDecCoins()
			}
			total.Add(sdk.NewDecCoinFromCoin(*reqs[i].GetAmount()))

			transfers[reqs[i].DestAddr] = total
		}
	}
	return transfers
}
func (k Keeper) swapOut(_ sdk.Context, stt string, reqs []types.Request) map[string]sdk.DecCoins {

	refunds := make(map[string]sdk.DecCoins)
	for i := range reqs {
		if stt == types.SwapStatusRejected {
			total, found := refunds[reqs[i].SrcAddr]
			if !found {
				total = sdk.NewDecCoins()
			}

			total = total.Add(sdk.NewDecCoinFromCoin(*reqs[i].GetAmount())).Add(sdk.NewDecCoinFromCoin(*reqs[i].GetFee()))
			refunds[reqs[i].SrcAddr] = total
		}
	}

	return refunds
}

//MoveRequest move the request to the store base on status
//Delete request form store and add this request to store. This action also update batch ID if it's existed
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

	for i := range reqs {
		req := &reqs[i]
		fromStore.Delete(GetRequestIDBytes(req.Id))
		req.Status = toStt
		// just needing the batch ID in case of approve swap out
		if isOut && batchID != nil && toStt == types.SwapStatusApproved {
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

//SendCoinToAddr use bank module to and send out coin to the shareLedger address from module address
func (k Keeper) SendCoinToAddr(ctx sdk.Context, c map[string]sdk.DecCoins) error {

	for address, coin := range c {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%+v", err)
		}
		bc, err := denom.NormalizeToBaseCoins(coin, false)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "%v", err)
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, bc); err != nil {
			return err
		}
	}

	return nil
}
