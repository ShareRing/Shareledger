package gentlemint

import (
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgEnrollIdSigners(ctx sdk.Context, keeper Keeper, msg MsgEnrollIDSigners) (*sdk.Result, error) {
	if err := CheckAccountOperator(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeEnrollIdSigner)
	for _, addr := range msg.IDSigners {
		idSigner := types.NewAccState(addr, types.Active)
		keeper.SetIdSigner(ctx, idSigner)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollIdSigner),
		),
	)
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgRevokeIdSigners(ctx sdk.Context, keeper Keeper, msg MsgRevokeIDSigners) (*sdk.Result, error) {
	if err := CheckAccountOperator(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeRevokeIdSigner)
	for _, addr := range msg.IDSigners {
		acc := keeper.GetIdSigner(ctx, addr)
		if acc.Status != types.Active {
			return nil, types.ErrDoesNotExist
		}

		acc.Status = types.Inactive
		keeper.SetIdSigner(ctx, acc)

		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeIdSigner),
		),
	)
	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func IsIdSigner(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {

	idSigner := k.GetIdSigner(ctx, address)
	if idSigner.Status != types.Active {
		return false
	}

	return true
}

func checkIdSigner(ctx sdk.Context, address sdk.AccAddress, k Keeper) error {
	if !IsIdSigner(ctx, address, k) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotIssuer)
	}
	return nil
}
