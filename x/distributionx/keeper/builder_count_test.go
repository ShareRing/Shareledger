package keeper_test

import (
	"strconv"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) createNBuilderCount(n int) []types.BuilderCount {
	items := make([]types.BuilderCount, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		s.dKeeper.SetBuilderCount(s.ctx, items[i])
	}
	return items
}

func (s *KeeperTestSuite) TestBuilderCountGet() {
	items := s.createNBuilderCount(10)
	for _, item := range items {
		resp, found := s.dKeeper.GetBuilderCount(s.ctx,
			item.Index,
		)
		s.Require().True(found)
		s.Require().Equal(&item, &resp)
	}
}

func (s *KeeperTestSuite) TestBuilderCountRemove() {
	items := s.createNBuilderCount(10)
	for _, item := range items {
		s.dKeeper.RemoveBuilderCount(s.ctx,
			item.Index,
		)
		_, found := s.dKeeper.GetBuilderCount(s.ctx,
			item.Index,
		)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestBuilderCountGetAll() {
	items := s.createNBuilderCount(10)
	s.Require().ElementsMatch(items, s.dKeeper.GetAllBuilderCount(s.ctx))
}

func (s *KeeperTestSuite) TestIncBuilderCount() {
	items := s.createNBuilderCount(3)
	s.dKeeper.IncBuilderCount(s.ctx, "test")
	_, found := s.dKeeper.GetBuilderCount(s.ctx, "test")
	s.Require().True(found)
	resp := s.dKeeper.GetAllBuilderCount(s.ctx)
	s.Require().Equal(len(items)+1, len(resp))
}
