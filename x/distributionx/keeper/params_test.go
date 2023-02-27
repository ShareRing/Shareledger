package keeper_test

// TODO:
func (s *KeeperTestSuite) TestGetParams() {
	params := s.dKeeper.GetParams(s.Ctx)
	_ = params
}

// TODO:
func (s *KeeperTestSuite) TestSetParams() {
	params := s.dKeeper.GetParams(s.Ctx)
	s.dKeeper.SetParams(s.Ctx, params)
}
