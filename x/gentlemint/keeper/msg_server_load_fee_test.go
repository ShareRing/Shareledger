package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestLoadFee() {
	wctx := sdk.WrapSDKContext(s.ctx)
	coin := sdk.NewDecCoin("shrp", sdk.NewInt(2))

	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)

	rate := s.gKeeper.GetExchangeRateD(s.ctx)
	boughtBase, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(coin), rate, false)
	cost, err := denom.NormalizeToBaseCoin(denom.BaseUSD, sdk.NewDecCoinsFromCoins(boughtBase), rate, true)

	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{tenMilNSHR, oneThousandCent})
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, sdk.Coins{boughtBase}).Return(nil)
	s.bankKeeper.EXPECT().GetSupply(gomock.Any(), gomock.Any()).Return(tenSHR)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, sdk.Coins{cost}).Return(nil)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), types.ModuleName, gomock.Any()).Return(nil)
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	req := types.MsgLoadFee{
		Creator: testAcc,
		Shrp:    &coin,
	}

	_, err = s.msgServer.LoadFee(wctx, &req)
	s.Require().NoError(err)
}
