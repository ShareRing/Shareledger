package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestUpdateSwapFee() {
	wctx := sdk.WrapSDKContext(s.ctx)
	fiveSHR := sdk.NewDecCoin("shr", sdk.NewInt(5))
	fiveMilNSHR, err := denom.NormalizeToBaseCoin(denom.Base, sdk.NewDecCoins(fiveSHR), sdk.NewDec(0), false)
	s.Require().NoError(err)
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
	msgRequest := types.MsgUpdateSwapFee{
		Creator: "",
		Network: "eth",
		In:      &fiveSHR,
		Out:     &fiveSHR,
	}

	s.swapKeeper.SetSchema(s.ctx, schema)
	_, err = s.msgServer.UpdateSwapFee(wctx, &msgRequest)
	s.Require().NoError(err)
	schema, found := s.swapKeeper.GetSchema(s.ctx, "eth")
	s.Require().True(found)
	s.Require().Equal(fiveMilNSHR, *schema.Fee.In)
	s.Require().Equal(fiveMilNSHR, *schema.Fee.Out)

	msgRequest1 := types.MsgUpdateSwapFee{
		Creator: "",
		Network: "bsc",
		In:      &fiveSHR,
		Out:     &fiveSHR,
	}
	_, err = s.msgServer.UpdateSwapFee(wctx, &msgRequest1)
	s.Require().NotNil(err)

	msgRequest2 := types.MsgUpdateSwapFee{
		Creator: "",
		Network: "eth",
		In:      &sdk.DecCoin{Denom: "test", Amount: sdk.NewDec(5)},
		Out:     &fiveSHR,
	}
	_, err = s.msgServer.UpdateSwapFee(wctx, &msgRequest2)
	s.Require().NotNil(err)

}
