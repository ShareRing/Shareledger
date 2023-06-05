package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestBuyShr() {
	amount := "10"
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)

	shrCoin := sdk.NewDecCoinFromDec(denom.Shr, sdk.MustNewDecFromStr(amount))
	rate := s.gKeeper.GetExchangeRateD(s.ctx)
	coin, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(shrCoin), rate, false)
	cost, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoinsFromCoins(coin), rate, true)

	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, sdk.Coins{cost}).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, sdk.Coins{cost}).Return(nil)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{tenSHR, oneThousandCent})
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, sdk.Coins{coin}).Return(nil)
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenSHR)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, sdk.Coins{coin}).Return(nil)
	for _, tc := range []struct {
		desc        string
		request     *types.MsgBuyShr
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgBuyShr{
				Creator: testAcc,
				Amount:  amount,
			},
			expectedErr: false,
		},
		{
			desc: "InvalidRequest",
			request: &types.MsgBuyShr{
				Creator: "",
				Amount:  "10shr",
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)
			_, err := s.msgServer.BuyShr(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
