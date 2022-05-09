package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) RevokeRelayers(goCtx context.Context, msg *types.MsgRevokeRelayers) (*types.MsgRevokeRelayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		if err := k.removeRelayer(ctx, addr); err != nil {
			return nil, err
		}

	}

	return &types.MsgRevokeRelayersResponse{}, nil
}
