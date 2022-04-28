package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Batch(c context.Context, req *types.QueryGetBatchRequest) (*types.QueryGetBatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	batch := k.GetBatchsByIDs(ctx, req.Ids)

	return &types.QueryGetBatchResponse{Batch: batch}, nil
}
