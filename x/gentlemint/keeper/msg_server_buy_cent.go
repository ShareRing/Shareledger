package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) BuyCent(goCtx context.Context, msg *types.MsgBuyCent) (*types.MsgBuyCentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	buyerAddr := msg.GetSigners()[0]

	amt, _ := sdk.NewIntFromString(msg.Amount)
	shrpAmount := sdk.NewCoins(sdk.NewCoin(types.DenomSHRP, amt))
	centAmount := sdk.NewCoins(sdk.NewCoin(types.DenomCent, amt.Mul(sdk.NewInt(100))))

	if err := k.burnCoins(ctx, buyerAddr, shrpAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins for address %v", shrpAmount, buyerAddr.String())
	}
	if err := k.loadCoins(ctx, buyerAddr, centAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "load %v coins for address %v", centAmount, buyerAddr.String())
	}

	log := fmt.Sprintf("Successfull exchange %v shrp to cent for address %s", msg.Amount, buyerAddr.String())

	return &types.MsgBuyCentResponse{
		Log: log,
	}, nil
}
