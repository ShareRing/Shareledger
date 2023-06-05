package keeper_test

import (
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestKeeperReject() {

	pendingRequests := s.createNRequest(emptySrcRequest, 1)
	s.Require().NotEmpty(pendingRequests)

	// test from status pending move to status reject
	_, err := s.swapKeeper.RejectSwap(s.ctx, types.NewMsgReject("", []uint64{pendingRequests[0].GetId()}))
	s.Require().Error(err)

	emptySrcRequest.SrcNetwork = types.NetworkNameShareLedger
	pendingRequests = s.createNRequest(emptySrcRequest, 1)
	s.Require().NotEmpty(pendingRequests)

	// test from status pending move to status reject
	rejectedRequests, err := s.swapKeeper.RejectSwap(s.ctx, types.NewMsgReject("", []uint64{pendingRequests[0].GetId()}))
	s.Require().NoError(err)
	s.Require().Equal(types.SwapStatusRejected, rejectedRequests[0].GetStatus())
	_, found := s.swapKeeper.GetRequest(s.ctx, pendingRequests[0].GetId())
	s.Require().False(found)
}
