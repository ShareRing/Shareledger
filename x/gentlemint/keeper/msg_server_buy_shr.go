package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

func (k msgServer) BuyShr(goCtx context.Context, msg *types.MsgBuyShr) (*types.MsgBuyShrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	shrCoin := sdk.NewDecCoinFromDec(denom.Shr, sdk.MustNewDecFromStr(fmt.Sprintf(msg.Amount)))
	coin := denom.NormalizeCoins(sdk.NewDecCoins(shrCoin), sdk.NewDec(1))
	if err := k.buyPShr(ctx, coin.Amount, msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "buy %v pshr to %v", msg.Amount, msg.Creator)
	}
	return &types.MsgBuyShrResponse{
		Log: fmt.Sprintf("Successfull buy %v pshr for address %s", msg.Amount, msg.Creator),
	}, nil
}
