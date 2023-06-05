package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

// test load 500.000nshr to dest account, but the fee need to be paid is 2shr.
// so need to load more coin to paid for the dest address to paid for creator
func (s *KeeperTestSuite) TestLoad() {
	coinsWantToLoad := sdk.NewDecCoins(fiveHundredThousandNSHRDecCoin)
	currentDestAccBalance := sdk.NewCoin("nshr", sdk.NewInt(1000000000))
	baseCoins, err := denom.NormalizeToBaseCoins(coinsWantToLoad, false)
	exchangeRate := s.gKeeper.GetExchangeRateD(s.ctx)
	msgTypeFee := s.gKeeper.GetFeeByMsg(s.ctx, &types.MsgLoad{})
	realFee, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(msgTypeFee), exchangeRate, true)
	expectedAmountToLoadMore := realFee.Sub(currentDestAccBalance)
	expectedCent, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoinsFromCoins(expectedAmountToLoadMore), exchangeRate, true)

	acc1, err := sdk.AccAddressFromBech32(testAcc1)
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)
	s.bankKeeper.EXPECT().GetBalance(gomock.Any(), acc1, gomock.Any()).Return(currentDestAccBalance)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc1).Return(sdk.Coins{currentDestAccBalance, oneThousandCent})
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, baseCoins).Return(nil)
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenMilNSHR)
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenMilNSHR)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc1, baseCoins).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc1, sdk.NewCoins(expectedAmountToLoadMore)).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc1, types.ModuleName, sdk.Coins{expectedCent}).Return(nil)
	s.bankKeeper.EXPECT().SendCoins(gomock.Any(), acc1, acc, sdk.NewCoins(realFee)).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, sdk.Coins{expectedCent}).Return(nil)
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, sdk.Coins{expectedAmountToLoadMore}).Return(nil)

	for _, tc := range []struct {
		desc        string
		request     *types.MsgLoad
		expectedErr bool
	}{
		{
			desc: "Completed",
			request: &types.MsgLoad{
				Creator: testAcc,
				Address: testAcc1,
				Coins:   coinsWantToLoad,
			},
			expectedErr: false,
		},
		{
			desc: "InvalidRequest",
			request: &types.MsgLoad{
				Creator: "",
				Address: "",
				Coins:   []sdk.DecCoin{},
			},
			expectedErr: true,
		},
	} {
		s.Run(tc.desc, func() {
			wctx := sdk.WrapSDKContext(s.ctx)

			_, err := s.msgServer.Load(wctx, tc.request)
			if tc.expectedErr {
				s.Require().NotNil(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}
