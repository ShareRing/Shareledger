package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (s *KeeperTestSuite) TestQueryMinimumGasPrices() {
	s.gKeeper.SetMinGasPriceParam(s.ctx, sdk.NewDecCoins(tenDecCoin))
	resp, err := s.gKeeper.MinimumGasPrices(sdk.WrapSDKContext(s.ctx), nil)

	s.Require().NoError(err)
	s.Require().Equal(tenDecCoin, resp.MinimumGasPrices[0])
}
