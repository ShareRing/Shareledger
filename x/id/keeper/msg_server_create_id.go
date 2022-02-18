package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/id/types"
)

func (k msgServer) CreateId(goCtx context.Context, msg *types.MsgCreateId) (*types.MsgCreateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	id := msg.ToID()

	// check existing
	if k.IsExist(ctx, &id) {
		return nil, sdkerrors.Wrap(types.ErrIdExisted, id.String())
	}

	k.SetID(ctx, &id)

	// Emit 2 events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventCreateID,
			sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddress),
			sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddress),
			sdk.NewAttribute(types.EventAttrId, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddress),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventCreateID),
		),
	})

	return &types.MsgCreateIdResponse{}, nil
}
