package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetLoader(goCtx context.Context, req *types.QueryGetLoaderRequest) (*types.QueryGetLoaderResponse, error) {
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

	return &types.QueryGetLoaderResponse{
		AccState: &types.AccState{
			Key:     m.Key,
			Address: m.Address,
			Status:  m.Status,
		},
	}, nil
}
