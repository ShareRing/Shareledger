package keeper_test

func (s *KeeperTestSuite) TestSetParams() {
	params := s.dKeeper.GetParams(s.ctx)
	params.TxThreshold = 50
	s.dKeeper.SetParams(s.ctx, params)
	params = s.dKeeper.GetParams(s.ctx)
	s.Require().Equal(uint32(50), params.TxThreshold)
}
