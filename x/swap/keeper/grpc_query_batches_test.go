package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestBatches_ok() {
	items := s.createNBatch([]uint64{}, "", 3)
	req := &types.QueryBatchesRequest{
		Network:    "",
		Ids:        []uint64{0, 1, 2},
		Pagination: &query.PageRequest{},
	}
	resp, err := s.swapKeeper.Batches(s.ctx, req)
	s.Require().NoError(err)
	numBatchReturn := len(resp.Batches)
	s.Require().Equal(3, numBatchReturn)
	for index, item := range items {
		s.Require().Equal(item, resp.Batches[index])
	}
}

func (s *KeeperTestSuite) TestBatches_wrongBatcheId() {
	items := s.createNBatch([]uint64{}, "", 3)
	req := &types.QueryBatchesRequest{
		Network:    "",
		Ids:        []uint64{1, 2, 3},
		Pagination: &query.PageRequest{},
	}
	resp, err := s.swapKeeper.Batches(s.ctx, req)
	s.Require().NoError(err)
	numBatchReturn := len(resp.Batches)
	expectedNumber := 2
	s.Require().Equal(expectedNumber, numBatchReturn)
	for index, item := range resp.Batches {
		s.Require().Equal(items[index+1], item)
	}
}
