package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollVoter(goCtx context.Context, msg *types.MsgEnrollVoter) (*types.MsgEnrollVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	addr, _ := sdk.AccAddressFromBech32(msg.Address)
	k.activeVoter(ctx, addr)
	return &types.MsgEnrollVoterResponse{}, nil
}
