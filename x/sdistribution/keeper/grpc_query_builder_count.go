package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/sdistribution/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BuilderCountAll(c context.Context, req *types.QueryAllBuilderCountRequest) (*types.QueryAllBuilderCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var builderCounts []types.BuilderCount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	builderCountStore := prefix.NewStore(store, types.KeyPrefix(types.BuilderCountKeyPrefix))

	pageRes, err := query.Paginate(builderCountStore, req.Pagination, func(key []byte, value []byte) error {
		var builderCount types.BuilderCount
		if err := k.cdc.Unmarshal(value, &builderCount); err != nil {
			return err
		}

		builderCounts = append(builderCounts, builderCount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBuilderCountResponse{BuilderCount: builderCounts, Pagination: pageRes}, nil
}

func (k Keeper) BuilderCount(c context.Context, req *types.QueryGetBuilderCountRequest) (*types.QueryGetBuilderCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBuilderCount(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBuilderCountResponse{BuilderCount: val}, nil
}
