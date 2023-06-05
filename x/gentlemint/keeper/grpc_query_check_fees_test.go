package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func (s *KeeperTestSuite) TestCheckFees() {
	req := types.QueryCheckFeesRequest{
		Address: "",
		Actions: []string{""},
	}

	// nil request
	_, err := s.gKeeper.CheckFees(s.ctx, nil)
	s.Require().NotNil(err)

	// empty address in request
	_, err = s.gKeeper.CheckFees(s.ctx, &req)
	s.Require().NotNil(err)

	// test sufficient fee
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{tenSHR, tenMilNSHR})
	req.Address = testAcc
	req.Actions = []string{"gentlemint_load"}
	resp, err := s.gKeeper.CheckFees(s.ctx, &req)
	s.Require().NoError(err)
	s.Require().True(resp.SufficientFee)

	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), acc).Return(sdk.Coins{zeroNSHR})
	resp, err = s.gKeeper.CheckFees(s.ctx, &req)
	s.Require().NoError(err)
	s.Require().False(resp.SufficientFee)
}
