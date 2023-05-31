package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestCompleteBatch_ok() {
	reqs := s.createNRequest(normalRequest, 1)
	batches := s.createNBatch([]uint64{reqs[0].GetId()}, "pending", 1)
	err := s.swapKeeper.MoveRequest(s.ctx, "pending", "approved", []types.Request{reqs[0]}, nil, true)
	s.Require().NoError(err)

	req := types.NewMsgCompleteBatch("", batches[0].GetId())
	_, found := s.swapKeeper.GetBatch(s.ctx, 0)
	s.Require().True(found)
	_, err = s.msgServer.CompleteBatch(sdk.WrapSDKContext(s.ctx), req)
	s.Require().NoError(err)
	_, found = s.swapKeeper.GetBatch(s.ctx, 0)
	s.Require().False(found)

}

func (s *KeeperTestSuite) TestCompleteBatch_batchNotFound() {
	reqs := s.createNRequest(normalRequest, 1)
	batches := s.createNBatch([]uint64{reqs[0].GetId()}, "pending", 1)

	req := types.NewMsgCompleteBatch("", batches[0].GetId())
	_, err := s.msgServer.CompleteBatch(sdk.WrapSDKContext(s.ctx), req)
	s.Require().NotEqual(nil, err)
}

func (s *KeeperTestSuite) TestCompleteBatch_requestStatusNotApproved() {
	reqs := s.createNRequest(normalRequest, 1)
	batches := s.createNBatch([]uint64{reqs[0].GetId()}, "pending", 1)

	req := types.NewMsgCompleteBatch("", batches[0].GetId())
	_, err := s.msgServer.CompleteBatch(sdk.WrapSDKContext(s.ctx), req)
	s.Require().Error(err)
}
