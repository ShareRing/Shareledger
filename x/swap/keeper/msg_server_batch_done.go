package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) CompleteBatch(goCtx context.Context, msg *types.MsgCompleteBatch) (*types.MsgCompleteBatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	batch, found := k.GetBatch(ctx, msg.GetBatchId())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "batch id=%s not found", msg.GetBatchId())
	}
	k.RemoveBatch(ctx, batch.Id)
	requests, err := k.getRequestsFromIds(ctx, batch.TxIds, types.SwapStatusApproved)
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

	events := sdk.NewEvent(types.EventTypeBatchDone,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(types.EventAttrBatchId, fmt.Sprintf("%v", batch.Id)),
		sdk.NewAttribute(types.EventAttrBatchNetwork, batch.Network),
		sdk.NewAttribute(types.EventAttrBatchTxIDs, strings.Join(reqIds, ",")),
	)
	ctx.EventManager().EmitEvent(
		events,
	)
	return &types.MsgCompleteBatchResponse{}, nil
}
