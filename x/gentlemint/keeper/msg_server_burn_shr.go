package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) BurnPShr(goCtx context.Context, msg *types.MsgBurnPShr) (*types.MsgBurnPShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	shrCoins, err := types.ParsePShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}
	if err := k.burnCoins(ctx, msg.GetSigners()[0], shrCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", shrCoins, msg.Creator)
	}

	return &types.MsgBurnPShrResponse{
		Log: fmt.Sprintf("Successfully burn %v pshr", msg.Amount),
	}, nil
}
