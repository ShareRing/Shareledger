package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestSetActionLevelFee() {

	for _, tc := range []struct {
		desc        string
		request     *types.MsgSetActionLevelFee
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgSetActionLevelFee{
				Creator: testAcc,
				Action:  "gentlemint_load",
				Level:   "min",
			},
			expectedErr: false,
		},
		{
			desc: "LevelNotFound",
			request: &types.MsgSetActionLevelFee{
				Creator: testAcc,
				Action:  "gentlemint_load",
				Level:   "test",
			},
			expectedErr: true,
		},
		{
			desc: "ActionNotFound",
			request: &types.MsgSetActionLevelFee{
				Creator: testAcc,
				Action:  strconv.Itoa(100000),
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)
			_, err := s.msgServer.SetActionLevelFee(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)

				_, found := s.gKeeper.GetActionLevelFee(s.ctx,
					tc.request.Action,
				)
				s.Require().True(found)
			}
		})
	}
}

func (s *KeeperTestSuite) TestDeleteActionLevelFee() {
	for _, tc := range []struct {
		desc        string
		request     *types.MsgDeleteActionLevelFee
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActionLevelFee{
				Creator: testAcc,
				Action:  "gentlemint_load",
			},
			expectedErr: false,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActionLevelFee{
				Creator: testAcc,
				Action:  strconv.Itoa(100000),
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)

			// update the fee table element to new value
			_, err := s.msgServer.SetActionLevelFee(wctx, &types.MsgSetActionLevelFee{
				Creator: testAcc,
				Action:  "gentlemint_load",
				Level:   "low",
			})
			s.Require().NoError(err)

			_, err = s.msgServer.DeleteActionLevelFee(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)
				_, found := s.gKeeper.GetActionLevelFee(s.ctx,
					tc.request.Action,
				)
				s.Require().False(found)
			}
		})
	}
}
