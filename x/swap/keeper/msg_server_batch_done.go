package keeper

import (
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) SetBatchDone(goCtx context.Context, msg *types.MsgSetBatchDone) (*types.MsgSetBatchDoneResponse, error) {
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

	k.MoveRequest(ctx, types.SwapStatusApproved, types.SwapStatusDone, requests, nil, true)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeBatchDone).
			AppendAttributes(
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(types.EventTypeAttrBatchID, fmt.Sprintf("%v", batch.Id)),
				sdk.NewAttribute(types.EventTypeAttrBatchTxIDs, fmt.Sprintf("%v", batch.TxIds)),
			))
	return &types.MsgSetBatchDoneResponse{}, nil
}
