package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeDocIssuers(goCtx context.Context, msg *types.MsgRevokeDocIssuers) (*types.MsgRevokeDocIssuersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	event := sdk.NewEvent(types.EventTypeRevokeDocIssuer)
	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		if err := k.revokeDocIssuer(ctx, addr); err != nil {
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
				sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeDocIssuer),
			),
		},
	)

	return &types.MsgRevokeDocIssuersResponse{}, nil
}
