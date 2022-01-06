package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeIdSigners(goCtx context.Context, msg *types.MsgRevokeIdSigners) (*types.MsgRevokeIdSignersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	event := sdk.NewEvent(types.EventTypeRevokeIdSigner)
	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		if err := k.revokeIdSigner(ctx, addr); err != nil {
			return nil, err
		}
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, a))
	}
	events := []sdk.Event{
		event,
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeIdSigner),
		),
	}
	ctx.EventManager().EmitEvents(events)

	return &types.MsgRevokeIdSignersResponse{}, nil
}
