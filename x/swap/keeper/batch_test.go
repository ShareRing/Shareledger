package keeper_test

import (
	"encoding/binary"

	"github.com/sharering/shareledger/x/swap/keeper"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) createNBatch(requestIds []uint64, status string, n int) []types.Batch {
	items := make([]types.Batch, n)
	for i := range items {
		if len(requestIds) != 0 {
			items[i].RequestIds = requestIds
		}
		if status != "" {
			items[i].Status = status
		}
		items[i].Id = s.swapKeeper.AppendBatch(s.ctx, items[i])
	}
	return items
}

func (s *KeeperTestSuite) TestBatchGet() {
	items := s.createNBatch([]uint64{}, "", 10)
	for _, item := range items {
		got, found := s.swapKeeper.GetBatch(s.ctx, item.Id)
		s.Require().True(found)
		s.Require().Equal(&item, &got)
	}
}

func (s *KeeperTestSuite) TestBatchRemove() {
	items := s.createNBatch([]uint64{}, "", 10)
	for _, item := range items {
		s.swapKeeper.RemoveBatch(s.ctx, item.Id)
		_, found := s.swapKeeper.GetBatch(s.ctx, item.Id)
		s.Require().False(found)
	}
}

func (s *KeeperTestSuite) TestBatchGetAll() {
	items := s.createNBatch([]uint64{}, "", 10)
	s.Require().ElementsMatch(items, s.swapKeeper.GetAllBatch(s.ctx))
}

func (s *KeeperTestSuite) TestBatchCount() {
	items := s.createNBatch([]uint64{}, "", 10)
	count := uint64(len(items))
	s.Require().Equal(count, s.swapKeeper.GetBatchCount(s.ctx))
}

func (s *KeeperTestSuite) TestGetBatchesByIDs() {
	items := s.createNBatch([]uint64{}, "", 3)
	batch := s.swapKeeper.GetBatchesByIDs(s.ctx, []uint64{0, 1, 2})
	s.Require().Equal(len(items), len(batch))
}

func (s *KeeperTestSuite) TestGetBatchIDBytes() {
	_ = s.createNBatch([]uint64{}, "", 3)
	batchId := keeper.GetBatchIDBytes(1)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, 1)
	s.Require().Equal(batchId, bz)
}
