package keeper_test

import (
	"strconv"

	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) createNReward(n int) []types.Reward {
	items := make([]types.Reward, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		s.dKeeper.SetReward(s.ctx, items[i])
	}
	return items
}

func (s *KeeperTestSuite) TestGetReward() {
	items := s.createNReward(10)
	for _, item := range items {
		rst, found := s.dKeeper.GetReward(s.ctx,
			item.Index,
		)
		s.Require().True(found)
		s.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func (s *KeeperTestSuite) TestRemoveReward() {
	items := s.createNReward(10)
	for _, item := range items {
		s.dKeeper.RemoveReward(s.ctx,
			item.Index,
		)
		_, found := s.dKeeper.GetReward(s.ctx,
			item.Index,
		)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestGetAllReward() {
	items := s.createNReward(10)
	s.Require().ElementsMatch(
		items,
		s.dKeeper.GetAllReward(s.ctx),
	)
}
