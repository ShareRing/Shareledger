package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) AllSchemas(c context.Context, req *types.QueryAllSchemasRequest) (*types.QueryAllSchemasResponse, error) {
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
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
	}

	return &types.QueryAllSchemasResponse{Schemas: schemas, Pagination: pageRes}, nil
}

func (k Keeper) Schema(c context.Context, req *types.QueryGetSchemaRequest) (*types.QuerySchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSchema(
		ctx,
		req.Network,
	)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s does not exist in schema", req.Network)
	}

	return &types.QuerySchemaResponse{Schema: val}, nil
}
