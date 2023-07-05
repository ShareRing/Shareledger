package document_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/document"
	"github.com/sharering/shareledger/x/document/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DocumentKeeper(t)
	document.InitGenesis(ctx, *k, genesisState)
	got := document.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
