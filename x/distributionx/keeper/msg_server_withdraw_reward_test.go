package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func (s *KeeperTestSuite) TestWithdrawReward() {
	wctx := sdk.WrapSDKContext(s.ctx)
	amount := sdk.NewCoin("nshr", sdk.NewInt(10000))
	destAcc, err := sdk.AccAddressFromBech32("shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf")
	s.Require().NoError(err)

	reward := types.Reward{
		Index:  "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
		Amount: []sdk.Coin{amount},
	}

	request := types.MsgWithdrawReward{
		Creator: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
	}

	s.dKeeper.SetReward(s.ctx, reward)

	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, destAcc, sdk.Coins{amount}).Return(nil)
	resp, err := s.msgServer.WithdrawReward(wctx, &request)
	s.Require().NoError(err)
	s.Require().Equal(resp.Amount, sdk.Coins{amount})
}
