package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyerAddr, types.ModuleName, shrpAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "send %v to module %v", shrpAmount, types.ModuleName)
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, centAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "mint %v in module %v", centAmount, types.ModuleName)
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, shrpAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "burns %v coins in module %v", shrpAmount, types.ModuleName)
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyerAddr, centAmount); err != nil {
		return nil, sdkerrors.Wrapf(err, "send %v coins to account %v", centAmount, buyerAddr.String())
	}

	log := fmt.Sprintf("Successfull exchange %v shrp to cent for address %s", msg.Amount, buyerAddr.String())

	return &types.MsgBuyCentResponse{
		Log: log,
	}, nil
}
