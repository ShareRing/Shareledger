package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeAccountOperators(goCtx context.Context, msg *types.MsgRevokeAccountOperators) (*types.MsgRevokeAccountOperatorsResponse, error) {
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
	return &types.MsgRevokeAccountOperatorsResponse{}, nil
}
