package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/require"
)

func createNRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Request {
	items := make([]types.Request, n)
	for i := range items {
		m, err := keeper.AppendPendingRequest(ctx, items[i])
		if err != nil {
			panic(err)
		}
		items[i].Id = m.Id
	}
	return items
}

func TestRequestCount(t *testing.T) {
	keeper, ctx := keepertest.SwapKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRequestCount(ctx))
}
