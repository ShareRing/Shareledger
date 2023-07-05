package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) NextBatchId(goCtx context.Context, req *types.QueryNextBatchIdRequest) (*types.QueryNextBatchIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	batchCount := k.GetBatchCount(ctx)

	return &types.QueryNextBatchIdResponse{NextCount: batchCount}, nil
}
