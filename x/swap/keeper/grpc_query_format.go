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

func (k Keeper) FormatAll(c context.Context, req *types.QueryAllFormatRequest) (*types.QueryAllFormatResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var formats []types.Format
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	formatStore := prefix.NewStore(store, types.KeyPrefix(types.FormatKeyPrefix))

	pageRes, err := query.Paginate(formatStore, req.Pagination, func(key []byte, value []byte) error {
		var format types.Format
		if err := k.cdc.Unmarshal(value, &format); err != nil {
			return err
		}

		formats = append(formats, format)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFormatResponse{Format: formats, Pagination: pageRes}, nil
}

func (k Keeper) Format(c context.Context, req *types.QueryGetFormatRequest) (*types.QueryGetFormatResponse, error) {
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

	return &types.QueryGetFormatResponse{Format: val}, nil
}