package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	denom "github.com/sharering/shareledger/x/utils/demo"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) Load(goCtx context.Context, msg *types.MsgLoad) (*types.MsgLoadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	destAdd, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	coins, err := sdk.ParseDecCoins(msg.Coins)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}
	baseCoins, err := denom.NormalizeToBaseCoins(coins, false)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, err.Error())
	}
	if !k.BaseMintPossible(ctx, baseCoins.AmountOf(denom.Base)) {
		return nil, sdkerrors.Wrapf(types.ErrBaseSupplyExceeded, "load %v", baseCoins)
	}
	if err := k.loadCoins(ctx, destAdd, baseCoins); err != nil {
		return nil, err
	}

	// Pay fee for loader who is a creator of this message
	if msg.Creator != msg.Address {
		if err := k.bankKeeper.SendCoins(ctx, destAdd, msg.GetSigners()[0], types.FeeLoadSHRP); err != nil {
			return nil, sdkerrors.Wrapf(err, "pay fee, %v, to approver, %v", types.FeeLoadSHRP, msg.Creator)
		}
	}

	return &types.MsgLoadResponse{}, nil
}