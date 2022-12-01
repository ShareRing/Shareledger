package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/sdistribution/keeper"
	"github.com/sharering/shareledger/x/sdistribution/types"
	"github.com/stretchr/testify/require"
)

func createNBuilderList(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BuilderList {
	items := make([]types.BuilderList, n)
	for i := range items {
		items[i].Id = keeper.AppendBuilderList(ctx, items[i])
	}
	return items
}

func TestBuilderListGet(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNBuilderList(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetBuilderList(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestBuilderListRemove(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNBuilderList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBuilderList(ctx, item.Id)
		_, found := keeper.GetBuilderList(ctx, item.Id)
		require.False(t, found)
	}
}

func TestBuilderListGetAll(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNBuilderList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBuilderList(ctx)),
	)
}

func TestBuilderListCount(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNBuilderList(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetBuilderListCount(ctx))
}
