package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

func (s *KeeperTestSuite) TestRequestIn() {
	ctx := sdk.WrapSDKContext(s.ctx)
	msgRequest := types.MsgRequestIn{
		Creator:     "",
		SrcAddress:  "",
		DestAddress: "",
		Network:     "eth",
		Amount:      &sdk.DecCoin{Denom: "shr", Amount: sdk.NewDec(5)},
		TxEvents:    []*types.TxEvent{},
	}
	schema := types.Schema{
		Network:          "eth",
		Creator:          "",
		Schema:           "",
		ContractExponent: 0,
		Fee: &types.Fee{
			In:  &tenSHR,
			Out: &tenSHR,
		},
	}
	s.swapKeeper.SetSchema(s.ctx, schema)
	id, err := s.msgServer.RequestIn(ctx, &msgRequest)
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), id.Id)
	msgRequest1 := types.MsgRequestIn{
		Creator:     "",
		SrcAddress:  "",
		DestAddress: "",
		Network:     "bsc",
		Amount:      &sdk.DecCoin{Denom: "shr", Amount: sdk.NewDec(5)},
		TxEvents:    []*types.TxEvent{},
	}
	_, err = s.msgServer.RequestIn(ctx, &msgRequest1)
	s.Require().Error(err)
	msgRequest2 := types.MsgRequestIn{
		Creator:     "",
		SrcAddress:  "",
		DestAddress: "",
		Network:     "eth",
		Amount:      &sdk.DecCoin{Denom: "shr", Amount: sdk.NewDec(0)},
		TxEvents:    []*types.TxEvent{},
	}
	_, err = s.msgServer.RequestIn(s.ctx, &msgRequest2)
	s.Require().Error(err)
}
