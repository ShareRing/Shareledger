package asset_test

import (
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/asset/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetKeeper(t)
	asset.InitGenesis(ctx, *k, genesisState)
	got := asset.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
