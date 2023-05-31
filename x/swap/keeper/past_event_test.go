package keeper_test

import (
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestAllPastEventGenesis() {
	req1 := s.createNRequest(request1, 1)
	req2 := s.createNRequest(request2, 1)

	s.createPastTxEvent(request1)
	s.createPastTxEvent(request2)
	resp := s.swapKeeper.AllPastEventGenesis(s.ctx)
	s.Require().Equal(2, len(resp))
	s.Require().Equal(req1[0].DestAddr, resp[0].DestAddr)
	s.Require().Equal(req2[0].DestAddr, resp[1].DestAddr)
}

func (s *KeeperTestSuite) TestSetPastEventFromGenesis() {
	reqs := []types.PastTxEventGenesis{
		{
			SrcAddr:  "",
			DestAddr: "",
			TxHash:   "0xXXX",
			LogIndex: 0,
		},
		{
			SrcAddr:  "",
			DestAddr: "",
			TxHash:   "0xABC",
			LogIndex: 1,
		},
		{
			SrcAddr:  "",
			DestAddr: "",
			TxHash:   "0xXYZ",
			LogIndex: 0,
		},
	}

	s.swapKeeper.SetPastEventFromGenesis(s.ctx, reqs)
	resp := s.swapKeeper.AllPastEventGenesis(s.ctx)
	s.Require().Equal(3, len(resp))
}
