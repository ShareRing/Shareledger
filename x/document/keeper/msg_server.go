package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
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

func (k msgServer) CreateDocument(goCtx context.Context, msg *types.MsgCreateDoc) (*types.MsgCreateDocResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	doc := types.Document{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof, Data: msg.Data, Version: 0}

	// Check doc is existed
	existingDoc := k.GetDoc(ctx, doc)
	if len(existingDoc.Proof) > 0 {
		return nil, types.ErrDocExisted
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

	return &types.MsgCreateDocResponse{}, nil
}

func (k msgServer) CreateDocumentInBatch(goCtx context.Context, msg *types.MsgCreateDocBatch) (*types.MsgCreateDocBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for i := 0; i < len(msg.Holder); i++ {
		doc := types.Document{Issuer: msg.Issuer, Holder: msg.Holder[i], Proof: msg.Proof[i], Data: msg.Data[i], Version: 0}

		// Check doc is existed
		existingDoc := k.GetDoc(ctx, doc)
		if len(existingDoc.Proof) > 0 {
			return nil, types.ErrDocExisted
		}

		k.SetDoc(ctx, &doc)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreateDoc,
				sdk.NewAttribute(types.EventAttrIssuer, msg.Issuer),
				sdk.NewAttribute(types.EventAttrHolder, string(msg.Holder[i])),
				sdk.NewAttribute(types.EventAttrProof, string(msg.Proof[i])),
				sdk.NewAttribute(types.EventAttrData, string(msg.Data[i])),
			),
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateDoc),
		),
	)

	return &types.MsgCreateDocBatchResponse{}, nil
}

func (k msgServer) UpdateDocument(goCtx context.Context, msg *types.MsgUpdateDoc) (*types.MsgUpdateDocResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	queryDoc := types.Document{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof, Data: msg.Data, Version: 0}

	// Check doc is existed
	existingDoc := k.GetDocByProof(ctx, queryDoc)
	if len(existingDoc.Proof) == 0 {
		return nil, types.ErrDocNotExisted
	}

	if existingDoc.Holder != msg.Holder || existingDoc.Issuer != queryDoc.Issuer {
		return nil, types.ErrDocNotExisted
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

	return &types.MsgUpdateDocResponse{}, nil
}

func (k msgServer) RevokeDocument(goCtx context.Context, msg *types.MsgRevokeDoc) (*types.MsgRevokeDocResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	queryDoc := types.Document{Issuer: msg.Issuer, Holder: msg.Holder, Proof: msg.Proof}

	// Check doc is existed
	existingDoc := k.GetDocByProof(ctx, queryDoc)
	if len(existingDoc.Proof) == 0 {
		return nil, types.ErrDocNotExisted
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

	return &types.MsgRevokeDocResponse{}, nil
}
