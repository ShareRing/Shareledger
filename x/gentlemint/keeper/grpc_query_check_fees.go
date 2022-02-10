package keeper

import (
	"context"
	denom "github.com/sharering/shareledger/x/utils/demo"

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

	fee := sdk.NewCoin(denom.PShr, sdk.NewInt(0))
	for _, a := range req.Actions {
		af := k.GetPShrFeeByActionKey(ctx, a)
		fee = fee.Add(af)
	}
	result.ConvertedFee = &fee

	currentBalances := k.bankKeeper.GetAllBalances(ctx, addr)
	currentShr := sdk.NewCoin(denom.PShr, currentBalances.AmountOf(denom.PShr))
	result.SufficientFee = currentShr.IsGTE(fee)
	result.SufficientFundForFee = result.SufficientFee // sufficient fee is true, sufficient fund for fee will be true by default
	if !result.SufficientFee {
		rate := k.GetExchangeRateD(ctx)
		calculatedPShrFromShrpFund := denom.NormalizeCoins(
			sdk.NewDecCoinsFromCoins(sdk.NewCoins(
				sdk.NewCoin(denom.ShrP, currentBalances.AmountOf(denom.ShrP)),
				sdk.NewCoin(denom.Cent, currentBalances.AmountOf(denom.Cent)),
			)...), rate)

		// Should check for the whole not partial fee to avoid a case that:
		// User have enough token to send out but not enough fee. So we need to buy whole fee token to let user be able to send out their current balance.
		result.SufficientFundForFee = calculatedPShrFromShrpFund.IsGTE(fee)

		if result.SufficientFundForFee {
			dC := denom.ToDecShrPCoin(sdk.NewDecCoinsFromCoins(fee), rate)
			result.CostLoadingFee = &dC
		}
	}

	return &result, nil
}
