package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/sdistribution/keeper"
	"github.com/sharering/shareledger/x/sdistribution/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNReward(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Reward {
	items := make([]types.Reward, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetReward(ctx, items[i])
	}
	return items
}

func TestRewardGet(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNReward(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetReward(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRewardRemove(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNReward(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveReward(ctx,
			item.Index,
		)
		_, found := keeper.GetReward(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRewardGetAll(t *testing.T) {
	keeper, ctx := keepertest.SdistributionKeeper(t)
	items := createNReward(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReward(ctx)),
	)
}
