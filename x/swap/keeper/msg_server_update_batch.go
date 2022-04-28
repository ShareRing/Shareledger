package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) UpdateBatch(goCtx context.Context, msg *types.MsgUpdateBatch) (*types.MsgUpdateBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateBatchResponse{}, nil
}
