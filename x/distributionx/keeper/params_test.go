package keeper_test

import (
	"testing"

	testkeeper "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DistributionxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
