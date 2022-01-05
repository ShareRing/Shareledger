package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) BuyShr(goCtx context.Context, msg *types.MsgBuyShr) (*types.MsgBuyShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	coins, err := types.ParseShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}
	if err := k.buyShr(ctx, coins.AmountOf(types.DenomSHR), msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "buy %v shr to %v", msg.Amount, msg.Creator)
	}
	return &types.MsgBuyShrResponse{
		Log: fmt.Sprintf("Successfull buy %v shr for address %s", msg.Amount, msg.Creator),
	}, nil
}

func (k msgServer) buyShr(ctx sdk.Context, amount sdk.Int, buyer sdk.AccAddress) error {
	if !k.ShrMintPossible(ctx, amount) {
		return sdkerrors.Wrap(types.ErrSHRSupplyExceeded, amount.String())
	}

	rate := k.GetExchangeRateF(ctx)

	currentBalance := k.bankKeeper.GetAllBalances(ctx, buyer)
	currentShrpBalance := sdk.NewCoins(
		sdk.NewCoin(types.DenomSHRP, currentBalance.AmountOf(types.DenomSHRP)),
		sdk.NewCoin(types.DenomCent, currentBalance.AmountOf(types.DenomCent)),
	)

	cost, err := types.GetCostShrpForShr(currentShrpBalance, amount, rate)
	if err != nil {
		return sdkerrors.Wrapf(err, "current %v balance", currentShrpBalance)
	}
	if cost.Sub.Empty() {
		return sdkerrors.ErrInsufficientFunds
	}

	if !cost.Add.Empty() {
		if err := k.loadCoins(ctx, buyer, cost.Add); err != nil {
			return sdkerrors.Wrapf(err, "%v coins in return", cost.Add)
		}
	}
	if err := k.burnCoins(ctx, buyer, cost.Sub); err != nil {
		return sdkerrors.Wrapf(err, "charge %v coins", cost.Sub)
	}
	boughtShr := sdk.NewCoins(sdk.NewCoin(types.DenomSHR, amount))
	if err := k.loadCoins(ctx, buyer, boughtShr); err != nil {
		return sdkerrors.Wrapf(err, "send %v coins", boughtShr)
	}
	return nil
}
