package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) Reject(goCtx context.Context, msg *types.MsgReject) (*types.MsgRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reqs, err := k.Keeper.RejectSwap(ctx, msg)
	if err != nil {
		return nil, err
	}
	total := sdk.NewDecCoins()
	reqIds := make([]string, 0, len(reqs))
	for _, r := range reqs {
		total = total.Add(sdk.NewDecCoinFromCoin(*r.GetAmount()))
		reqIds = append(reqIds, fmt.Sprintf("%v", r.Id))
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeSwapApprove,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeApproverAction, types.SwapStatusRejected),
			sdk.NewAttribute(types.EventTypeRejectArr, msg.Creator),
			sdk.NewAttribute(types.EventTypeBatchTotal, total.String()),
			sdk.NewAttribute(types.EventTypeSwapId, strings.Join(reqIds, ",")),
		),
	)

	return &types.MsgRejectResponse{}, nil
}
