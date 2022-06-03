package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) UpdateBatch(goCtx context.Context, msg *types.MsgUpdateBatch) (*types.MsgUpdateBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetBatch(ctx, msg.GetBatchId())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "batch id=%s not found", msg.GetBatchId())
	}
	if msg.GetStatus() != types.BatchStatusPending && msg.GetStatus() != types.BatchStatusDone {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "status %s is not valid in blockchain", msg.GetStatus())
	}
	batch.Status = msg.GetStatus()
	k.SetBatch(ctx, batch)

	//The batch is done we detele the swap request
	if batch.GetStatus() == types.SwapStatusDone {
		k.RemoveRequestFromStore(ctx, batch.GetTxIds())
	}

	return &types.MsgUpdateBatchResponse{}, nil
}
