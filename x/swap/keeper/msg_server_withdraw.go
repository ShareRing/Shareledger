package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.Amount), true)
	if err != nil {
		k.Logger(ctx).Error("normalizer the base_coins fail", "error", err.Error())
		return &types.MsgWithdrawResponse{Status: types.TxnStatusFail}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "input coin for withdraw is invalid %s", err)
	}

	recAddr, err := sdk.AccAddressFromBech32(msg.GetReceiver())
	if err != nil {
		k.Logger(ctx).Error("getting receiver address fail", "error", err)
		return &types.MsgWithdrawResponse{Status: types.TxnStatusFail}, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "receiver address is invalid %s", err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recAddr, baseCoins)
	if err != nil {
		k.Logger(ctx).Error("withdraw fail", "error", err)
		return &types.MsgWithdrawResponse{Status: types.TxnStatusFail}, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithDraw,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeWithdrawReceiver, msg.GetReceiver()),
			sdk.NewAttribute(types.EventTypeDepositAmount, msg.GetAmount().String()),
		),
	)
	return &types.MsgWithdrawResponse{Status: types.TxnStatusSuccess}, nil
}
