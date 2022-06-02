package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"strings"
)

func (k msgServer) ApproveOut(goCtx context.Context, msg *types.MsgApproveOut) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	batchId := k.AppendBatch(ctx, types.Batch{
		Signature: msg.Signature,
		TxIds:     msg.Ids,
		Status:    types.BatchStatusPending,
	})
	reqs, err := k.ChangeStatusRequests(ctx, msg.Ids, types.SwapStatusApproved, &batchId, true)
	if err != nil {
		return nil, err
	}

	total := sdk.NewDecCoins()
	reqIds := make([]string, 0, len(reqs))
	for _, r := range reqs {
		total = total.Add(sdk.NewDecCoinFromCoin(r.Amount))
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
	}

	events := sdk.NewEvent(types.EventTypeSwapApprove,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(types.EventTypeBatchId, fmt.Sprintf("%v", batchId)),
		sdk.NewAttribute(types.EventTypeApproverAction, types.SwapStatusApproved),
		sdk.NewAttribute(types.EventTypeApproverAddr, msg.Creator),
		sdk.NewAttribute(types.EventTypeBatchTotal, total.String()),
		sdk.NewAttribute(types.EventTypeSwapId, strings.Join(reqIds, ",")),
	)

	ctx.EventManager().EmitEvent(
		events,
	)

	return &types.MsgApproveResponse{
		BatchId: batchId,
	}, nil
}
