package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (k Keeper) ApproveIn(ctx sdk.Context, msg *types.MsgApproveIn) ([]types.Request, error) {

	reqs, err := k.ChangeStatusRequestsImprovement(ctx, msg.GetTxnIDs(), types.SwapStatusApproved, nil)
	if err != nil {
		return nil, err
	}

	return reqs, nil
}
