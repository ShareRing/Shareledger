package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LevelFeeAll(c context.Context, req *types.QueryAllLevelFeeRequest) (*types.QueryAllLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var levelFees []types.LevelFee
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	levelFeeStore := prefix.NewStore(store, types.KeyPrefix(types.LevelFeeKeyPrefix))

	pageRes, err := query.Paginate(levelFeeStore, req.Pagination, func(key []byte, value []byte) error {
		var levelFee types.LevelFee
		if err := k.cdc.Unmarshal(value, &levelFee); err != nil {
			return err
		}

		levelFees = append(levelFees, levelFee)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLevelFeeResponse{LevelFee: levelFees, Pagination: pageRes}, nil
}

func (k Keeper) LevelFee(c context.Context, req *types.QueryGetLevelFeeRequest) (*types.QueryGetLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLevelFee(
	    ctx,
	    req.Level,
        )
	if !found {
	    return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetLevelFeeResponse{LevelFee: val}, nil
}