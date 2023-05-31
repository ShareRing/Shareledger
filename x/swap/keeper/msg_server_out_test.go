package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestRequestOut() {
	ctx := sdk.WrapSDKContext(s.ctx)
	acc, err := sdk.AccAddressFromBech32(testAddr)
	amount := sdk.NewDecCoin("nshr", sdk.NewInt(500))
	amount1 := sdk.NewDecCoin("shr", sdk.NewInt(0))
	expectedAmount, err := denom.NormalizeToBaseCoins(sdk.NewDecCoins(amount), true)
	sumCoin := expectedAmount.Add(tenSHR)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), acc, types.ModuleName, sumCoin).Return(nil)

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
	msgRequest := types.MsgRequestOut{
		Creator:     testAddr,
		SrcAddress:  "",
		DestAddress: "",
		Network:     "eth",
		Amount:      &amount,
	}
	s.swapKeeper.SetSchema(s.ctx, schema)

	// the creator account need to pay to type of denom (shr, nshr) when swap out, so need to be sure
	// the creator account these two types of denom balance.
	// - the nshr pay for swap out amount.
	// - the shr pay for swap fee.
	resp, err := s.msgServer.RequestOut(ctx, &msgRequest)
	s.Require().NoError(err)
	s.Require().Equal(uint64(0), resp.Id)
	request, found := s.swapKeeper.GetRequest(s.ctx, 0)
	s.Require().True(found)
	s.Require().Equal(types.SwapStatusPending, request.Status)

	msgRequest1 := types.MsgRequestOut{
		Creator:     testAddr,
		SrcAddress:  "",
		DestAddress: "",
		Network:     "bsc",
		Amount:      &amount,
	}

	_, err = s.msgServer.RequestOut(ctx, &msgRequest1)
	s.Require().Error(err)
	msgRequest2 := types.MsgRequestOut{
		Creator:     testAddr,
		SrcAddress:  "",
		DestAddress: "",
		Network:     "eth",
		Amount:      &amount1,
	}

	_, err = s.msgServer.RequestOut(ctx, &msgRequest2)
	s.Require().Error(err)
}
