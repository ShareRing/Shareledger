package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) LoadPShr(goCtx context.Context, msg *types.MsgLoadPShr) (*types.MsgLoadPShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coins, err := types.ParsePShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}

	if !k.PShrMintPossible(ctx, coins.AmountOf(types.DenomPSHR)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SHR possible mint exceeded")
	}

	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, msg.Address)
	}

	if err := k.loadCoins(ctx, addr, coins); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully loaded pshr {address: %s, amount %v}", msg.Address, coins)

	return &types.MsgLoadPShrResponse{
		Log: log,
	}, nil
}
