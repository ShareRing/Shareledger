package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SwapManagers(goCtx context.Context, req *types.QuerySwapManagersRequest) (*types.QuerySwapManagersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	list := k.IterateAccState(ctx, types.AccStateKeySwapManager)
	swapManager := make([]*types.AccState, 0, len(list))
	for _, i := range list {
		swapManager = append(swapManager, &types.AccState{
			Key:     i.Key,
			Address: i.Address,
			Status:  i.Status,
		})
	}

	return &types.QuerySwapManagersResponse{SwapManagers: swapManager}, nil
}
