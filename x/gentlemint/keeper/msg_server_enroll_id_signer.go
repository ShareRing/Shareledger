package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnrollIdSigner(goCtx context.Context, msg *types.MsgEnrollIdSigner) (*types.MsgEnrollIdSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !k.IsAccountOperator(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAccountOperator)
	}
	event := sdk.NewEvent(types.EventTypeEnrollIdSigner)

	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeIdSigner(ctx, addr)
		event = event.AppendAttributes(sdk.NewAttribute(types.EventAttrAddress, addr.String()))
	}
	events := []sdk.Event{
		event,
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeEnrollIdSigner),
		),
	}
	ctx.EventManager().EmitEvents(events)
	return &types.MsgEnrollIdSignerResponse{}, nil
}
