package keeper

import (
	"context"
	"fmt"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// LoadShrp function is used to load the given amount of ShrP to the given recipient
// - Automatically buy 10SHR for the recipient if there is less than 10 shr
// - Send 1SHR from recipient to loader as the loading fee
func (k msgServer) LoadShrp(goCtx context.Context, msg *types.MsgLoadShrp) (*types.MsgLoadShrpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(err, msg.Address)
	}

	oldCoins := k.bankKeeper.GetAllBalances(ctx, receiverAddr)
	amt, err := types.ParseShrpCoinsStr(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, msg.Amount)
	}

	adjustCoins, err := types.AddShrpCoins(oldCoins, amt)
	if err != nil {
		return nil, err
	}

	if adjustCoins.Add.Len() > 0 {
		if err := k.loadCoins(ctx, receiverAddr, adjustCoins.Add); err != nil {
			return nil, sdkerrors.Wrapf(err, "load coins, %v, to address %v", amt, msg.Address)
		}
	}
	if adjustCoins.Sub.Len() > 0 {
		if err := k.burnCoins(ctx, receiverAddr, adjustCoins.Sub); err != nil {
			return nil, sdkerrors.Wrapf(err, "burn coins, %v, from address %v", adjustCoins.Sub, msg.Address)
		}
	}

	oldCoins = k.bankKeeper.GetAllBalances(ctx, receiverAddr)
	oldShr := oldCoins.AmountOf(denom.PShr)

	// if there is less than required shr amount, buy more.
	if oldShr.LT(types.RequiredPSHRAmt) {
		if err := k.buyPShr(ctx, types.RequiredPSHRAmt, receiverAddr); err != nil {
			return nil, sdkerrors.Wrapf(err, "buy minimum required pShr, %v, for address %v", types.RequiredPSHRAmt, receiverAddr.String())
		}
	}

	//Pay fee for loader who is a creator of this message
	if err := k.bankKeeper.SendCoins(ctx, receiverAddr, msg.GetSigners()[0], types.FeeLoadSHRP); err != nil {
		return nil, sdkerrors.Wrapf(err, "pay fee, %v, to approver, %v", types.FeeLoadSHRP, msg.Creator)
	}
	log := fmt.Sprintf("Successfully load ShrP {amount %s, address: %s}", msg.Amount, msg.Address)

	return &types.MsgLoadShrpResponse{
		Log: log,
	}, nil
}
