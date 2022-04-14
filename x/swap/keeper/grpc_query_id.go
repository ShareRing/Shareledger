package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdAll(c context.Context, req *types.QueryAllIdRequest) (*types.QueryAllIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ids []types.Id
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	idStore := prefix.NewStore(store, types.KeyPrefix(types.IdKeyPrefix))

	pageRes, err := query.Paginate(idStore, req.Pagination, func(key []byte, value []byte) error {
		var id types.Id
		if err := k.cdc.Unmarshal(value, &id); err != nil {
			return err
		}

		ids = append(ids, id)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIdResponse{Id: ids, Pagination: pageRes}, nil
}

func (k Keeper) Id(c context.Context, req *types.QueryGetIdRequest) (*types.QueryGetIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetId(
		ctx,
		req.IDType,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIdResponse{Id: val}, nil
}
