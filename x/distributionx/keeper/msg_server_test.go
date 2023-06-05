package keeper_test

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) SetupMsgServer() (types.MsgServer, context.Context) {
	return keeper.NewMsgServerImpl(*s.dKeeper), sdk.WrapSDKContext(s.ctx)
}
