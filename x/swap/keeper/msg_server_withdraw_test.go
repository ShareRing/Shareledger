package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestWithdraw() {
	moduleAcc, err := sdk.AccAddressFromBech32(string(swapModuleAddress))
	s.Require().NoError(err)
	destAcc, err := sdk.AccAddressFromBech32(string(testAddr))
	s.Require().NoError(err)
	wctx := sdk.WrapSDKContext(s.ctx)

	for _, tc := range []struct {
		desc     string
		request  *types.MsgWithdraw
		response *types.MsgWithdrawResponse
		testFunc func(*types.MsgWithdraw, *types.MsgWithdrawResponse)
	}{
		{
			desc: "ModuleOutOfCoin",
			request: &types.MsgWithdraw{
				Creator:  "",
				Receiver: testAddr,
				Amount:   sdk.NewDecCoin("shr", sdk.NewInt(5)),
			},
			testFunc: func(request *types.MsgWithdraw, response *types.MsgWithdrawResponse) {
				baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(request.Amount), true)
				s.accKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(moduleAcc)
				s.bankKeeper.EXPECT().SpendableCoins(gomock.Any(), moduleAcc).Return(sdk.Coins{tenSHR, tenMilNSHR})
				s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, destAcc, baseCoins).Return(fmt.Errorf("Not nil"))

				status, err := s.msgServer.Withdraw(wctx, request)
				s.Require().Error(err)
				s.Require().Equal(response, status)
			},
			response: &types.MsgWithdrawResponse{
				Status: types.TxnStatusFail,
			},
		},
		{
			desc: "ok",
			request: &types.MsgWithdraw{
				Creator:  "",
				Receiver: testAddr,
				Amount:   sdk.NewDecCoin("shr", sdk.NewInt(5)),
			},
			testFunc: func(request *types.MsgWithdraw, response *types.MsgWithdrawResponse) {
				baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(request.Amount), true)
				s.accKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(moduleAcc)
				s.bankKeeper.EXPECT().SpendableCoins(gomock.Any(), moduleAcc).Return(sdk.Coins{tenSHR, tenMilNSHR})
				s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, destAcc, baseCoins).Return(nil)
				status, err := s.msgServer.Withdraw(wctx, request)
				s.Require().NoError(err)
				s.Require().Equal(response, status)

			},
			response: &types.MsgWithdrawResponse{
				Status: types.TxnStatusSuccess,
			},
		},
		{
			desc: "ErrInvalidCoins",
			request: &types.MsgWithdraw{
				Creator:  "",
				Receiver: testAddr,
				Amount:   sdk.NewDecCoin("shr", sdk.NewInt(0)),
			},
			testFunc: func(request *types.MsgWithdraw, response *types.MsgWithdrawResponse) {
				status, err := s.msgServer.Withdraw(wctx, request)
				s.Require().Error(err)
				s.Require().Equal(response, status)

			},
			response: &types.MsgWithdrawResponse{
				Status: types.TxnStatusFail,
			},
		},
		{
			desc: "ErrInvalidAddress",
			request: &types.MsgWithdraw{
				Creator:  "",
				Receiver: "test1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
				Amount:   sdk.NewDecCoin("shr", sdk.NewInt(5)),
			},
			testFunc: func(request *types.MsgWithdraw, response *types.MsgWithdrawResponse) {
				status, err := s.msgServer.Withdraw(wctx, request)
				status, err = s.msgServer.Withdraw(wctx, request)
				s.Require().Error(err)
				s.Require().Equal(response, status)

			},
			response: &types.MsgWithdrawResponse{
				Status: types.TxnStatusFail,
			},
		},
	} {
		s.Run(tc.desc, func() {
			tc.testFunc(tc.request, tc.response)
		})
	}
}
