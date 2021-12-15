package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDocument(goCtx context.Context, msg *types.MsgCreateDocument) (*types.MsgCreateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	doc := types.Document{
		Issuer:  msg.Issuer,
		Holder:  msg.Holder,
		Proof:   msg.Proof,
		Data:    msg.Data,
		Version: 0,
	}

	// check existing doc
	_, found := k.GetDoc(ctx, doc)
	if found {
		return nil, sdkerrors.Wrap(types.ErrDocExisted, doc.String())
	}

	k.SetDoc(ctx, &doc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateDoc,
			sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer),
			sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder)),
			sdk.NewAttribute(types.EventAttrProof, string(msg.Proof)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateDoc),
		),
	})

	return &types.MsgCreateDocumentResponse{}, nil
}
