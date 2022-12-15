package distributionx_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RewardList: []types.Reward{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BuilderCountList: []types.BuilderCount{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		BuilderListList: []types.BuilderList{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BuilderListCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DistributionxKeeper(t)
	distributionx.InitGenesis(ctx, *k, genesisState)
	got := distributionx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RewardList, got.RewardList)
	require.ElementsMatch(t, genesisState.BuilderCountList, got.BuilderCountList)
	require.ElementsMatch(t, genesisState.BuilderListList, got.BuilderListList)
	require.Equal(t, genesisState.BuilderListCount, got.BuilderListCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
