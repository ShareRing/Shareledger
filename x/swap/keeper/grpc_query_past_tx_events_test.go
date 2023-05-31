package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) createPastTxEvent(request types.Request) {
	destAcc, err := sdk.AccAddressFromBech32(request.DestAddr)
	s.Require().NoError(err)

	s.swapKeeper.SetPastTxEvent(s.ctx, destAcc, "", request.TxEvents)
}

func (s *KeeperTestSuite) TestPastTxEvent() {
	msgRequest := types.QueryPastTxEventRequest{
		TxHash:   "0xXXX",
		LogIndex: 0,
	}
	req := s.createNRequest(request1, 1)
	s.createPastTxEvent(request1)
	resp, err := s.swapKeeper.PastTxEvent(s.ctx, nil)
	s.Require().Error(err)

	resp, err = s.swapKeeper.PastTxEvent(s.ctx, &msgRequest)
	s.Require().NoError(err)
	s.Require().Equal(req[0].DestAddr, resp.Event.DestAddr)
}

func (s *KeeperTestSuite) TestPastTxEventsByTxHash() {
	msgRequest := types.QueryPastTxEventsByTxHashRequest{
		TxHash: "0xXXX",
	}
	req := s.createNRequest(request1, 1)

	s.createPastTxEvent(request1)
	resp, err := s.swapKeeper.PastTxEventsByTxHash(s.ctx, nil)
	s.Require().Error(err)

	resp, err = s.swapKeeper.PastTxEventsByTxHash(s.ctx, &msgRequest)
	s.Require().NoError(err)
	s.Require().Equal(req[0].DestAddr, resp.Events[0].DestAddr)
}

func (s *KeeperTestSuite) TestPastTxEvents() {
	_ = s.createNRequest(request1, 1)
	_ = s.createNRequest(request2, 1)

	s.createPastTxEvent(request1)
	s.createPastTxEvent(request2)

	reqMsg := types.QueryPastTxEventsRequest{
		Pagination: &query.PageRequest{
			Key:        []byte{},
			Offset:     0,
			Limit:      0,
			CountTotal: false,
			Reverse:    false,
		},
	}

	resp, err := s.swapKeeper.PastTxEvents(s.ctx, nil)
	s.Require().Error(err)

	resp, err = s.swapKeeper.PastTxEvents(s.ctx, &reqMsg)
	s.Require().NoError(err)
	s.Require().Equal(2, len(resp.Events))
}
