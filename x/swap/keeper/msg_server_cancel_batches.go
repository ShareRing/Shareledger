package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CancelBatches(goCtx context.Context, msg *types.MsgCancelBatches) (*types.MsgCancelBatchesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batches := k.GetBatchesByIDs(ctx, msg.GetIds())
	var requestIDs []uint64
	batchIDs := make([]string, 0, len(batches))
	for _, batch := range batches {
		k.RemoveBatch(ctx, batch.GetId())
		requestIDs = append(requestIDs, batch.GetRequestIds()...)
		batchIDs = append(batchIDs, fmt.Sprintf("%x", batch.Id))
	}

	var zeroBatchNum uint64 = 0
	_, err := k.ChangeStatusRequests(ctx, requestIDs, types.BatchStatusPending, &zeroBatchNum, true)
	if err != nil {
		return &types.MsgCancelBatchesResponse{},
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can't change request to pending %s", err)
	}
	reqIDs := make([]string, 0, len(requestIDs))
	for _, i := range requestIDs {
		reqIDs = append(reqIDs, fmt.Sprintf("%v", i))
	}

	ctx.EventManager().EmitEvent(
		types.NewCancelBatchEvent(msg.Creator, batchIDs, reqIDs),
	)

	return &types.MsgCancelBatchesResponse{}, nil
}
