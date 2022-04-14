package keeper_test

import (
	"testing"

	testkeeper "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.OutFee, k.OutFee(ctx))
}
