package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EnrollLoaders(goCtx context.Context, msg *types.MsgEnrollLoaders) (*types.MsgEnrollLoadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	log := "SHRP loaders' addresses: "
	for _, a := range msg.Addresses {
		log = fmt.Sprintf("%s %s", log, a)
		addr, err := sdk.AccAddressFromBech32(a)
		if err != nil {
			return nil, err
		}
		k.activeShrpLoader(ctx, addr)
		if err := k.gk.LoadAllowanceLoader(ctx, addr); err != nil {
			return nil, err
		}
	}

	return &types.MsgEnrollLoadersResponse{
		Log: fmt.Sprintf("Successfully enroll SHRP loader %s", log),
	}, nil
}
