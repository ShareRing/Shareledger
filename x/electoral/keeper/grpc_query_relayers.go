package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/electoral/types"
)

func (k Keeper) Relayers(goCtx context.Context, req *types.QueryRelayersRequest) (*types.QueryRelayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	list := k.IterateAccState(ctx, types.AccStateKeyRelayer)
	relayers := make([]*types.AccState, 0, len(list))
	for _, i := range list {
		relayers = append(relayers, &types.AccState{
			Key:     i.Key,
			Address: i.Address,
			Status:  i.Status,
		})
	}

	return &types.QueryRelayersResponse{Relayers: relayers}, nil
}
