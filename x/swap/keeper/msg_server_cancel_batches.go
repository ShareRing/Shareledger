package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CancelBatches(goCtx context.Context, msg *types.MsgCancelBatches) (*types.MsgCancelBatchesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batches := k.GetBatchesByIDs(ctx, msg.GetIds())
	var requestIDs []uint64
	for _, batch := range batches {
		k.RemoveBatch(ctx, batch.GetId())
		requestIDs = append(requestIDs, batch.GetTxIds()...)
	}

	var zeroBatchNum uint64 = 0
	_, err := k.ChangeStatusRequests(ctx, requestIDs, types.BatchStatusPending, &zeroBatchNum, true)
	if err != nil {
		return &types.MsgCancelBatchesResponse{},
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can't change request to pending %s", err)
	}

	return &types.MsgCancelBatchesResponse{}, nil
}
