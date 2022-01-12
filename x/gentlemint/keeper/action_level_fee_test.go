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

func createNActionLevelFee(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionLevelFee {
	items := make([]types.ActionLevelFee, n)
	for i := range items {
		items[i].Action = strconv.Itoa(i)

		keeper.SetActionLevelFee(ctx, items[i])
	}
	return items
}

func TestActionLevelFeeGet(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNActionLevelFee(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionLevelFee(ctx,
			item.Action,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestActionLevelFeeRemove(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNActionLevelFee(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionLevelFee(ctx,
			item.Action,
		)
		_, found := keeper.GetActionLevelFee(ctx,
			item.Action,
		)
		require.False(t, found)
	}
}

func TestActionLevelFeeGetAll(t *testing.T) {
	keeper, ctx := keepertest.GentlemintKeeper(t)
	items := createNActionLevelFee(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllActionLevelFee(ctx))
}
