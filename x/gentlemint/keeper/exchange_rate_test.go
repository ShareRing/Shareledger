package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) createTestExchangeRate() types.ExchangeRate {
	item := types.ExchangeRate{
		Rate: "200.1",
	}
	s.gKeeper.SetExchangeRate(s.ctx, item)
	return item
}

func (s *KeeperTestSuite) TestGetExchangeRate() {
	item := s.createTestExchangeRate()
	rst, found := s.gKeeper.GetExchangeRate(s.ctx)
	s.Require().True(found)
	s.Require().Equal(item, rst)
}

func (s *KeeperTestSuite) TestRemoveExchangeRate() {
	s.createTestExchangeRate()
	s.gKeeper.RemoveExchangeRate(s.ctx)
	v, found := s.gKeeper.GetExchangeRate(s.ctx)
	s.Require().True(found)
	s.Require().Equal(types.DefaultExchangeRateSHRPToSHR, sdk.MustNewDecFromStr(v.Rate))
}

func (s *KeeperTestSuite) TestGetExchangeRateD() {
	item := s.createTestExchangeRate()
	resp := s.gKeeper.GetExchangeRateD(s.ctx)
	s.Require().Equal(sdk.MustNewDecFromStr(item.Rate), resp)
}
