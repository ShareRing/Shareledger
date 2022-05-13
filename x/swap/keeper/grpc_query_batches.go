package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Batches(goCtx context.Context, req *types.QueryBatchesRequest) (*types.QueryBatchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	bMap := make(map[uint64]struct{})
	for i := range req.GetIds() {
		bMap[req.Ids[i]] = struct{}{}
	}
	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))

	var batches []types.Batch

	pageRes, err := query.FilteredPaginate(batchStore, req.GetPagination(), func(key []byte, value []byte, accumulate bool) (bool, error) {
		var val types.Batch
		if err := k.cdc.Unmarshal(value, &val); err != nil {
			return false, err
		}
		if req.GetStatus() != "" {
			if val.Status != req.GetStatus() {
				return false, nil
			}
		}

		if req.GetNetwork() != "" {
			if val.Network != req.GetNetwork() {
				return false, nil
			}
		}
		if len(req.GetIds()) > 0 {
			_, matched := bMap[val.Id]
			if !matched {
				return false, nil
			}
		}

		if accumulate {
			batches = append(batches, val)
		}
		return true, nil
	})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "getting batches fail")
	}

	return &types.QueryBatchesResponse{
		Batches:    batches,
		Pagination: pageRes,
	}, nil
}
