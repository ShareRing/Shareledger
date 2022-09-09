package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/query"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Swap(goCtx context.Context, req *types.QuerySwapRequest) (*types.QuerySwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.BasicValidate(); err != nil {
		return nil, err
	}
	store := k.GetStoreRequestMap(ctx)[req.Status]
	var swapRequests []types.Request
	filterIds := make(map[uint64]struct{})
	var filteredIds bool
	if len(req.Ids) > 0 {
		for _, id := range req.Ids {
			filterIds[id] = struct{}{}
		}
		filteredIds = true
	}
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var val types.Request
		if err := k.cdc.Unmarshal(value, &val); err != nil {
			return false, err
		}
		if filteredIds {
			if _, found := filterIds[val.Id]; !found {
				return false, nil
			}
		}
		if req.DestAddr != "" && !strings.EqualFold(req.DestAddr, val.DestAddr) {
			return false, nil
		}
		if req.SrcAddr != "" && !strings.EqualFold(req.SrcAddr, val.SrcAddr) {
			return false, nil
		}
		if req.DestNetwork != "" && !strings.EqualFold(req.DestNetwork, val.DestNetwork) {
			return false, nil
		}
		if req.SrcNetwork != "" && !strings.EqualFold(req.SrcNetwork, val.SrcNetwork) {
			return false, nil
		}

		if accumulate {
			swapRequests = append(swapRequests, val)
		}

		return true, nil
	})

	return &types.QuerySwapResponse{
		Swaps:      swapRequests,
		Pagination: pageRes,
	}, err
}
