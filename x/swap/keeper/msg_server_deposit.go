package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(*msg.Amount), true)
	if err != nil {
		k.Logger(ctx).Error("normalizer the base_coins fail", "error", err.Error())
		return &types.MsgDepositResponse{Status: types.TxnStatusFail}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%+v", err)
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, baseCoins)
	if err != nil {
		k.Logger(ctx).Error("sending coin to swapping module fail", "error", err.Error())
		return &types.MsgDepositResponse{Status: types.TxnStatusFail}, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeDepositAddr, msg.GetSigners()[0].String()),
			sdk.NewAttribute(types.EventTypeDepositAmount, msg.GetAmount().String()),
		),
	)

	return &types.MsgDepositResponse{Status: types.TxnStatusSuccess}, nil
}
