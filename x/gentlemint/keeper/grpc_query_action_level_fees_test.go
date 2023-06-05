package keeper_test

import (
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestActionLevelFees() {
	items := s.createNActionLevelFee(5)

	// nil request
	_, err := s.gKeeper.ActionLevelFees(s.ctx, nil)
	s.Require().Error(err)

	// normal request
	resp, err := s.gKeeper.ActionLevelFees(s.ctx, &types.QueryActionLevelFeesRequest{})
	s.Require().NoError(err)
	// only check the actions fee we just created, ignore the default one from get from fee_table
	s.Require().ElementsMatch(items, resp.ActionLevelFee[:5])
}

func (s *KeeperTestSuite) TestActionLevelFee() {
	items := s.createNActionLevelFee(1)

	_, err := s.gKeeper.ActionLevelFee(s.ctx, nil)
	s.Require().Error(err)

	resp, err := s.gKeeper.ActionLevelFee(s.ctx, &types.QueryActionLevelFeeRequest{Action: "0"})
	s.Require().NoError(err)
	s.Require().Equal(items[0].Action, *&resp.Action)
}
