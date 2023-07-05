package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k Keeper) GetMinGasPriceParam(ctx sdk.Context) sdk.DecCoins {
	minGasPrice := sdk.DecCoins{}
	k.paramSpace.Get(ctx, types.ParamStoreKeyMinGasPrices, &minGasPrice)
	return minGasPrice
}

// set the params
func (k Keeper) SetMinGasPriceParam(ctx sdk.Context, min sdk.DecCoins) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyMinGasPrices, min)
}
