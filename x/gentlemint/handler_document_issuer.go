package gentlemint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func handleMsgEnrollDocummentIssuer(ctx sdk.Context, keeper Keeper, msg types.MsgEnrollDocIssuers) (*sdk.Result, error) {
	if err := CheckAccountOperator(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeEnrollDocIssuer)

	for _, addr := range msg.Issuers {
		acc := types.NewAccState(addr, types.Active)
		keeper.SetDocIssuer(ctx, acc)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollDocIssuer),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgRevokeDocumentIssuers(ctx sdk.Context, keeper Keeper, msg types.MsgRevokeDocIssuers) (*sdk.Result, error) {
	if err := CheckAccountOperator(ctx, msg.GetSigners()[0], keeper); err != nil {
		return nil, err
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeRevokeDocIssuer)

	for _, addr := range msg.Issuers {
		acc := keeper.GetDocIssuer(ctx, addr)
		if acc.Status != types.Active {
			return nil, types.ErrDoesNotExist
		}

		acc.Status = types.Inactive
		keeper.SetDocIssuer(ctx, acc)

		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	// Emit event
	ctx.EventManager().EmitEvent(event)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Approver.String()),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeDocIssuer),
		),
	)

	return &sdk.Result{
		Events: ctx.EventManager().Events(),
	}, nil
}

func IsDocIssuer(ctx sdk.Context, address sdk.AccAddress, k Keeper) bool {

	acc := k.GetDocIssuer(ctx, address)
	if acc.Status != types.Active {
		return false
	}

	return true
}
