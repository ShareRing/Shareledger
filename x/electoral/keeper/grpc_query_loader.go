package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/electoral/types"
)

func (k Keeper) Loader(goCtx context.Context, req *types.QueryLoaderRequest) (*types.QueryLoaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	key := types.GenAccStateIndexKey(addr, types.AccStateKeyShrpLoaders)
	m, f := k.GetAccState(ctx, key)
	if !f {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryLoaderResponse{
		AccState: &types.AccState{
			Key:     m.Key,
			Address: m.Address,
			Status:  m.Status,
		},
	}, nil
}
