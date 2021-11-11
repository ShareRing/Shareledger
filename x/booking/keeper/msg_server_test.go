package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ShareRing/Shareledger/testutil/keeper"
	"github.com/ShareRing/Shareledger/x/booking/keeper"
	"github.com/ShareRing/Shareledger/x/booking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BookingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
