package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperTestSuite) TestExchangeRateQuery() {
	wctx := sdk.WrapSDKContext(s.ctx)
	item := s.createTestExchangeRate()
	for _, tc := range []struct {
		desc     string
		request  *types.QueryExchangeRateRequest
		response *types.QueryExchangeRateResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryExchangeRateRequest{},
			response: &types.QueryExchangeRateResponse{Rate: item.Rate},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.gKeeper.ExchangeRate(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().Equal(sdk.MustNewDecFromStr(tc.response.Rate), sdk.MustNewDecFromStr(response.Rate))
			}
		})
	}
}
