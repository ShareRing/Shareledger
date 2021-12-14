package electoral_test

import (
	"testing"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/electoral"
	"github.com/ShareRing/Shareledger/x/electoral/types"
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ElectoralKeeper(t)
	electoral.InitGenesis(ctx, *k, genesisState)
	got := electoral.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.AccStateList, len(genesisState.AccStateList))
	require.Subset(t, genesisState.AccStateList, got.AccStateList)
	// this line is used by starport scaffolding # genesis/test/assert
}
