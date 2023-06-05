package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestBurn() {
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)
	baseCoins, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(tenDecCoin), false)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, baseCoins).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, baseCoins).Return(nil)
	for _, tc := range []struct {
		desc        string
		request     *types.MsgBurn
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgBurn{
				Creator: testAcc,
				Coins:   sdk.NewDecCoins(tenDecCoin),
			},
			expectedErr: false,
		},
		{
			desc: "InvalidRequest",
			request: &types.MsgBurn{
				Creator: testAcc,
				Coins: []sdk.DecCoin{{
					Denom:  "",
					Amount: sdk.NewDec(-10),
				}},
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)
			_, err := s.msgServer.Burn(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
