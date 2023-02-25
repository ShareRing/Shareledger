package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/sharering/shareledger/app/apptesting"
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
	gKeeper     *keeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)
	s.gKeeper = &s.App.GentleMintKeeper
}

func (s *KeeperTestSuite) TestBaseMintPossible_false() {
	s.False(s.gKeeper.BaseMintPossible(s.Ctx, types.MaxBaseSupply))

}

func (s *KeeperTestSuite) TestBaseMintPossible_true() {
	s.True(s.gKeeper.BaseMintPossible(s.Ctx, math.NewInt(1000)))
}

func (s *KeeperTestSuite) TestLoadAllowanceLoader_increaseBalance() {
	coinBefore := s.App.BankKeeper.GetBalance(s.Ctx, s.TestAccs[0], denom.Base)
	s.Nil(s.gKeeper.LoadAllowanceLoader(s.Ctx, s.TestAccs[0]))
	coinAfter := s.App.BankKeeper.GetBalance(s.Ctx, s.TestAccs[0], denom.Base)
	s.Equal(coinBefore.Add(types.AllowanceLoader[0]), coinAfter)
}
