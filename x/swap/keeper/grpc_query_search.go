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

func (k Keeper) Search(goCtx context.Context, req *types.QuerySearchRequest) (*types.QuerySearchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := req.BasicValidate(); err != nil {
		return nil, err
	}
	store := k.GetStoreRequestMap(ctx)[req.Status]
	var requests []types.Request

	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var val types.Request
		if err := k.cdc.Unmarshal(value, &val); err != nil {
			return false, err
		}
		if req.Id != 0 && req.Id != val.Id {
			return false, nil
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
			requests = append(requests, val)
		}

		return true, nil
	})

	return &types.QuerySearchResponse{
		Request:    requests,
		Pagination: pageRes,
	}, err
}
