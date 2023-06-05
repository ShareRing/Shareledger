package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperTestSuite) TestLevelFeeQuerySingle() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNLevelFee(2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryLevelFeeRequest
		response *types.QueryLevelFeeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryLevelFeeRequest{
				Level: msgs[0].Level,
			},
			response: &types.QueryLevelFeeResponse{
				LevelFee: types.LevelFeeDetail{
					Level:        msgs[0].Level,
					Creator:      "",
					OriginalFee:  tenDecCoin,
					ConvertedFee: &tenMilNSHR,
				},
			}, //LevelFee: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryLevelFeeRequest{
				Level: msgs[1].Level,
			},
			response: &types.QueryLevelFeeResponse{
				LevelFee: types.LevelFeeDetail{
					Level:        msgs[1].Level,
					Creator:      "",
					OriginalFee:  tenDecCoin,
					ConvertedFee: &tenMilNSHR,
				},
			}, //LevelFee: msgs[1]
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.gKeeper.LevelFee(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().Equal(tc.response, response)
			}
		})
	}
}

func (s *KeeperTestSuite) TestLevelFeesQuery() {
	wctx := sdk.WrapSDKContext(s.ctx)
	msgs := s.createNLevelFee(2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryLevelFeesRequest
		response *types.QueryLevelFeesResponse
		err      error
	}{
		{
			desc:    "First",
			request: &types.QueryLevelFeesRequest{},
			response: &types.QueryLevelFeesResponse{
				LevelFees: []types.LevelFeeDetail{
					{
						Level:        msgs[0].Level,
						Creator:      "",
						OriginalFee:  tenDecCoin,
						ConvertedFee: &tenMilNSHR,
					},
					{
						Level:        msgs[1].Level,
						Creator:      "",
						OriginalFee:  tenDecCoin,
						ConvertedFee: &tenMilNSHR,
					},
				},
			},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		s.Run(tc.desc, func() {
			response, err := s.gKeeper.LevelFees(wctx, tc.request)
			if tc.err != nil {
				s.Require().ErrorIs(err, tc.err)
			} else {
				s.Require().Equal(tc.response.LevelFees, response.LevelFees[:2])
			}
		})
	}
}
