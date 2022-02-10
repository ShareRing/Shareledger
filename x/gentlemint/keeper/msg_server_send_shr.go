package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) SendPShr(goCtx context.Context, msg *types.MsgSendPShr) (*types.MsgSendPShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid message")
	}
	amt, _ := sdk.NewIntFromString(msg.Amount)

	oldCoin := k.bankKeeper.GetBalance(ctx, msg.GetSigners()[0], types.DenomPSHR)

	if oldCoin.Amount.LT(amt) {
		shrToBuy := sdk.NewCoin(types.DenomPSHR, amt.Sub(oldCoin.Amount))
		if err := k.buyPShr(ctx, shrToBuy.Amount, msg.GetSigners()[0]); err != nil {
			return nil, sdkerrors.Wrapf(err, "buy %v pshr for address %v", shrToBuy, msg.Creator)
		}
	}
	sendCoins := sdk.NewCoins(sdk.NewCoin(types.DenomPSHR, amt))
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoins(ctx, msg.GetSigners()[0], receiverAddr, sendCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "send coins, %+v, from %s to %s", sendCoins, msg.Creator, msg.Address)
	}
	log := fmt.Sprintf("Successfully Send SHR {amount %s, from: %s, to: %s}", msg.Amount, msg.Creator, msg.Address)

	return &types.MsgSendPShrResponse{
		Log: log,
	}, nil
}
