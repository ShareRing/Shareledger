package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) createNRequest(request types.Request, n int) []types.Request {
	items := make([]types.Request, n)
	for i := 0; i < n; i++ {
		m, err := s.swapKeeper.AppendPendingRequest(s.ctx, request)
		if err != nil {
			panic(err)
		}
		items[i] = m
	}
	return items
}

func (s *KeeperTestSuite) TestNextRequestId() {
	_ = s.createNRequest(normalRequest, 3)

	resp, err := s.swapKeeper.NextRequestId(s.ctx, &types.QueryNextRequestIdRequest{})
	s.Require().NoError(err)
	s.Require().Equal(3, int(resp.NextCount))

	s.swapKeeper.RemoveBatch(s.ctx, 0)
	resp, err = s.swapKeeper.NextRequestId(sdk.WrapSDKContext(s.ctx), &types.QueryNextRequestIdRequest{})
	s.Require().NoError(err)
	s.Require().Equal(3, int(resp.NextCount))

}

func (s *KeeperTestSuite) TestNextRequestId_invalidRequest() {
	_ = s.createNRequest(normalRequest, 1)

	_, err := s.swapKeeper.NextRequestId(s.ctx, nil)
	s.Require().Error(err)
}
