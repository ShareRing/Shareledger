package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateIdBatch(goCtx context.Context, msg *types.MsgCreateIdBatch) (*types.MsgCreateIdBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateIdBatchResponse{}, nil
}
