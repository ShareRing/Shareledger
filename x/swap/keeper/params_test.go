package keeper_test

import "github.com/sharering/shareledger/x/swap/types"

func (s *KeeperTestSuite) TestGetParams() {
	params := types.DefaultParams()

	s.swapKeeper.SetParams(s.ctx, params)

	s.Require().EqualValues(params, s.swapKeeper.GetParams(s.ctx))
}
