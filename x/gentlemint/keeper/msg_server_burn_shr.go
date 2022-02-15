package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

func (k msgServer) BurnShr(goCtx context.Context, msg *types.MsgBurnShr) (*types.MsgBurnShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	v, err := sdk.NewDecFromStr(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	baseCoins, err := denom.NormalizeCoins(sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom.Shr, v)), nil)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if err := k.burnCoins(ctx, msg.GetSigners()[0], sdk.NewCoins(baseCoins)); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", baseCoins, msg.Creator)
	}

	return &types.MsgBurnShrResponse{
		Log: fmt.Sprintf("Successfully burn %v", baseCoins),
	}, nil
}
