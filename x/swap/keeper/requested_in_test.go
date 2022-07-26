package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRequestedIn(keeper *keeper.Keeper, ctx sdk.Context, n int) []string {
	items := make([]string, n)
	for i := 0; i <= n; i++ {
		keeper.SetRequestedIn(ctx, sdk.AccAddress{}, "", []string{strconv.Itoa(i)})
		items[i] = strconv.Itoa(i)
	}
	return items
}

func TestRequestedInGet(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequestedIn(ctx,
			item,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRequestedInRemove(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequestedIn(ctx,
			item,
		)
		_, found := keeper.GetRequestedIn(ctx,
			item,
		)
		require.False(t, found)
	}
}

func TestRequestedInGetAll(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRequestedIn(ctx)),
	)
}
