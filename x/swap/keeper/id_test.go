package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Id {
	items := make([]types.Id, n)
	for i := range items {
		items[i].IDType = strconv.Itoa(i)

		keeper.NextId(ctx, items[i].IDType)
	}
	return items
}

func TestIdNext(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	for i := 0; i < 10; i++ {
		nextV := keeper.NextId(ctx, "TypeID")
		require.Equal(t, i+1, nextV)
	}
}
