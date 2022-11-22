package sdistribution_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/sdistribution"
	"github.com/sharering/shareledger/x/sdistribution/types"
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SdistributionKeeper(t)
	sdistribution.InitGenesis(ctx, *k, genesisState)
	got := sdistribution.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RewardList, got.RewardList)
	// this line is used by starport scaffolding # genesis/test/assert
}
