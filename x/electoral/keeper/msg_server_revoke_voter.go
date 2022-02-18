package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeVoter(goCtx context.Context, msg *types.MsgRevokeVoter) (*types.MsgRevokeVoterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	addr, _ := sdk.AccAddressFromBech32(msg.Address)
	if err := k.revokeVoter(ctx, addr); err != nil {
		return nil, err
	}

	return &types.MsgRevokeVoterResponse{}, nil
}
