package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDocument(goCtx context.Context, msg *types.MsgUpdateDocument) (*types.MsgUpdateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateDocumentResponse{}, nil
}
