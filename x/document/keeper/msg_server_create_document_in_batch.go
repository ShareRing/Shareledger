package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDocumentInBatch(goCtx context.Context, msg *types.MsgCreateDocumentInBatch) (*types.MsgCreateDocumentInBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateDocumentInBatchResponse{}, nil
}
