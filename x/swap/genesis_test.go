package swap_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/swap"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdList: []types.Id{
			{
				IDType: "0",
			},
			{
				IDType: "1",
			},
		},
		RequestList: []types.Request{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RequestCount: 2,
		BatchList: []types.Batch{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BatchCount: 2,
		FormatList: []types.Format{
		{
			Network: "0",
},
		{
			Network: "1",
},
	},
	// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwapKeeper(t)
	swap.InitGenesis(ctx, *k, genesisState)
	got := swap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdList, got.IdList)
	require.ElementsMatch(t, genesisState.RequestList, got.RequestList)
	require.Equal(t, genesisState.RequestCount, got.RequestCount)
	require.ElementsMatch(t, genesisState.BatchList, got.BatchList)
	require.Equal(t, genesisState.BatchCount, got.BatchCount)
	require.ElementsMatch(t, genesisState.FormatList, got.FormatList)
// this line is used by starport scaffolding # genesis/test/assert
}
