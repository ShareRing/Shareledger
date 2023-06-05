package keeper_test

import (
	"strconv"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) createNActionLevelFee(n int) []types.ActionLevelFee {
	items := make([]types.ActionLevelFee, n)
	for i := range items {
		items[i].Action = strconv.Itoa(i)

		s.createActionLevelFee(items[i])
	}
	return items
}

func (s *KeeperTestSuite) createActionLevelFee(input types.ActionLevelFee) {
	s.gKeeper.SetActionLevelFee(s.ctx, input)
}

func (s *KeeperTestSuite) TestActionLevelFeeGet() {
	items := s.createNActionLevelFee(10)
	for _, item := range items {
		rst, found := s.gKeeper.GetActionLevelFee(s.ctx,
			item.Action,
		)
		s.Require().True(found)
		s.Require().Equal(item, rst)
	}
}
func (s *KeeperTestSuite) TestActionLevelFeeRemove() {
	items := s.createNActionLevelFee(10)
	for _, item := range items {
		s.gKeeper.RemoveActionLevelFee(s.ctx,
			item.Action,
		)
		_, found := s.gKeeper.GetActionLevelFee(s.ctx,
			item.Action,
		)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestActionLevelFeeGetAll() {
	items := s.createNActionLevelFee(10)
	s.Require().ElementsMatch(items, s.gKeeper.GetAllActionLevelFee(s.ctx))
}
