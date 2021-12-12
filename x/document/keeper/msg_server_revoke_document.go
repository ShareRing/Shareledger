package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeDocument(goCtx context.Context, msg *types.MsgRevokeDocument) (*types.MsgRevokeDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeDocumentResponse{}, nil
}
