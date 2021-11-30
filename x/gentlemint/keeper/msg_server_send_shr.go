package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SendShr(goCtx context.Context, msg *types.MsgSendShr) (*types.MsgSendShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid message")
	}
	amt, _ := sdk.NewIntFromString(msg.Amount)

	oldCoin := k.bankKeeper.GetBalance(ctx, msg.GetSigners()[0], types.DemonSHR)

	if oldCoin.Amount.LT(amt) {
		shrToBuy := sdk.NewCoin(types.DemonSHR, amt.Sub(oldCoin.Amount))
		if err := k.buyShr(ctx, shrToBuy.Amount, msg.GetSigners()[0]); err != nil {
			return nil, sdkerrors.Wrapf(err, "buy %v shr for address %v", shrToBuy, msg.Creator)
		}
	}
	sendCoins := sdk.NewCoins(sdk.NewCoin(types.DemonSHR, amt))
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
