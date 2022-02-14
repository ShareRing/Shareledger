package keeper

import (
	"context"
	"fmt"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) LoadShr(goCtx context.Context, msg *types.MsgLoadShr) (*types.MsgLoadShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	v, e := sdk.NewDecFromStr(msg.Amount)
	if e != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%v", e)
	}
	if v.LTE(sdk.NewDec(0)) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "amount should be larger than 0")
	}

	shrCoin := sdk.NewDecCoinFromDec(denom.Shr, v)
	buyCoin, err := denom.NormalizeCoins(sdk.NewDecCoins(shrCoin), nil)
	if err != nil {
		return nil, err
	}

	if !k.PShrMintPossible(ctx, buyCoin.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "PShr possible mint exceeded")
	}

	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, msg.Address)
	}

	if err := k.loadCoins(ctx, addr, sdk.NewCoins(buyCoin)); err != nil {
		return nil, err
	}
	log := fmt.Sprintf("Successfully loaded pshr {address: %s, amount %v}", msg.Address, buyCoin)

	return &types.MsgLoadShrResponse{
		Log: log,
	}, nil
}
