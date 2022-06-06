package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"strings"
)

func (k msgServer) ApproveIn(goCtx context.Context, msg *types.MsgApproveIn) (*types.MsgApproveInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reqs, err := k.ChangeStatusRequests(ctx, msg.GetIds(), types.SwapStatusApproved, nil, false)
	if err != nil {
		return nil, err
	}
	total := sdk.NewCoins()
	reqIds := make([]string, 0, len(reqs))
	for _, r := range reqs {
		total = total.Add(r.Amount)
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
	}

	events := sdk.NewEvent(types.EventTypeSwapApprove,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		sdk.NewAttribute(types.EventAttrApproverAction, types.SwapStatusApproved),
		sdk.NewAttribute(types.EventAttrSwapType, types.SwapTypeIn),
		sdk.NewAttribute(types.EventAttrApproverAddr, msg.Creator),
		sdk.NewAttribute(types.EventAttrBatchTotal, total.String()),
		sdk.NewAttribute(types.EventAttrSwapId, strings.Join(reqIds, ",")),
	)

	ctx.EventManager().EmitEvent(
		events,
	)

	return &types.MsgApproveInResponse{}, nil
}
