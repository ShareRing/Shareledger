package electoralbk_test

import (
	"testing"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/electoralbk"
	"github.com/ShareRing/Shareledger/x/electoralbk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.electoralbkKeeper(t)
	electoralbk.InitGenesis(ctx, *k, genesisState)
	got := electoralbk.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
