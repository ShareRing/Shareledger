package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
)

func createTestAuthority(keeper *keeper.Keeper, ctx sdk.Context) types.Authority {
	item := types.Authority{}
	keeper.SetAuthority(ctx, item)
	return item
}

func TestAuthorityGet(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	item := createTestAuthority(keeper, ctx)
	rst, found := keeper.GetAuthority(ctx)
	require.True(t, found)
	require.Equal(t, item, rst)
}
func TestAuthorityRemove(t *testing.T) {
	keeper, ctx := keepertest.ElectoralKeeper(t)
	createTestAuthority(keeper, ctx)
	keeper.RemoveAuthority(ctx)
	_, found := keeper.GetAuthority(ctx)
	require.False(t, found)
}
