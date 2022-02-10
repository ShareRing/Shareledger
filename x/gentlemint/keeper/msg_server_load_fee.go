package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// LoadFee Not use this function
// This func is executed in antee
func (k msgServer) LoadFee(goCtx context.Context, msg *types.MsgLoadFee) (*types.MsgLoadFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.LoadFeeFundFromShrp(ctx, msg)
	return &types.MsgLoadFeeResponse{}, err
}
