package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnrollAccountOperator(goCtx context.Context, msg *types.MsgEnrollAccountOperator) (*types.MsgEnrollAccountOperatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if !k.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAuthority)
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
	return &types.MsgEnrollAccountOperatorResponse{}, nil
}
