package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CompleteBatch(goCtx context.Context, msg *types.MsgCompleteBatch) (*types.MsgCompleteBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetBatch(ctx, msg.GetBatchId())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "batch id=%s not found", msg.GetBatchId())
	}
	k.RemoveBatch(ctx, batch.Id)
	requests, err := k.getRequestsFromIds(ctx, batch.ReqIDs, types.SwapStatusApproved)
	if err != nil {
		return nil, err
	}

	if err := k.MoveRequest(ctx, types.SwapStatusApproved, types.SwapStatusDone, requests, nil, true); err != nil {
		return nil, err
	}
	reqIds := make([]string, 0, len(requests))
	for _, r := range requests {
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
	}

	ctx.EventManager().EmitEvent(
		types.NewCompleteBatchEvent(msg.Creator, batch.Id, reqIds),
	)
	return &types.MsgCompleteBatchResponse{}, nil
}
