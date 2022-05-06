package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollApprovers(goCtx context.Context, msg *types.MsgEnrollApprovers) (*types.MsgEnrollApproversResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	event := sdk.NewEvent(types.EventTypeEnrollApprover)

	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeApprover(ctx, addr)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}
	events := []sdk.Event{
		event,
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollApprover),
		),
	}
	ctx.EventManager().EmitEvents(events)
	return &types.MsgEnrollApproversResponse{}, nil
}
