package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestSetLevelFee() {
	wctx := sdk.WrapSDKContext(s.ctx)
	request := &types.MsgSetLevelFee{
		Creator: testAcc,
		Level:   "test",
	}
	_, err := s.msgServer.SetLevelFee(wctx, request)
	s.Require().NotNil(err)

	request.Fee = sdk.NewDecCoin("shrp", sdk.NewInt(10))
	_, err = s.msgServer.SetLevelFee(wctx, request)
	s.Require().NoError(err)
	resp, found := s.gKeeper.GetLevelFee(s.ctx,
		request.Level,
	)
	s.Require().True(found)
	s.Require().Equal(request.Level, resp.Level)
}

func (s *KeeperTestSuite) TestDeleteLevelFee() {
	for _, tc := range []struct {
		desc        string
		request     *types.MsgDeleteLevelFee
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteLevelFee{
				Creator: testAcc,
				Level:   strconv.Itoa(0),
			},
			expectedErr: false,
		},
		{
			desc: "InvalidRequest",
			request: &types.MsgDeleteLevelFee{
				Creator: testAcc,
				Level:   strconv.Itoa(100000),
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)

			_, err := s.msgServer.SetLevelFee(wctx, &types.MsgSetLevelFee{
				Creator: testAcc,
				Level:   strconv.Itoa(0),
				Fee:     sdk.NewDecCoin("shrp", sdk.NewInt(10)),
			})
			s.Require().NoError(err)
			_, err = s.msgServer.DeleteLevelFee(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)
				_, found := s.gKeeper.GetLevelFee(s.ctx,
					tc.request.Level,
				)
				s.Require().False(found)
			}
		})
	}
}
