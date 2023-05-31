package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestSetPastTxEvent() {
	events := types.TxEvent{
		TxHash:   "0xXXX",
		Sender:   "",
		LogIndex: 0,
	}
	addr, err := sdk.AccAddressFromBech32(testAddr)
	s.Require().NoError(err)
	s.swapKeeper.SetPastTxEvent(s.ctx, addr, "", []*types.TxEvent{&events})
	resp, found := s.swapKeeper.GetPastTxEvent(s.ctx, "0xXXX", 0)
	s.Require().True(found)
	s.Require().Equal(testAddr, resp.DestAddr)
}

func (s *KeeperTestSuite) TestRemovePastTxEvent() {
	events := types.TxEvent{
		TxHash:   "0xXXX",
		Sender:   "",
		LogIndex: 0,
	}
	addr, err := sdk.AccAddressFromBech32(testAddr)
	s.Require().NoError(err)
	s.swapKeeper.SetPastTxEvent(s.ctx, addr, "", []*types.TxEvent{&events})
	s.swapKeeper.RemovePastTxEvent(s.ctx, "0xXXX", 0)
	_, found := s.swapKeeper.GetPastTxEvent(s.ctx, "0xXXX", 0)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestGetPastTxEventsByTxHash() {
	events := types.TxEvent{
		TxHash:   "0xXXX",
		Sender:   "",
		LogIndex: 0,
	}
	addr, err := sdk.AccAddressFromBech32(testAddr)
	s.Require().NoError(err)
	s.swapKeeper.SetPastTxEvent(s.ctx, addr, "", []*types.TxEvent{&events})
	event := s.swapKeeper.GetPastTxEventsByTxHash(s.ctx, "0xXXX")
	s.Require().Equal(testAddr, event[0].DestAddr)
}
