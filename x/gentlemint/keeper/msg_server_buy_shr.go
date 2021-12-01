package keeper

import (
	"context"
	"fmt"

	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyShr(goCtx context.Context, msg *types.MsgBuyShr) (*types.MsgBuyShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	coins, err := types.ParseShrCoinsStr(msg.Amount)
	if err != nil {
		return nil, err
	}
	if err := k.buyShr(ctx, coins.AmountOf(types.DenomSHR), msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "buy %v shr to %v", msg.Amount, msg.Creator)
	}
	return &types.MsgBuyShrResponse{
		Log: fmt.Sprintf("Successfull buy %d shr for address %s", msg.Amount, msg.Creator),
	}, nil
}

func (k msgServer) buyShr(ctx sdk.Context, amount sdk.Int, buyer sdk.AccAddress) error {
	if !k.ShrMintPossible(ctx, amount) {
		return sdkerrors.Wrap(types.ErrSHRSupplyExceeded, amount.String())
	}
}
