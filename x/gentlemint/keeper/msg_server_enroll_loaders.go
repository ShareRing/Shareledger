package keeper

import (
	"context"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollLoaders(goCtx context.Context, msg *types.MsgEnrollLoaders) (*types.MsgEnrollLoadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	for _, a := range msg.Addresses {
		add, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.setSHRPLoaderStatus(ctx, add, types.StatusSHRPLoaderActived)
	}

	return &types.MsgEnrollLoadersResponse{}, nil
}
