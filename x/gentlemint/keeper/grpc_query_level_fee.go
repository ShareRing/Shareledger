package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/constant"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LevelFees(c context.Context, req *types.QueryLevelFeesRequest) (*types.QueryLevelFeesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	defaultsLevelFees := constant.DefaultFeeLevel
	storedLevelFees := k.GetAllLevelFee(ctx)
	levelFees := make([]types.LevelFeeDetail, 0, len(defaultsLevelFees)+len(storedLevelFees))
	exchangeRate := k.GetExchangeRateD(ctx)

	for _, lf := range storedLevelFees {
		decCoins, err := sdk.ParseDecCoins(lf.Fee)
		if err != nil {
			return nil, err
		}
		convertedFee, err := denom.NormalizeCoins(decCoins, &exchangeRate)
		if err != nil {
			return nil, err
		}
		levelFees = append(levelFees, types.LevelFeeDetail{
			Level:        lf.Level,
			OriginalFee:  lf.Fee,
			ConvertedFee: &convertedFee,
			Creator:      lf.Creator,
		})
		delete(defaultsLevelFees, constant.DefaultLevel(lf.Level))
	}
	for l, f := range defaultsLevelFees {
		decCoins := sdk.NewDecCoins(f)

		convertedFee, err := denom.NormalizeCoins(decCoins, &exchangeRate)
		if err != nil {
			return nil, err
		}

		levelFees = append(levelFees, types.LevelFeeDetail{
			Level:        string(l),
			OriginalFee:  f.String(),
			ConvertedFee: &convertedFee,
		})
	}
	sort.Slice(levelFees, func(i, j int) bool {
		return levelFees[i].Level < levelFees[j].Level
	})

	return &types.QueryLevelFeesResponse{LevelFees: levelFees}, nil
}

func (k Keeper) LevelFee(c context.Context, req *types.QueryLevelFeeRequest) (*types.QueryLevelFeeResponse, error) {
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

	return &types.QueryLevelFeeResponse{LevelFee: val}, nil
}
