package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/constant"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActionLevelFeeAll(c context.Context, req *types.QueryAllActionLevelFeeRequest) (*types.QueryAllActionLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionLevelFees []types.ActionLevelFee
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionLevelFeeStore := prefix.NewStore(store, types.KeyPrefix(types.ActionLevelFeeKeyPrefix))

	pageRes, err := query.Paginate(actionLevelFeeStore, req.Pagination, func(key []byte, value []byte) error {
		var actionLevelFee types.ActionLevelFee
		if err := k.cdc.Unmarshal(value, &actionLevelFee); err != nil {
			return err
		}

		actionLevelFees = append(actionLevelFees, actionLevelFee)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionLevelFeeResponse{ActionLevelFee: actionLevelFees, Pagination: pageRes}, nil
}

func (k Keeper) ActionLevelFee(c context.Context, req *types.QueryGetActionLevelFeeRequest) (*types.QueryGetActionLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	defaultLevel := string(constant.MinFee)

	val, found := k.GetActionLevelFee(
		ctx,
		req.Action,
	)
	if found {
		defaultLevel = val.Level
	}

	return &types.QueryGetActionLevelFeeResponse{
		Action: val.Action,
		Level:  defaultLevel,
		Fee:    k.GetFeeByLevel(ctx, defaultLevel).String(),
	}, nil
}
