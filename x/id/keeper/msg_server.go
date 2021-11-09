package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateId(goCtx context.Context, msg *types.MsgCreateId) (*types.MsgCreateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := msg.ToID()

	// Check existing
	if k.IsExist(ctx, &id) {
		return nil, types.ErrIdExisted
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

func (k msgServer) CreateIdInBatch(goCtx context.Context, msg *types.MsgCreateIdBatch) (*types.MsgCreateIdInBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for i := 0; i < len(msg.Id); i++ {
		id := types.NewID(msg.Id[i], msg.IssuerAddress, msg.BackupAddress[i], msg.OwnerAddress[i], msg.ExtraData[i])
		// Check id existing
		if k.IsExist(ctx, &id) {
			return nil, types.ErrIdExisted
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
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventCreateIDBatch),
		),
	})

	return &types.MsgCreateIdInBatchResponse{}, nil
}

func (k msgServer) UpdateId(goCtx context.Context, msg *types.MsgUpdateId) (*types.MsgUpdateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.GetIDByIdString(ctx, msg.Id)

	if id.IsEmpty() {
		return nil, types.ErrIdNotExisted
	}

	// Update extra data
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

func (k msgServer) ReplaceIdOwner(goCtx context.Context, msg *types.MsgReplaceIdOwner) (*types.MsgReplaceIdOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.GetIDByIdString(ctx, msg.Id)

	// Check if the id is existed or not
	if id.IsEmpty() {
		return nil, types.ErrIdNotExisted
	}

	// Check if the new owner has id or not
	idOfNewOwner := k.GetIdOnlyByAddress(ctx, sdk.AccAddress(msg.OwnerAddress))
	if len(idOfNewOwner) > 0 {
		return nil, types.ErrIdExisted
	}

	// Check right backup
	if id.Data.BackupAddress != msg.BackupAddress {
		return nil, types.ErrWrongBackupAddr
	}

	// Update owner
	id.Data.OwnerAddress = msg.OwnerAddress
	k.SetID(ctx, id)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventReplaceIDOwner,
		sdk.NewAttribute(types.EventAttrOwner, msg.OwnerAddress),
		sdk.NewAttribute(types.EventAttrId, msg.Id),
	))

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
