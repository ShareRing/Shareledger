package keeper_test

import (
	"testing"

    "github.com/sharering/shareledger/x/swap/keeper"
    "github.com/sharering/shareledger/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/stretchr/testify/require"
)

func createNBatch(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Batch {
	items := make([]types.Batch, n)
	for i := range items {
		items[i].Id = keeper.AppendBatch(ctx, items[i])
	}
	return items
}

func TestBatchGet(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNBatch(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBatch(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBatchRemove(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNBatch(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBatch(ctx, item.Id)
		_, found := keeper.GetBatch(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBatchGetAll(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNBatch(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBatch(ctx)),
	)
}

func TestBatchCount(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNBatch(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBatchCount(ctx))
}
