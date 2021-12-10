package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SendShrp(goCtx context.Context, msg *types.MsgSendShrp) (*types.MsgSendShrpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	sentCoins, err := types.ParseShrpCoinsStr(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(err, msg.Amount)
	}

	senderAdd, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, msg.Creator)
	}
	oldCoins := k.bankKeeper.GetAllBalances(ctx, senderAdd)

	// Convert 1 current shrp to cent if sending ammount of cents is large than current cents
	if oldCoins.AmountOf(types.DenomCent).LT(sentCoins.AmountOf(types.DenomCent)) {
		if oldCoins.AmountOf(types.DenomSHRP).LTE(sentCoins.AmountOf(types.DenomSHRP)) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "address %v has %v", senderAdd.String(), oldCoins)
		}
		// Exchange 1 shrp to 100 cents for senderAdd
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAdd, types.ModuleName, types.OneShrP); err != nil {
			return nil, sdkerrors.Wrapf(err, "send %v coins to module", types.OneShrP)
		}
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, types.OneHundredCents); err != nil {
			return nil, sdkerrors.Wrapf(err, "mint %v coins", types.OneHundredCents)
		}
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, types.OneShrP); err != nil {
			return nil, sdkerrors.Wrapf(err, "burn %v coins", types.OneShrP)
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, senderAdd, types.OneHundredCents); err != nil {
			return nil, sdkerrors.Wrapf(err, "send %v coins from module to account", types.OneHundredCents)
		}
		// end exchange
	}

	toAdd, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(err, msg.Address)
	}
	if err := k.bankKeeper.SendCoins(ctx, senderAdd, toAdd, sentCoins); err != nil {
		return nil, sdkerrors.Wrapf(err, "%v send %v coins to %v", senderAdd.String(), sentCoins, toAdd.String())
	}

	log := fmt.Sprintf("Successfully Send SHRP {amount %s, from: %s, to: %s}", msg.Amount, msg.Creator, msg.Address)

	return &types.MsgSendShrpResponse{
		Log: log,
	}, nil
}
