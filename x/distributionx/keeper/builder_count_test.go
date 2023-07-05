package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBuilderCount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BuilderCount {
	items := make([]types.BuilderCount, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetBuilderCount(ctx, items[i])
	}
	return items
}

func TestBuilderCountGet(t *testing.T) {
	keeper, ctx := keepertest.DistributionxKeeper(t)
	items := createNBuilderCount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBuilderCount(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBuilderCountRemove(t *testing.T) {
	keeper, ctx := keepertest.DistributionxKeeper(t)
	items := createNBuilderCount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBuilderCount(ctx,
			item.Index,
		)
		_, found := keeper.GetBuilderCount(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestBuilderCountGetAll(t *testing.T) {
	keeper, ctx := keepertest.DistributionxKeeper(t)
	items := createNBuilderCount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBuilderCount(ctx)),
	)
}
