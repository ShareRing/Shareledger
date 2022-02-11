package keeper

import (
	"context"
	"fmt"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) SendShr(goCtx context.Context, msg *types.MsgSendShr) (*types.MsgSendShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid message")
	}
	decShr, err := sdk.NewDecFromStr(msg.Amount)
	amt := denom.NormalizeCoins(sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom.Shr, decShr)), sdk.NewDec(1)).Amount

	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	oldCoin := k.bankKeeper.GetBalance(ctx, msg.GetSigners()[0], denom.PShr)

	if oldCoin.Amount.LT(amt) {
		shrToBuy := sdk.NewCoin(denom.PShr, amt.Sub(oldCoin.Amount))
		if err := k.buyPShr(ctx, shrToBuy.Amount, msg.GetSigners()[0]); err != nil {
			return nil, sdkerrors.Wrapf(err, "buy %v pshr for address %v", shrToBuy, msg.Creator)
		}
	}
	sendCoins := sdk.NewCoins(sdk.NewCoin(denom.PShr, amt))
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoins(ctx, msg.GetSigners()[0], receiverAddr, sendCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "send coins, %+v, from %s to %s", sendCoins, msg.Creator, msg.Address)
	}
	log := fmt.Sprintf("Successfully Send SHR {amount %s, from: %s, to: %s}", msg.Amount, msg.Creator, msg.Address)

	return &types.MsgSendShrResponse{
		Log: log,
	}, nil
}
