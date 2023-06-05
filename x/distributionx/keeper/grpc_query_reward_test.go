package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/distributionx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperTestSuite) TestRewardQuerySingle() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNReward(2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRewardRequest
		response *types.QueryGetRewardResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRewardRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRewardResponse{Reward: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRewardRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRewardResponse{Reward: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRewardRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.dKeeper.Reward(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.response, response)
			}
		})
	}
}

func (s *KeeperTestSuite) TestRewardQueryPaginated() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNReward(5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRewardRequest {
		return &types.QueryAllRewardRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	s.Run("ByOffset", func() {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := s.dKeeper.RewardAll(wctx, request(nil, uint64(i), uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.Reward), step)
			s.Require().Subset(msgs, resp.Reward)
		}
	})
	s.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := s.dKeeper.RewardAll(wctx, request(next, 0, uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.Reward), step)
			s.Require().Subset(msgs, resp.Reward)
			next = resp.Pagination.NextKey
		}
	})
	s.Run("Total", func() {
		resp, err := s.dKeeper.RewardAll(wctx, request(nil, 0, 0, true))
		s.Require().NoError(err)
		s.Require().Equal(len(msgs), int(resp.Pagination.Total))
		s.Require().ElementsMatch(msgs, resp.Reward)
	})
	s.Run("InvalidRequest", func() {
		_, err := s.dKeeper.RewardAll(wctx, nil)
		s.Require().ErrorIs(err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
