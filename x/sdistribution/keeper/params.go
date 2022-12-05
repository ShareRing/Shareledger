package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/sdistribution/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	fmt.Printf("get params: %v\n", params)

	k.paramstore.SetParamSet(ctx, &types.Params{
		WasmMasterBuilder: 1000,
	})

	k.paramstore.GetParamSet(ctx, &params)
	fmt.Printf("get params: %v\n", params)

	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	fmt.Printf("set params: %v\n", params)
	k.paramstore.SetParamSet(ctx, &params)
}
