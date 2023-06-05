package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestSwap() {
	pendingRequests := s.createNRequest(normalRequest, 3)
	approverequests := s.createNRequest(normalRequest, 3)
	s.swapKeeper.MoveRequest(s.ctx, types.SwapStatusPending, types.SwapStatusApproved, approverequests, nil, true)
	req := &types.QuerySwapRequest{
		Status:      types.SwapStatusPending,
		Ids:         []uint64{},
		SrcAddr:     "",
		DestAddr:    "",
		SrcNetwork:  "",
		DestNetwork: "",
		Pagination:  &query.PageRequest{},
	}

	resp, err := s.swapKeeper.Swap(s.ctx, nil)
	s.Require().NotNil(err)

	resp, err = s.swapKeeper.Swap(s.ctx, req)
	s.Require().NoError(err)
	s.Require().Equal(3, len(resp.Swaps))

	for index, request := range resp.Swaps {
		s.Require().Equal(request, pendingRequests[index])
	}

	req.Status = "approved"
	resp, err = s.swapKeeper.Swap(s.ctx, req)
	s.Require().NoError(err)
	s.Require().Equal(3, len(resp.Swaps))
	for index, request := range resp.Swaps {
		s.Require().Equal(request, approverequests[index])
	}

	req.Status = "test"
	resp, err = s.swapKeeper.Swap(s.ctx, req)
	s.Require().Error(err)
}
