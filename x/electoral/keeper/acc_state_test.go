package keeper_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/sharering/shareledger/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccState(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccState {
	items := make([]types.AccState, n)
	for i := range items {
		addr, _ := sdk.AccAddressFromBech32(sample.AccAddress())
		items[i].Key = string(types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner))
		items[i].Address = addr.String()
		keeper.SetAccState(ctx, items[i])
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Key < items[j].Key
	})
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
