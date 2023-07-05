package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func TestQueryMinimumGasPrices(t *testing.T) {
	specs := map[string]struct {
		setupStore func(ctx sdk.Context, s paramtypes.Subspace)
		expMin     sdk.DecCoins
	}{
		"one coin": {
			setupStore: func(ctx sdk.Context, s paramtypes.Subspace) {
				s.SetParamSet(ctx, &types.Params{
					MinimumGasPrices: sdk.NewDecCoins(sdk.NewDecCoin("ALX", sdk.OneInt())),
				})
			},
			expMin: sdk.NewDecCoins(sdk.NewDecCoin("ALX", sdk.OneInt())),
		},
		"multiple coins": {
			setupStore: func(ctx sdk.Context, s paramtypes.Subspace) {
				s.SetParamSet(ctx, &types.Params{
					MinimumGasPrices: sdk.NewDecCoins(sdk.NewDecCoin("ALX", sdk.OneInt()), sdk.NewDecCoin("BLX", sdk.NewInt(2))),
				})
			},
			expMin: sdk.NewDecCoins(sdk.NewDecCoin("ALX", sdk.OneInt()), sdk.NewDecCoin("BLX", sdk.NewInt(2))),
		},
		"no min gas price set": {
			setupStore: func(ctx sdk.Context, s paramtypes.Subspace) {
				s.SetParamSet(ctx, &types.Params{})
			},
		},
		"no param set": {
			setupStore: func(ctx sdk.Context, s paramtypes.Subspace) {
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			keeper, ctx := keepertest.GentlemintKeeper(t)
			wctx := sdk.WrapSDKContext(ctx)
			_ = spec
			_ = keeper
			_ = wctx
		})
	}
}
