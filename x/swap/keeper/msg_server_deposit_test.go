package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestDeposit() {
	acc, err := sdk.AccAddressFromBech32(testAddr)
	amount := sdk.NewDecCoin("shr", sdk.NewInt(10))
	amount1 := sdk.NewDecCoin("shr", sdk.NewInt(0))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, sdk.Coins{tenMilNSHR}).Return(nil)
	msgRequest := types.MsgDeposit{
		Creator: testAddr,
		Amount:  &amount,
	}
	_, err = s.msgServer.Deposit(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().NoError(err)
	msgRequest = types.MsgDeposit{
		Creator: testAddr,
		Amount:  &amount1,
	}
	_, err = s.msgServer.Deposit(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().Error(err)
}
