package keeper_test

import (
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestQueryFundBalance() {
	s.accKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(swapModuleAddress)
	s.bankKeeper.EXPECT().GetBalance(gomock.Any(), swapModuleAddress, gomock.Any()).Return(tenSHR)
	resp, err := s.swapKeeper.Balance(s.ctx, &types.QueryBalanceRequest{})
	s.Require().NoError(err)
	s.Require().Equal(resp.Balance, &tenSHR)
}
