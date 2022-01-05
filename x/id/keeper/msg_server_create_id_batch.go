package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/id/types"
)

func (k msgServer) CreateIds(goCtx context.Context, msg *types.MsgCreateIds) (*types.MsgCreateIdsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	for i := 0; i < len(msg.Id); i++ {
		data := types.BaseID{
			IssuerAddress: msg.IssuerAddress,
			BackupAddress: msg.BackupAddress[i],
			OwnerAddress:  msg.OwnerAddress[i],
			ExtraData:     msg.ExtraData[i],
		}

		id := types.Id{
			Id:   msg.Id[i],
			Data: &data,
		}

		// check id existing
		if k.IsExist(ctx, &id) {
			return nil, sdkerrors.Wrap(types.ErrIdExisted, id.String())
		}

		k.SetID(ctx, &id)
		event := sdk.NewEvent(
			types.EventCreateID,
			sdk.NewAttribute(types.EventAttrIssuer, msg.IssuerAddress),
			sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddress[i]),
			sdk.NewAttribute(types.EventAttrId, msg.Id[i]),
		)
		ctx.EventManager().EmitEvent(event)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.IssuerAddress),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventCreateIDs),
		),
	})

	return &types.MsgCreateIdsResponse{}, nil
}
