package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.OutFee(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// OutFee returns the OutFee param
func (k Keeper) OutFee(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyOutFee, &res)
	return
}
