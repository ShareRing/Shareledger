package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k msgServer) LoadFee(goCtx context.Context, msg *types.MsgLoadFee) (*types.MsgLoadFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.LoadFeeFromShrp(ctx, msg)
	return &types.MsgLoadFeeResponse{}, err
}
