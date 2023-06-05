package keeper_test

import "github.com/sharering/shareledger/x/gentlemint/keeper"

func (s *KeeperTestSuite) TestMigrate2to3() {
	migrator := keeper.NewMigrator(*s.gKeeper)

	err := migrator.Migrate2to3(s.ctx)
	s.Require().NoError(err)
	resp := s.gKeeper.GetMinGasPriceParam(s.ctx)
	s.Require().Empty(resp)
}
