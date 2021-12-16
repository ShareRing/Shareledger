package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateId(goCtx context.Context, msg *types.MsgUpdateId) (*types.MsgUpdateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	id, found := k.GetFullIDByIDString(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrIdNotExisted, msg.Id)
	}

	// update extra data
	id.Data.ExtraData = msg.ExtraData
	k.SetID(ctx, id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventUpdateID,
			sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddress),
			sdk.NewAttribute(types.EventAttrId, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddress),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventUpdateID),
		),
	})

	return &types.MsgUpdateIdResponse{}, nil
}
