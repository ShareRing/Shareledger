package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/document/types"
)

func (k msgServer) RevokeDocument(goCtx context.Context, msg *types.MsgRevokeDocument) (*types.MsgRevokeDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	queryDoc := types.Document{
		Issuer: msg.Issuer,
		Holder: msg.Holder,
		Proof:  msg.Proof,
	}

	// check existing doc
	existingDoc, found := k.GetDocByProof(ctx, queryDoc)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrDocNotExisted, existingDoc.String())
	}

	existingDoc.Version = int32(types.DocRevokeFlag)
	k.SetDoc(ctx, &existingDoc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRevokeDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer),
			sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder)),
			sdk.NewAttribute(types.EventAttrProof, string(msg.Proof)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeRevokeDoc),
		),
	})

	return &types.MsgRevokeDocumentResponse{}, nil
}
