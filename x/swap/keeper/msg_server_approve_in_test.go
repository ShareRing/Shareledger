package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestApproveIn() {
	acc, err := sdk.AccAddressFromBech32(testAddr)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), types.ModuleName, acc, sdk.Coins{tenSHR}).Return(nil)
	msgRequest := types.MsgApproveIn{
		Creator: "",
		Ids:     []uint64{0},
	}
	_, err = s.swapKeeper.AppendPendingRequest(s.ctx, types.Request{
		Id:          0,
		DestAddr:    testAddr,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      "pending",
		BatchId:     0,
		TxEvents: []*types.TxEvent{
			{
				TxHash: "0xXXX",
			},
		},
	})
	s.Require().NoError(err)
	_, err = s.msgServer.ApproveIn(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().NoError(err)

	// approve in duplicated request
	msgRequest = types.MsgApproveIn{
		Creator: "",
		Ids:     []uint64{1},
	}
	_, err = s.swapKeeper.AppendPendingRequest(s.ctx, types.Request{
		Id:          1,
		DestAddr:    testAddr,
		DestNetwork: types.NetworkNameShareLedger,
		Amount:      tenSHR,
		Fee:         tenSHR,
		Status:      "pending",
		BatchId:     0,
		TxEvents: []*types.TxEvent{
			{
				TxHash: "0xXXX",
			},
		},
	})
	s.Require().NoError(err)
	_, err = s.msgServer.ApproveIn(sdk.WrapSDKContext(s.ctx), &msgRequest)
	s.Require().NotNil(err)
}
