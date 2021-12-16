package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevokeLoaders(goCtx context.Context, msg *types.MsgRevokeLoaders) (*types.MsgRevokeLoadersResponse, error) {
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
		if err := k.revokeShrpLoader(ctx, addr); err != nil {
			return nil, err
		}
	}

	return &types.MsgRevokeLoadersResponse{}, nil
}
