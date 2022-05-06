package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeApprover(goCtx context.Context, msg *types.MsgRevokeApprover) (*types.MsgRevokeApproverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevokeApproverResponse{}, nil
}
