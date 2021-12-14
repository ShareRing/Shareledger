package keeper_test

import (
	"github.com/ShareRing/Shareledger/testutil/sample"
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
		addr, _ := sdk.AccAddressFromBech32(sample.AccAddress())
		items[i].Key = string(types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner))

		keeper.SetAccState(ctx, items[i])
	}
	return items
}

func TestAccStateGet(t *testing.T) {
	k, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetAccState(ctx,
			types.IndexKeyAccState(item.Key),
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestAccStateRemove(t *testing.T) {
	k, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(k, ctx, 10)
	for _, item := range items {
		k.RemoveAccState(ctx,
			types.IndexKeyAccState(item.Key),
		)
		_, found := k.GetAccState(ctx,
			types.IndexKeyAccState(item.Key),
		)
		require.False(t, found)
	}
}

func TestAccStateGetAll(t *testing.T) {
	k, ctx := keepertest.ElectoralKeeper(t)
	items := createNAccState(k, ctx, 10)
	require.ElementsMatch(t, items, k.GetAllAccState(ctx))
}
