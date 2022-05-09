package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollRelayers(goCtx context.Context, msg *types.MsgEnrollRelayers) (*types.MsgEnrollRelayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeRelayer(ctx, addr)

	}

	return &types.MsgEnrollRelayersResponse{}, nil
}
