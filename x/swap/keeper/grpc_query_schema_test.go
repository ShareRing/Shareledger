package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) createNSchema(n int) []types.Schema {
	items := make([]types.Schema, n)
	for i := range items {
		items[i].Network = strconv.Itoa(i)

		s.swapKeeper.SetSchema(s.ctx, items[i])
	}
	return items
}

func (s *KeeperTestSuite) TestFormatQuerySingle() {
	msgs := s.createNSchema(2)
	for _, tc := range []struct {
		desc     string
		request  *types.QuerySchemaRequest
		response *types.QuerySchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QuerySchemaRequest{
				Network: msgs[0].Network,
			},
			response: &types.QuerySchemaResponse{Schema: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QuerySchemaRequest{
				Network: msgs[1].Network,
			},
			response: &types.QuerySchemaResponse{Schema: msgs[1]},
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.swapKeeper.Schema(sdk.WrapSDKContext(s.ctx), tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(tc.err, err)
				return
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.response, response)
			}
		})
	}
}

func (s *KeeperTestSuite) TestFormatQueryPaginated() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNSchema(5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QuerySchemasRequest {
		return &types.QuerySchemasRequest{
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
			resp, err := s.swapKeeper.Schemas(wctx, request(nil, uint64(i), uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.Schemas), step)
			s.Require().Subset(msgs, resp.Schemas)
		}
	})
	s.Run("ByKey", func() {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := s.swapKeeper.Schemas(wctx, request(next, 0, uint64(step), false))
			s.Require().NoError(err)
			s.Require().LessOrEqual(len(resp.Schemas), step)
			s.Require().Subset(msgs, resp.Schemas)
			next = resp.Pagination.NextKey
		}
	})
	s.Run("Total", func() {
		resp, err := s.swapKeeper.Schemas(wctx, request(nil, 0, 0, true))
		s.Require().NoError(err)
		s.Require().Equal(len(msgs), int(resp.Pagination.Total))
		s.Require().ElementsMatch(msgs, resp.Schemas)
	})
}
