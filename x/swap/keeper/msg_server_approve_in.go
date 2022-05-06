package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) ApproveIn(goCtx context.Context, msg *types.MsgApproveIn) (*types.MsgApproveInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reqs, err := k.ChangeStatusRequests(ctx, msg.GetTxnIDs(), types.SwapStatusApproved, nil, false)
	if err != nil {
		return nil, err
	}
	total := sdk.NewDecCoins()
	reqIds := make([]string, 0, len(reqs))
	for _, r := range reqs {
		total = total.Add(*r.Amount)
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeSwapApprove,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeApproverAction, types.SwapStatusApproved),
			sdk.NewAttribute(types.EventTypeApproverAddr, msg.Creator),
			sdk.NewAttribute(types.EventTypeBatchTotal, total.String()),
			sdk.NewAttribute(types.EventTypeSwapId, strings.Join(reqIds, ",")),
		))

	return &types.MsgApproveInResponse{}, nil
}
