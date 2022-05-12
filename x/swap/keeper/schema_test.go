package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Schema {
	items := make([]types.Schema, n)
	for i := range items {
		items[i].Network = strconv.Itoa(i)

		keeper.SetSchema(ctx, items[i])
	}
	return items
}

func TestFormatGet(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSchema(ctx,
			item.Network,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFormatRemove(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSchema(ctx,
			item.Network,
		)
		_, found := keeper.GetSchema(ctx,
			item.Network,
		)
		require.False(t, found)
	}
}

func TestFormatGetAll(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSchema(ctx)),
	)
}
