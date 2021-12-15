package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDocumentInBatch(goCtx context.Context, msg *types.MsgCreateDocumentInBatch) (*types.MsgCreateDocumentInBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	for i := 0; i < len(msg.Holder); i++ {
		doc := types.Document{
			Issuer:  msg.Issuer,
			Holder:  msg.Holder[i],
			Proof:   msg.Proof[i],
			Data:    msg.Data[i],
			Version: 0,
		}

		// check existing doc
		_, found := k.GetDoc(ctx, doc)
		if found {
			return nil, sdkerrors.Wrap(types.ErrDocExisted, doc.String())
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

	return &types.MsgCreateDocumentInBatchResponse{}, nil
}
