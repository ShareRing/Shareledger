package keeper_test

import (
	"strconv"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) createNLevelFee(n int) []types.LevelFee {
	items := make([]types.LevelFee, n)
	for i := range items {
		items[i].Fee = tenDecCoin
		items[i].Level = strconv.Itoa(i)
		s.createLevelFee(items[i])
	}
	return items
}

func (s *KeeperTestSuite) createLevelFee(item types.LevelFee) {
	s.gKeeper.SetLevelFee(s.ctx, item)
}

func (s *KeeperTestSuite) TestLevelFeeGet() {
	items := s.createNLevelFee(10)
	for _, item := range items {
		resp, found := s.gKeeper.GetLevelFee(s.ctx,
			item.Level,
		)
		s.Require().True(found)
		s.Require().Equal(item, resp)
	}
}
func (s *KeeperTestSuite) TestLevelFeeRemove() {
	items := s.createNLevelFee(10)
	for _, item := range items {
		s.gKeeper.RemoveLevelFee(s.ctx,
			item.Level,
		)
		_, found := s.gKeeper.GetLevelFee(s.ctx,
			item.Level,
		)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestLevelFeeGetAll() {
	items := s.createNLevelFee(10)
	resp := s.gKeeper.GetAllLevelFee(s.ctx)
	s.Require().ElementsMatch(items, resp)
}
