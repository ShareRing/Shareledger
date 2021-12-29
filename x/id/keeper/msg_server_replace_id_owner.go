package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/id/types"
)

func (k msgServer) ReplaceIdOwner(goCtx context.Context, msg *types.MsgReplaceIdOwner) (*types.MsgReplaceIdOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	id, found := k.GetFullIDByIDString(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrIdNotExisted, msg.Id)
	}

	// check if the new owner has id or not
	a, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OwnerAddress)
	}
	idOfNewOwner := k.GetIdByAddress(ctx, a)
	if len(idOfNewOwner) > 0 {
		return nil, sdkerrors.Wrap(types.ErrOwnerHasID, msg.OwnerAddress)
	}

	// check right backup
	if id.Data.BackupAddress != msg.BackupAddress {
		return nil, sdkerrors.Wrap(types.ErrWrongBackupAddr, msg.BackupAddress)
	}

	// update owner
	id.Data.OwnerAddress = msg.OwnerAddress
	k.SetID(ctx, id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventReplaceIDOwner,
			sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddress),
			sdk.NewAttribute(types.EventAttrId, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.BackupAddress),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventReplaceIDOwner),
		),
	})

	return &types.MsgReplaceIdOwnerResponse{}, nil
}
