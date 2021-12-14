package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/keeper"
	"github.com/ShareRing/Shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccState(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccState {
	items := make([]types.AccState, n)
	for i := range items {
		items[i].Key = strconv.Itoa(i)

		keeper.SetAccState(ctx, items[i])
	}
	return items
}

func TestAccStateGet(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccState(ctx,
			item.Key,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestAccStateRemove(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccState(ctx,
			item.Key,
		)
		_, found := keeper.GetAccState(ctx,
			item.Key,
		)
		require.False(t, found)
	}
}

func TestAccStateGetAll(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllAccState(ctx))
}
