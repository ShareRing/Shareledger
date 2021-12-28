package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/x/document/keeper"
	"github.com/sharering/shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DocumentKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
