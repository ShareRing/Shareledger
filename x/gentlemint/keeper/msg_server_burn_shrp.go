package keeper

import (
	"context"
	"fmt"

	"github.com/sharering/shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BurnShrp(goCtx context.Context, msg *types.MsgBurnShrp) (*types.MsgBurnShrpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	amt, err := types.ParseShrpCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}

	if err := k.burnCoins(ctx, msg.GetSigners()[0], amt); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", amt, msg.Creator)
	}
	log := fmt.Sprintf("Successfully burn coins %s", msg.Amount)

	return &types.MsgBurnShrpResponse{
		Log: log,
	}, nil
}
