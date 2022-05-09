package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) RejectSwap(ctx sdk.Context, msg *types.MsgReject) ([]types.Request, error) {
	reqs, err := k.ChangeStatusRequests(ctx, msg.Ids, types.SwapStatusRejected, nil, true)
	if err != nil {
		return nil, err
	}
	return reqs, nil

}
