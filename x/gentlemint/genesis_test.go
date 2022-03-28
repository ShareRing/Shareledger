package gentlemint_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/gentlemint"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ExchangeRate: &types.ExchangeRate{
			Rate: "200",
		},
		LevelFeeList: []types.LevelFee{
			{
				Level: "0",
			},
			{
				Level: "1",
			},
		},
		ActionLevelFeeList: []types.ActionLevelFee{
			{
				Action: "0",
			},
			{
				Action: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GentlemintKeeper(t)
	gentlemint.InitGenesis(ctx, *k, genesisState)
	got := gentlemint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.ExchangeRate, got.ExchangeRate)
	require.Len(t, got.LevelFeeList, len(genesisState.LevelFeeList))
	require.Subset(t, genesisState.LevelFeeList, got.LevelFeeList)
	require.Len(t, got.ActionLevelFeeList, len(genesisState.ActionLevelFeeList))
	require.Subset(t, genesisState.ActionLevelFeeList, got.ActionLevelFeeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
