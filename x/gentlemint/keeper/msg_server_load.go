package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) Load(goCtx context.Context, msg *types.MsgLoad) (*types.MsgLoadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	destAdd, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	baseCoins, err := denom.NormalizeToBaseCoins(msg.Coins, false)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, err.Error())
	}
	if !k.BaseMintPossible(ctx, baseCoins.AmountOf(denom.Base)) {
		return nil, sdkerrors.Wrapf(types.ErrBaseSupplyExceeded, "load %v", baseCoins)
	}
	if err := k.loadCoins(ctx, destAdd, baseCoins); err != nil {
		return nil, err
	}

	// Pay fee for loader who is a creator of this transaction
	if msg.Creator != msg.Address {
		exchangeRate := k.GetExchangeRateD(ctx)
		loadDFee := k.GetFeeByMsg(ctx, msg)
		loadFee, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(loadDFee), exchangeRate, true)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, err.Error())
		}
		currentBalance := k.bankKeeper.GetBalance(ctx, destAdd, denom.Base)
		if currentBalance.IsLT(loadFee) {
			if err := k.buyBaseDenom(ctx, loadFee.Sub(currentBalance), destAdd); err != nil {
				return nil, err
			}
		}
		if err := k.bankKeeper.SendCoins(ctx, destAdd, msg.GetSigners()[0], sdk.NewCoins(loadFee)); err != nil {
			return nil, sdkerrors.Wrapf(err, "pay fee, %v, to approver, %v", loadFee, msg.Creator)
		}
	}

	return &types.MsgLoadResponse{}, nil
}
