package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/constant"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestGetBaseFeeByMsg() {
	fee := constant.DefaultFeeLevel["min"]
	usdRate := s.gKeeper.GetExchangeRateD(s.ctx)
	expected, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(fee), usdRate, true)
	s.Require().NoError(err)

	s.createActionLevelFee(testActionFee)
	resp, err := s.gKeeper.GetBaseFeeByMsg(s.ctx, msgTest)
	s.Require().NoError(err)
	s.Require().NotEmpty(resp, expected)
}

func (s *KeeperTestSuite) TestGetBaseDenomFeeByActionKey() {
	s.createActionLevelFee(testActionFee)
	resp, err := s.gKeeper.GetBaseDenomFeeByActionKey(s.ctx, testActionFee.Action)
	s.Require().NoError(err)
	s.Require().NotEmpty(resp)
}

func (s *KeeperTestSuite) TestLoadFeeFundFromShrp() {
	msg := types.MsgLoadFee{
		Creator: testAcc,
		Shrp:    nil,
	}

	rate := s.gKeeper.GetExchangeRateD(s.ctx)
	coin, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(tenDecCoin), rate, false)
	expectedCoin, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoinsFromCoins(coin), rate, true)

	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), denom.Base).Return(tenSHR)
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{tenSHR, oneThousandCent})
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, sdk.Coins{expectedCoin}).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, sdk.Coins{expectedCoin}).Return(nil)
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, sdk.Coins{coin}).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, sdk.Coins{coin}).Return(nil)

	// test shrp is nil
	err = s.gKeeper.LoadFeeFundFromShrp(s.ctx, &msg)
	s.Require().NotNil(err)

	// test normal case
	msg.Shrp = &tenDecCoin
	err = s.gKeeper.LoadFeeFundFromShrp(s.ctx, &msg)
	s.Require().NoError(err)
}
