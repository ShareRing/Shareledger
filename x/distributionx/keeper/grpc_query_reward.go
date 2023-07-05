package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func (k Keeper) RewardAll(c context.Context, req *types.QueryAllRewardRequest) (*types.QueryAllRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rewards []types.Reward
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rewardStore := prefix.NewStore(store, types.KeyPrefix(types.RewardKeyPrefix))

	pageRes, err := query.Paginate(rewardStore, req.Pagination, func(key []byte, value []byte) error {
		var reward types.Reward
		if err := k.cdc.Unmarshal(value, &reward); err != nil {
			return err
		}

		rewards = append(rewards, reward)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRewardResponse{Reward: rewards, Pagination: pageRes}, nil
}

func (k Keeper) Reward(c context.Context, req *types.QueryGetRewardRequest) (*types.QueryGetRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetReward(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRewardResponse{Reward: val}, nil
}
