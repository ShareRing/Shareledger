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
	shrCoin := sdk.NewDecCoinFromDec(denom.Shr, sdk.MustNewDecFromStr(msg.Amount))
	rate := k.GetExchangeRateD(ctx)
	coin, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(shrCoin), rate, false)
	if err != nil {
		return nil, err
	}
	if err := k.buyBaseDenom(ctx, coin, msg.GetSigners()[0]); err != nil {
		return nil, sdkerrors.Wrapf(err, "buy %+v to %v", coin, msg.Creator)
	}
	return &types.MsgBuyShrResponse{
		Log: fmt.Sprintf("Successfull buy %+v for address %s", coin, msg.Creator),
	}, nil
}
