package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestBalances() {
	req := types.QueryBalancesRequest{
		Address: "",
	}

	// nil request
	_, err := s.gKeeper.Balances(s.ctx, nil)
	s.Require().Error(err)

	// empty address in request
	_, err = s.gKeeper.Balances(s.ctx, &req)
	s.Require().Error(err)

	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{tenSHR})

	req.Address = testAcc
	resp, err := s.gKeeper.Balances(s.ctx, &req)
	s.Require().NoError(err)
	s.Require().NotEmpty(resp, tenDecCoin)
}
