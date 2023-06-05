package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (s *KeeperTestSuite) TestParams() {
	s.gKeeper.SetMinGasPriceParam(s.ctx, sdk.NewDecCoins(tenDecCoin))

	resp := s.gKeeper.GetMinGasPriceParam(s.ctx)
	s.Require().Equal(resp[0], tenDecCoin)
}
