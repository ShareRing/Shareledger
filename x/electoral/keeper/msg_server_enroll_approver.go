package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollApprover(goCtx context.Context, msg *types.MsgEnrollApprover) (*types.MsgEnrollApproverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollApproverResponse{}, nil
}
