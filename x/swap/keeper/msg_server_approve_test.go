package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestApprove() {
	reqs := s.createNRequest(normalRequest, 1)
	batchNum := s.swapKeeper.GetBatchCount(s.ctx)
	s.Require().Zero(batchNum)

	for _, tc := range []struct {
		desc        string
		request     *types.MsgApproveOut
		response    *types.MsgApproveOutResponse
		expectedErr bool
	}{
		{
			desc: "ok",
			request: &types.MsgApproveOut{
				Creator:   "",
				Signature: "",
				Ids:       []uint64{0},
			},
			response:    &types.MsgApproveOutResponse{},
			expectedErr: false,
		},
		{
			desc: "EmptyRequestId",
			request: &types.MsgApproveOut{
				Creator:   "",
				Signature: "",
				Ids:       []uint64{},
			},
			response:    &types.MsgApproveOutResponse{},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			resp, err := s.msgServer.ApproveOut(sdk.WrapSDKContext(s.ctx), tc.request)
			if tc.expectedErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				_, found := s.swapKeeper.GetBatch(s.ctx, resp.BatchId)
				s.Require().True(found)
				request, found := s.swapKeeper.GetRequest(s.ctx, reqs[0].GetId())
				s.Require().Equal("approved", request.Status)
				s.Require().Equal(resp.BatchId, request.BatchId)
			}
		})
	}
}
