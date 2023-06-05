package keeper_test

import (
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) createNBuilderList(n int) []types.BuilderList {
	items := make([]types.BuilderList, n)
	for i := range items {
		items[i].Id = s.dKeeper.AppendBuilderList(s.ctx, items[i])
	}
	return items
}

func (s *KeeperTestSuite) TestBuilderListGet() {
	items := s.createNBuilderList(10)
	for _, item := range items {
		got, found := s.dKeeper.GetBuilderList(s.ctx, item.Id)
		s.Require().True(found)
		s.Require().Equal(&item, &got)
	}
}

func (s *KeeperTestSuite) TestBuilderListRemove() {
	items := s.createNBuilderList(10)
	for _, item := range items {
		s.dKeeper.RemoveBuilderList(s.ctx, item.Id)
		_, found := s.dKeeper.GetBuilderList(s.ctx, item.Id)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestBuilderListGetAll() {
	items := s.createNBuilderList(10)
	s.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(s.dKeeper.GetAllBuilderList(s.ctx)),
	)
}

func (s *KeeperTestSuite) TestBuilderListCount() {
	items := s.createNBuilderList(10)
	count := uint64(len(items))
	s.Require().Equal(count, s.dKeeper.GetBuilderListCount(s.ctx))
}
