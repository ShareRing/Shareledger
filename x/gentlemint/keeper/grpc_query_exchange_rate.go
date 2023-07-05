package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k Keeper) ExchangeRate(c context.Context, req *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExchangeRate(ctx)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryExchangeRateResponse{Rate: val.Rate}, nil
}
