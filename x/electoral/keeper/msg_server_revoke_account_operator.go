package keeper

import (
	"context"

	"github.com/sharering/shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeAccountOperator(goCtx context.Context, msg *types.MsgRevokeAccountOperator) (*types.MsgRevokeAccountOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeRevokeAccOp)
	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		if err := k.revokeAccOperator(ctx, addr); err != nil {
			return nil, err
		}
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}

	ctx.EventManager().EmitEvents(
		[]sdk.Event{
			event,
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeAccOp),
			),
		},
	)
	return &types.MsgRevokeAccountOperatorResponse{}, nil
}
