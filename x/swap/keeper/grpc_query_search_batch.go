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

func (k Keeper) SearchBatch(goCtx context.Context, req *types.QuerySearchBatchRequest) (*types.QuerySearchBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	batchStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BatchKey))

	var batches []types.Batch

	pageRes, err := query.FilteredPaginate(batchStore, req.GetPagination(), func(key []byte, value []byte, accumulate bool) (bool, error) {
		var val types.Batch
		if err := k.cdc.Unmarshal(value, &val); err != nil {
			return false, err
		}
		if val.Status != req.GetStatus() {
			return false, nil
		}

		if accumulate {
			batches = append(batches, val)
		}
		return true, nil
	})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "getting batchs fail")
	}

	return &types.QuerySearchBatchResponse{
		Batch:      batches,
		Pagination: pageRes,
	}, nil
}
