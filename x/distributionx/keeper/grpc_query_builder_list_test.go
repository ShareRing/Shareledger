package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/distributionx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperTestSuite) TestBuilderListQuerySingle() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNBuilderList(2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBuilderListRequest
		response *types.QueryGetBuilderListResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBuilderListRequest{Id: msgs[0].Id},
			response: &types.QueryGetBuilderListResponse{BuilderList: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetBuilderListRequest{Id: msgs[1].Id},
			response: &types.QueryGetBuilderListResponse{BuilderList: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetBuilderListRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.dKeeper.BuilderList(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.response, response)
			}
		})
	}
}

func (s *KeeperTestSuite) TestBuilderListQueryPaginated() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNBuilderList(5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBuilderListRequest {
		return &types.QueryAllBuilderListRequest{
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
			resp, err := s.dKeeper.BuilderListAll(wctx, request(nil, uint64(i), uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.BuilderList), step)
			s.Require().Subset(msgs, resp.BuilderList)
		}
	})
	s.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := s.dKeeper.BuilderListAll(wctx, request(next, 0, uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.BuilderList), step)
			s.Require().Subset(msgs, resp.BuilderList)
			next = resp.Pagination.NextKey
		}
	})
	s.Run("Total", func() {
		resp, err := s.dKeeper.BuilderListAll(wctx, request(nil, 0, 0, true))
		s.Require().NoError(err)
		s.Require().Equal(len(msgs), int(resp.Pagination.Total))
		s.Require().ElementsMatch(msgs, resp.BuilderList)
	})
	s.Run("InvalidRequest", func() {
		_, err := s.dKeeper.BuilderListAll(wctx, nil)
		s.Require().ErrorIs(err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
