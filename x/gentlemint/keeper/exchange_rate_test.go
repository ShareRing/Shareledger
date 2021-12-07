package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/gentlemint/keeper"
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
)

func createTestExchangeRate(keeper *keeper.Keeper, ctx sdk.Context) types.ExchangeRate {
	item := types.ExchangeRate{}
	keeper.SetExchangeRate(ctx, item)
	return item
}

func TestExchangeRateGet(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	item := createTestExchangeRate(keeper, ctx)
	rst, found := keeper.GetExchangeRate(ctx)
	require.True(t, found)
	require.Equal(t, item, rst)
}
func TestExchangeRateRemove(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	createTestExchangeRate(keeper, ctx)
	keeper.RemoveExchangeRate(ctx)
	_, found := keeper.GetExchangeRate(ctx)
	require.False(t, found)
}
