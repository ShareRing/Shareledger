package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevokeLoaders(goCtx context.Context, msg *types.MsgRevokeLoaders) (*types.MsgRevokeLoadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !k.IsAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, types.ErrSenderIsNotAuthority)
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
	return &types.MsgRevokeLoadersResponse{
		Log: fmt.Sprintf("Successfully revoke SHRP loader %s", log),
	}, nil
}
