package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
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
		Log: fmt.Sprintf("Successfull buy %v shr for address %s", msg.Amount, msg.Creator),
	}, nil
}
