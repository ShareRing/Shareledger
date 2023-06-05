package keeper_test

import "github.com/sharering/shareledger/testutil/nullify"

func (s *KeeperTestSuite) TestFormatGet() {
	items := s.createNSchema(10)
	for _, item := range items {
		rst, found := s.swapKeeper.GetSchema(s.ctx,
			item.Network,
		)
		s.Require().True(found)
		s.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func (s *KeeperTestSuite) TestFormatRemove() {
	items := s.createNSchema(10)
	for _, item := range items {
		s.swapKeeper.RemoveSchema(s.ctx,
			item.Network,
		)
		_, found := s.swapKeeper.GetSchema(s.ctx,
			item.Network,
		)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestFormatGetAll() {
	items := s.createNSchema(10)
	s.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(s.swapKeeper.GetAllSchema(s.ctx)),
	)
}
