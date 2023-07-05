package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) NextRequestId(goCtx context.Context, req *types.QueryNextRequestIdRequest) (*types.QueryNextRequestIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reqCount := k.GetRequestCount(ctx)

	return &types.QueryNextRequestIdResponse{NextCount: reqCount}, nil
}
