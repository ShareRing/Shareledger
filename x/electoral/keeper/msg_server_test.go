package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ElectoralKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
