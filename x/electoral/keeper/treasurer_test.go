package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

func createTestTreasurer(keeper *keeper.Keeper, ctx sdk.Context) types.Treasurer {
	item := types.Treasurer{}
	keeper.SetTreasurer(ctx, item)
	return item
}

func TestTreasurerGet(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	item := createTestTreasurer(keeper, ctx)
	rst, found := keeper.GetTreasurer(ctx)
	require.True(t, found)
	require.Equal(t, item, rst)
}
func TestTreasurerRemove(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	createTestTreasurer(keeper, ctx)
	keeper.RemoveTreasurer(ctx)
	_, found := keeper.GetTreasurer(ctx)
	require.False(t, found)
}
