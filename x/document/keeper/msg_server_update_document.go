package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/document/types"
)

func (k msgServer) UpdateDocument(goCtx context.Context, msg *types.MsgUpdateDocument) (*types.MsgUpdateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	queryDoc := types.Document{
		Issuer:  msg.Issuer,
		Holder:  msg.Holder,
		Proof:   msg.Proof,
		Data:    msg.Data,
		Version: 0,
	}

	existingDoc, found := k.GetDocByProof(ctx, queryDoc)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrDocNotExisted, queryDoc.String())
	}

	if existingDoc.Holder != msg.Holder || existingDoc.Issuer != msg.Issuer {
		return nil, sdkerrors.Wrap(types.ErrDocNotExisted, queryDoc.String())
	}

	if existingDoc.Version == int32(types.DocRevokeFlag) {
		return nil, sdkerrors.Wrap(types.ErrDocRevoked, queryDoc.String())
	}

	existingDoc.Data = msg.Data
	existingDoc.Version = existingDoc.Version + 1

	k.SetDoc(ctx, &existingDoc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer),
			sdk.NewAttribute(types.EventAttrHolder, msg.Holder),
			sdk.NewAttribute(types.EventAttrProof, msg.Proof),
			sdk.NewAttribute(types.EventAttrData, msg.Data),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeUpdateDoc),
		),
	})

	return &types.MsgUpdateDocumentResponse{}, nil
}
