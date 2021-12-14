package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EnrollLoaders(goCtx context.Context, msg *types.MsgEnrollLoaders) (*types.MsgEnrollLoadersResponse, error) {
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
		k.activeShrpLoader(ctx, addr)
		if err := k.loadCoins(ctx, addr, types.AllowanceLoader); err != nil {
			return nil, err
		}
	}

	return &types.MsgEnrollLoadersResponse{
		Log: fmt.Sprintf("Successfully enroll SHRP loader %s", log),
	}, nil
}
