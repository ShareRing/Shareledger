package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
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

	return &types.MsgApproveInResponse{}, nil
}
