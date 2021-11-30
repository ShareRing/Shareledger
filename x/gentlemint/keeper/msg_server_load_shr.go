package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LoadShr(goCtx context.Context, msg *types.MsgLoadShr) (*types.MsgLoadShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.isAuthority(ctx, msg.GetSigners()[0]) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Approver's Address is not authority")
	}

	amt, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount)
	}

	if k.ShrMintPossible(ctx, amt) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SHR possible mint exceeded")
	}
	coins := sdk.NewCoins(sdk.NewCoin(types.DenomSHR, amt))
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, msg.Address)
	}

	if err := k.loadCoins(ctx, addr, coins); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully loaded shr {address: %s, amount %v}", msg.Address, coins)

	return &types.MsgLoadShrResponse{
		Log: log,
	}, nil
}
