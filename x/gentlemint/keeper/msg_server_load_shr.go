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
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coins, err := types.ParseShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}

	if k.ShrMintPossible(ctx, coins.AmountOf(types.DenomSHR)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "SHR possible mint exceeded")
	}

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
