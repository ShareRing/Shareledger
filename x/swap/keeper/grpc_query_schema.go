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

func (k Keeper) AllSchemas(c context.Context, req *types.QueryAllSchemasRequest) (*types.QueryAllSchemasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var schemas []types.Schema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.SignSchemaKeyPrefix))

	pageRes, err := query.Paginate(signSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var signSchema types.Schema
		if err := k.cdc.Unmarshal(value, &signSchema); err != nil {
			return err
		}

		schemas = append(schemas, signSchema)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSchemasResponse{Schemas: schemas, Pagination: pageRes}, nil
}

func (k Keeper) Schema(c context.Context, req *types.QueryGetSchemaRequest) (*types.QuerySchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSchema(
		ctx,
		req.Network,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QuerySchemaResponse{Schema: val}, nil
}
