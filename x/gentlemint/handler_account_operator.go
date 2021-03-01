package gentlemint

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgEnrollAccountOperator(ctx sdk.Context, keeper Keeper, msg types.MsgEnrollAccOperators) (*sdk.Result, error) {
	if err := checkAuthority(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(
		types.EventTypeEnrollAccOp,
	)

	for _, addr := range msg.Operators {
		acc := types.NewAccState(addr, types.Active)
		keeper.SetAccOp(ctx, acc)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))

	}
	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			// sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			// sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollAccOp),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgRevokeAccountOperator(ctx sdk.Context, keeper Keeper, msg types.MsgRevokeAccOperators) (*sdk.Result, error) {
	if err := checkAuthority(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeRevokeAccOp)

	// Only deactivate active accounts
	for _, addr := range msg.Operators {
		acc := keeper.GetAccOp(ctx, addr)
		if acc.Status != types.Active {
			return nil, types.ErrDoesNotExist
		}

		acc.Status = types.Inactive
		keeper.SetAccOp(ctx, acc)

		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeAccOp),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func IsAccountOperator(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {

	acc := k.GetAccOp(ctx, address)
	if acc.Status != types.Active {
		return false
	}

	return true
}

func checkAuthority(ctx sdk.Context, address sdk.AccAddress, k Keeper) error {
	if !IsAuthority(ctx, address, k) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAuthority)
	}
	return nil
}

func CheckAccountOperator(ctx sdk.Context, address sdk.AccAddress, k Keeper) error {
	if !IsAccountOperator(ctx, address, k) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAccountOperator)
	}
	return nil
}
