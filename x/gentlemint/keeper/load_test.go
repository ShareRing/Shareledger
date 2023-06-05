package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestLoadCoins() {
	coins := sdk.NewCoins(tenSHR, tenMilNSHR)
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)

	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), types.ModuleName, coins).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, coins).Return(nil)

	err = s.gKeeper.LoadCoins(s.ctx, acc, coins)
	s.Require().NoError(err)
}
