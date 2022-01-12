package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLevelFee(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LevelFee {
	items := make([]types.LevelFee, n)
	for i := range items {
		items[i].Level = strconv.Itoa(i)

		keeper.SetLevelFee(ctx, items[i])
	}
	return items
}

func TestLevelFeeGet(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNLevelFee(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLevelFee(ctx,
			item.Level,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestLevelFeeRemove(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNLevelFee(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLevelFee(ctx,
			item.Level,
		)
		_, found := keeper.GetLevelFee(ctx,
			item.Level,
		)
		require.False(t, found)
	}
}

func TestLevelFeeGetAll(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNLevelFee(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllLevelFee(ctx))
}
