package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestGetRequestCount() {
	items := s.createNRequest(normalRequest, 10)
	count := uint64(len(items))
	s.Require().Equal(count, s.swapKeeper.GetRequestCount(s.ctx))
}

func (s *KeeperTestSuite) TestSetRequestCount() {
	s.swapKeeper.SetRequestCount(s.ctx, 5)
	s.Require().Equal(uint64(5), s.swapKeeper.GetRequestCount(s.ctx))
}

func (s *KeeperTestSuite) TestAppendPendingRequest() {
	request := types.Request{
		Id:          0,
		SrcAddr:     "",
		DestAddr:    "",
		SrcNetwork:  "",
		DestNetwork: "",
		Amount:      sdk.Coin{},
		Fee:         sdk.Coin{},
		Status:      "pending",
		BatchId:     0,
		TxEvents:    []*types.TxEvent{},
	}
	_, err := s.swapKeeper.AppendPendingRequest(s.ctx, request)
	s.Require().NoError(err)

}

func (s *KeeperTestSuite) TestChangeStatusRequests() {
	var batchId uint64

	// Name:    "pass",
	req := s.createNRequest(normalRequest, 1)
	batchId = s.swapKeeper.AppendBatch(s.ctx, types.Batch{
		Signature:  "",
		RequestIds: []uint64{req[0].Id},
		Status:     types.BatchStatusPending,
	})
	status, err := s.swapKeeper.ChangeStatusRequests(s.ctx, []uint64{0}, types.SwapStatusApproved, &batchId, true)
	s.Require().NoError(err)
	s.Require().Equal(types.SwapStatusApproved, status[0].Status)

	// Name:    "ErrInvalidRequest",
	_, err = s.swapKeeper.ChangeStatusRequests(s.ctx, nil, types.SwapStatusApproved, &batchId, true)
	s.Require().NotNil(err)

	// Name:    "NoBatchIdProvided",
	_, err = s.swapKeeper.ChangeStatusRequests(s.ctx, []uint64{0}, types.SwapStatusApproved, nil, true)
	s.Require().NotNil(err)
}

func (s *KeeperTestSuite) TestGetRequestFromStore() {
	req := s.createNRequest(normalRequest, 1)
	stores := s.swapKeeper.GetStoreRequestMap(s.ctx)
	store := stores[types.SwapStatusPending]

	// normal case
	resp, found := s.swapKeeper.GetRequestFromStore(s.ctx, store, req[0].Id)
	s.Require().True(found)
	s.Require().Equal(req[0], resp)

	// wrong request id
	_, found = s.swapKeeper.GetRequestFromStore(s.ctx, store, 10)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestGetAllRequest() {
	req := s.createNRequest(normalRequest, 1)
	req1 := s.createNRequest(request1, 1)
	// approve the request1
	err := s.swapKeeper.MoveRequest(s.ctx, types.SwapStatusPending, types.SwapStatusApproved, req1, nil, true)
	s.Require().NoError(err)

	stores := s.swapKeeper.GetStoreRequestMap(s.ctx)
	store := stores[types.SwapStatusApproved]
	resp, found := s.swapKeeper.GetRequestFromStore(s.ctx, store, req1[0].Id)
	s.Require().True(found)
	req = append(req, resp)

	requests := s.swapKeeper.GetAllRequest(s.ctx)
	s.Require().ElementsMatch(req, requests)

}
