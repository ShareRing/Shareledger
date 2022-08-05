package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k msgServer) ApproveIn(goCtx context.Context, msg *types.MsgApproveIn) (*types.MsgApproveInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	requests, err := k.getRequestsFromIds(ctx, msg.GetIds(), types.SwapStatusPending)
	for _, r := range requests {
		for _, hash := range r.TxEvents {
			_, found := k.GetPastTxEvent(ctx, hash.TxHash, hash.LogIndex)
			if found {
				return nil, sdkerrors.Wrap(types.ErrDuplicatedSwapIn, "tx hash was processed in blockchain")
			}
		}
	}

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

	ctx.EventManager().EmitEvent(
		types.NewApproveInEvent(msg.Creator, reqIds, total),
	)

	return &types.MsgApproveInResponse{}, nil
}
