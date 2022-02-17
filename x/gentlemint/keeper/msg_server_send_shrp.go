package keeper

//import (
//	"context"
//	"fmt"
//	denom "github.com/sharering/shareledger/x/utils/demo"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	"github.com/sharering/shareledger/x/gentlemint/types"
//)
//
//func (k msgServer) SendShrp(goCtx context.Context, msg *types.MsgSendShrp) (*types.MsgSendShrpResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	if err := msg.ValidateBasic(); err != nil {
//		return nil, err
//	}
//	sentCoins, err := types.ParseShrpCoinsStr(msg.Amount)
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, msg.Amount)
//	}
//
//	senderAdd, err := sdk.AccAddressFromBech32(msg.Creator)
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, msg.Creator)
//	}
//	oldCoins := k.bankKeeper.GetAllBalances(ctx, senderAdd)
//
//	// Convert 1 current shrp to cent if sending ammount of cents is large than current cents
//	if oldCoins.AmountOf(denom.BaseUSD).LT(sentCoins.AmountOf(denom.BaseUSD)) {
//		if oldCoins.AmountOf(denom.ShrP).LTE(sentCoins.AmountOf(denom.ShrP)) {
//			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "address %v has %v", senderAdd.String(), oldCoins)
//		}
//		// Exchange 1 shrp to 100 cents for senderAdd
//		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAdd, types.ModuleName, denom.OneUSD); err != nil {
//			return nil, sdkerrors.Wrapf(err, "send %v coins to module", denom.OneUSD)
//		}
//		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, denom.OneHundredCents); err != nil {
//			return nil, sdkerrors.Wrapf(err, "mint %v coins", denom.OneHundredCents)
//		}
//		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, denom.OneUSD); err != nil {
//			return nil, sdkerrors.Wrapf(err, "burn %v coins", denom.OneUSD)
//		}
//		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, senderAdd, denom.OneHundredCents); err != nil {
//			return nil, sdkerrors.Wrapf(err, "send %v coins from module to account", denom.OneHundredCents)
//		}
//		// end exchange
//	}
//
//	toAdd, err := sdk.AccAddressFromBech32(msg.Address)
//	if err != nil {
//		return nil, sdkerrors.Wrap(err, msg.Address)
//	}
//	if err := k.bankKeeper.SendCoins(ctx, senderAdd, toAdd, sentCoins); err != nil {
//		return nil, sdkerrors.Wrapf(err, "%v send %v coins to %v", senderAdd.String(), sentCoins, toAdd.String())
//	}
//
//	log := fmt.Sprintf("Successfully Send ShrP {amount %s, from: %s, to: %s}", msg.Amount, msg.Creator, msg.Address)
//
//	return &types.MsgSendShrpResponse{
//		Log: log,
//	}, nil
//}
