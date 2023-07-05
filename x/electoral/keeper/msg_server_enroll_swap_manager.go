package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/electoral/types"
)

func (k msgServer) EnrollSwapManagers(goCtx context.Context, msg *types.MsgEnrollSwapManagers) (*types.MsgEnrollSwapManagersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, a := range msg.Addresses {
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeSwapManager(ctx, addr)

	}

	return &types.MsgEnrollSwapManagersResponse{}, nil
}
