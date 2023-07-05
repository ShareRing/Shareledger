package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/distributionx/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	params := types.NewParams()
	k.paramstore.GetParamSetIfExists(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
