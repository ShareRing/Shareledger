package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func (k Keeper) BuilderListAll(c context.Context, req *types.QueryAllBuilderListRequest) (*types.QueryAllBuilderListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var builderLists []types.BuilderList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	builderListStore := prefix.NewStore(store, types.KeyPrefix(types.BuilderListKey))

	pageRes, err := query.Paginate(builderListStore, req.Pagination, func(key []byte, value []byte) error {
		var builderList types.BuilderList
		if err := k.cdc.Unmarshal(value, &builderList); err != nil {
			return err
		}

		builderLists = append(builderLists, builderList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBuilderListResponse{BuilderList: builderLists, Pagination: pageRes}, nil
}

func (k Keeper) BuilderList(c context.Context, req *types.QueryGetBuilderListRequest) (*types.QueryGetBuilderListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	builderList, found := k.GetBuilderList(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetBuilderListResponse{BuilderList: builderList}, nil
}
