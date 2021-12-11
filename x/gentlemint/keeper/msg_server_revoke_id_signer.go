package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevokeIdSigner(goCtx context.Context, msg *types.MsgRevokeIdSigner) (*types.MsgRevokeIdSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !k.isAccountOperator(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAccountOperator)
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
	return &types.MsgRevokeIdSignerResponse{}, nil
}