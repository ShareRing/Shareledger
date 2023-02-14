package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (k Keeper) MinimumGasPrices(stdCtx context.Context, _ *types.QueryMinimumGasPricesRequest) (*types.QueryMinimumGasPricesResponse, error) {
	var minGasPrices sdk.DecCoins
	ctx := sdk.UnwrapSDKContext(stdCtx)
	if k.paramsSpace.Has(ctx, types.ParamStoreKeyMinGasPrices) {
		k.paramsSpace.Get(ctx, types.ParamStoreKeyMinGasPrices, &minGasPrices)
	}
	return &types.QueryMinimumGasPricesResponse{
		MinimumGasPrices: minGasPrices,
	}, nil
}
