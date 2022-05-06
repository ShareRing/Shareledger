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

func (k Keeper) SignSchemaAll(c context.Context, req *types.QueryAllSignSchemasRequest) (*types.QueryAllSignSchemasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var formats []types.SignSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	formatStore := prefix.NewStore(store, types.KeyPrefix(types.FormatKeyPrefix))

	pageRes, err := query.Paginate(formatStore, req.Pagination, func(key []byte, value []byte) error {
		var format types.SignSchema
		if err := k.cdc.Unmarshal(value, &format); err != nil {
			return err
		}

		formats = append(formats, format)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignSchemasResponse{Schemas: formats, Pagination: pageRes}, nil
}

func (k Keeper) SignSchema(c context.Context, req *types.QueryGetSignSchemaRequest) (*types.QuerySignSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFormat(
		ctx,
		req.Network,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QuerySignSchemaResponse{SignSchema: val}, nil
}
