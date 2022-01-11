package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/constant"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LevelFeeAll(c context.Context, req *types.QueryAllLevelFeeRequest) (*types.QueryAllLevelFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	defaultsLevelFees := constant.DefaultFeeLevel
	storedLevelFees := k.GetAllLevelFee(ctx)
	levelFees := make([]types.LevelFee, 0, len(defaultsLevelFees)+len(storedLevelFees))
	for _, lf := range storedLevelFees {
		levelFees = append(levelFees, types.LevelFee{
			Level:   lf.Level,
			Fee:     lf.Fee,
			Creator: lf.Creator,
		})
		delete(defaultsLevelFees, constant.DefaultLevel(lf.Level))
	}
	for l, f := range defaultsLevelFees {
		levelFees = append(levelFees, types.LevelFee{
			Level: string(l),
			Fee:   f.String(),
		})
	}
	sort.Slice(levelFees, func(i, j int) bool {
		return levelFees[i].Level < levelFees[j].Level
	})

	return &types.QueryAllLevelFeeResponse{LevelFee: levelFees}, nil
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
