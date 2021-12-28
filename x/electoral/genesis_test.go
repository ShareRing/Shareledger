package electoral_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/electoral"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		AccStateList: []types.AccState{
			{
				Key: "0",
			},
			{
				Key: "1",
			},
		},
		Authority: &types.Authority{
			Address: "address",
		},
		Treasurer: &types.Treasurer{
			Address: "address",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ElectoralKeeper(t)
	electoral.InitGenesis(ctx, *k, genesisState)
	got := electoral.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.AccStateList, len(genesisState.AccStateList))
	require.Subset(t, genesisState.AccStateList, got.AccStateList)
	require.Equal(t, genesisState.Authority, got.Authority)
	require.Equal(t, genesisState.Treasurer, got.Treasurer)
	// this line is used by starport scaffolding # genesis/test/assert
}
