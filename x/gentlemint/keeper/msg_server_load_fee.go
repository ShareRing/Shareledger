package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) LoadFee(goCtx context.Context, msg *types.MsgLoadFee) (*types.MsgLoadFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	decCoins, err := sdk.ParseDecCoins(msg.Shrp)
	if err != nil {
		return nil, err
	}
	shrp := types.ShrpDecCoinsToCoins(decCoins)
	boughtShr := types.ShrpToShr(shrp, k.GetExchangeRateD(ctx))

	if err := k.buyShr(ctx, boughtShr.Amount, msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "load fee %+v shr with %+v", boughtShr, shrp)
	}

	return &types.MsgLoadFeeResponse{}, nil
}
