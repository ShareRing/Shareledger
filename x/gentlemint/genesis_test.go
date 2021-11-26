package gentlemint_test

import (
	"testing"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/gentlemint"
	"github.com/ShareRing/Shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GentlemintKeeper(t)
	gentlemint.InitGenesis(ctx, *k, genesisState)
	got := gentlemint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
