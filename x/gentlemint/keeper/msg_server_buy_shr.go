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
	coin, err := denom.NormalizeCoins(sdk.NewDecCoins(shrCoin), nil)
	if err != nil {
		return nil, err
	}
	if err := k.buyBaseDenom(ctx, coin.Amount, msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "buy %+v to %v", coin, msg.Creator)
	}
	return &types.MsgBuyShrResponse{
		Log: fmt.Sprintf("Successfull buy %+v for address %s", coin, msg.Creator),
	}, nil
}
