package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CheckFees(goCtx context.Context, req *types.QueryCheckFeesRequest) (*types.QueryCheckFeesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var result types.QueryCheckFeesResponse

	fee := sdk.NewCoin(types.DenomSHR, sdk.NewInt(0))
	for _, a := range req.Actions {
		af := k.GetShrFeeByActionKey(ctx, a)
		fee = fee.Add(af)
	}
	result.ShrFee = fee.String()

	currentBalances := k.bankKeeper.GetAllBalances(ctx, addr)
	currentShr := sdk.NewCoin(types.DenomSHR, currentBalances.AmountOf(types.DenomSHR))
	result.SufficientFee = currentShr.IsGTE(fee)
	result.SufficientFundForFee = result.SufficientFee // sufficient fee is true, sufficient fund for fee will be true by default
	if !result.SufficientFee {
		rate := k.GetExchangeRateD(ctx)
		currentShrFromShrp := types.ShrpToShr(
			sdk.NewCoins(
				sdk.NewCoin(types.DenomSHRP, currentBalances.AmountOf(types.DenomSHRP)),
				sdk.NewCoin(types.DenomCent, currentBalances.AmountOf(types.DenomCent)),
			),
			rate,
		)
		neededShr := fee.Sub(currentShr)

		result.SufficientFundForFee = currentShrFromShrp.IsGTE(neededShr)

		if result.SufficientFundForFee {
			result.ShrpCostLoadingFee = types.ShrToShrp(neededShr, rate).String()
		}
	}

	return &result, nil
}
