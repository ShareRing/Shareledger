package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	fee := sdk.NewCoin(denom.Base, sdk.NewInt(0))
	for _, a := range req.Actions {
		af, err := k.GetBaseDenomFeeByActionKey(ctx, a)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "get %v fee by action %v", denom.Base, a)
		}
		fee = fee.Add(af)
	}
	result.ConvertedFee = &fee

	currentBalances := k.bankKeeper.GetAllBalances(ctx, addr)
	currentShr := sdk.NewCoin(denom.Base, currentBalances.AmountOf(denom.Base))
	result.SufficientFee = currentShr.IsGTE(fee)
	result.SufficientFundForFee = result.SufficientFee // sufficient fee is true, sufficient fund for fee will be true by default
	if !result.SufficientFee {
		rate := k.GetExchangeRateD(ctx)
		calculatedBaseDenomFromShrpFund, err := denom.NormalizeToBaseCoin(denom.Base,
			sdk.NewDecCoinsFromCoins(sdk.NewCoins(
				sdk.NewCoin(denom.ShrP, currentBalances.AmountOf(denom.ShrP)),
				sdk.NewCoin(denom.BaseUSD, currentBalances.AmountOf(denom.BaseUSD)),
			)...), rate, true)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "calculate %v from %v fund %v", denom.Base, denom.ShrP, currentBalances)
		}
		// Should check for the whole not partial fee to avoid a case that:
		// User have enough token to send out but not enough fee. So we need to buy whole fee token to let user be able to send out their current balance.
		result.SufficientFundForFee = calculatedBaseDenomFromShrpFund.IsGTE(fee)

		if result.SufficientFundForFee {
			dC, err := denom.To(sdk.NewDecCoinsFromCoins(fee), denom.ShrP, rate)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "convert fee,%v, to usd", fee)
			}
			result.CostLoadingFee = &dC
		}
	}

	return &result, nil
}
