package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollVoter(goCtx context.Context, msg *types.MsgEnrollVoter) (*types.MsgEnrollVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEnrollVoterResponse{}, nil
}
