package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/gentlemint/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestSend() {
	req := types.MsgSend{
		Creator: "",
		Address: "",
		Coins:   sdk.NewDecCoins(tenDecCoin),
	}

	// invalid request
	_, err := s.msgServer.Send(s.ctx, &req)
	s.Require().Error(err)

	// normal request
	acc, err := sdk.AccAddressFromBech32(testAcc)
	s.Require().NoError(err)
	acc1, err := sdk.AccAddressFromBech32(testAcc1)
	s.Require().NoError(err)
	coin, err := denom.NormalizeToBaseCoins(req.Coins, false)
	s.bankKeeper.EXPECT().SendCoins(gomock.Any(), acc, acc1, coin).Return(nil)

	req.Creator = testAcc
	req.Address = testAcc1
	_, err = s.msgServer.Send(s.ctx, &req)
	s.Require().NoError(err)
}
