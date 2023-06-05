package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestSetExchange() {
	req := types.MsgSetExchange{
		Creator: testAcc,
		Rate:    "777",
	}

	_, err := s.msgServer.SetExchange(sdk.WrapSDKContext(s.ctx), &req)
	s.Require().NoError(err)

	// query exchange rate and check
	resp, found := s.gKeeper.GetExchangeRate(s.ctx)
	s.Require().True(found)
	s.Require().Equal(resp.Rate, req.Rate)
}
