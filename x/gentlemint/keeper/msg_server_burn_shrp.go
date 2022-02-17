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
//func (k msgServer) BurnShrp(goCtx context.Context, msg *types.MsgBurnShrp) (*types.MsgBurnShrpResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//	if err := msg.ValidateBasic(); err != nil {
//		return nil, err
//	}
//
//	amt, err := sdk.NewDecFromStr(msg.Amount)
//	if err != nil {
//		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
//	}
//	if err != nil {
//		return nil, err
//	}
//	rate := k.GetExchangeRateD(ctx)
//	bUSD, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoins(sdk.NewDecCoinFromDec(denom.ShrP, amt)), rate, true)
//
//	if err := k.burnCoins(ctx, msg.GetSigners()[0], sdk.NewCoins(bUSD)); err != nil {
//		return nil, sdkerrors.Wrapf(err, "burns %v coins from %v", bUSD, msg.Creator)
//	}
//	log := fmt.Sprintf("Successfully burn coins %s", msg.Amount)
//
//	return &types.MsgBurnShrpResponse{
//		Log: log,
//	}, nil
//}
