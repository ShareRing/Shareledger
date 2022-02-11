package keeper

import (
	"context"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) BuyPShr(goCtx context.Context, msg *types.MsgBuyPShr) (*types.MsgBuyPShrResponse, error) {
	panic("implement me")
	//ctx := sdk.UnwrapSDKContext(goCtx)
	//
	//if err := msg.ValidateBasic(); err != nil {
	//	return nil, err
	//}
	//coins, err := types.ParsePShrCoinsStr(msg.Amount)
	//if err != nil {
	//	return nil, err
	//}
	//if err := k.buyPShr(ctx, coins.AmountOf(denom.PShr), msg.GetSigners()[0]); err != nil {
	//	return nil, sdkerrors.Wrapf(err, "buy %v pshr to %v", msg.Amount, msg.Creator)
	//}
	//return &types.MsgBuyPShrResponse{
	//	Log: fmt.Sprintf("Successfull buy %v pshr for address %s", msg.Amount, msg.Creator),
	//}, nil
}
