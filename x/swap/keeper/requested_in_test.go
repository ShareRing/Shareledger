package keeper_test

import (
	"strconv"
	"testing"

	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRequestedIn(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RequestedIn {
	items := make([]types.RequestedIn, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
        
		keeper.SetRequestedIn(ctx, items[i])
	}
	return items
}

func TestRequestedInGet(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequestedIn(ctx,
		    item.Address,
            
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRequestedInRemove(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequestedIn(ctx,
		    item.Address,
            
		)
		_, found := keeper.GetRequestedIn(ctx,
		    item.Address,
            
		)
		require.False(t, found)
	}
}

func TestRequestedInGetAll(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequestedIn(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRequestedIn(ctx)),
	)
}
