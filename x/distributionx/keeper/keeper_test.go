package keeper_test

import (
	"testing"

	"github.com/sharering/shareledger/app/apptesting"
	"github.com/sharering/shareledger/x/distributionx/keeper"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
	dKeeper     keeper.Keeper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)
	s.dKeeper = s.App.DistributionxKeeper
	// set dev_pool_account
	params := s.dKeeper.GetParams(s.Ctx)
	params.DevPoolAccount = s.TestAccs[0].String()
	s.dKeeper.SetParams(s.Ctx, params)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
