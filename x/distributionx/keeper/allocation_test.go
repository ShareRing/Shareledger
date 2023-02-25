package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/distributionx/types"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

// empty master builder list -> all money will transfer to dev pool account
func (s *KeeperTestSuite) TestAllocateTokens_EmptyMasterBuilderList() {
	dev_pool := s.TestAccs[0].String()
	// send some money to feewasmpool
	coins := sdk.Coins{sdk.Coin{
		Denom:  denom.Base,
		Amount: math.NewInt(1000000000),
	}}
	s.Nil(s.App.BankKeeper.MintCoins(s.Ctx, gentleminttypes.ModuleName, coins))
	s.Nil(s.App.BankKeeper.SendCoinsFromModuleToModule(s.Ctx, gentleminttypes.ModuleName, types.FeeWasmName, coins))

	s.dKeeper.AllocateTokens(s.Ctx)
	reward, found := s.dKeeper.GetReward(s.Ctx, dev_pool)
	s.True(found)
	s.True(reward.Amount.IsEqual(coins))
}
