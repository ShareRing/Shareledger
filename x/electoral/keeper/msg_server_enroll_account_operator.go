package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollAccountOperators(goCtx context.Context, msg *types.MsgEnrollAccountOperators) (*types.MsgEnrollAccountOperatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(
		types.EventTypeEnrollAccOp,
	)
	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeAccOperator(ctx, addr)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, a))
	}
	ctx.EventManager().EmitEvents([]sdk.Event{
		event,
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollAccOp),
		),
	},
	)

	return &types.MsgEnrollAccountOperatorsResponse{}, nil
}
