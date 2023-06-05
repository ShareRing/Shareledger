package keeper_test

import (
	"cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/mock/gomock"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func (s *KeeperTestSuite) TestAllocateTokens() {
	dev_pool := "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf"
	testFeeWasmModuleAcc := authtypes.NewEmptyModuleAccount(types.FeeWasmName)
	testFeeNativeAcc := authtypes.NewEmptyModuleAccount(types.FeeNativeName)
	amount := sdk.Coin{
		Denom:  denom.Base,
		Amount: math.NewInt(1000000000),
	}

	contractInfo := wasmtypes.ContractInfo{
		CodeID:  0,
		Creator: "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
	}

	// create builder list with contract address
	builder := types.BuilderList{
		Id:              0,
		ContractAddress: dev_pool,
	}

	contractAddress, err := sdk.AccAddressFromBech32(builder.ContractAddress)
	s.Require().NoError(err)
	s.accKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.FeeWasmName).Return(testFeeWasmModuleAcc)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), testFeeWasmModuleAcc.GetAddress()).Return(sdk.NewCoins(amount))
	s.accKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.FeeNativeName).Return(testFeeNativeAcc)
	s.bankKeeper.EXPECT().GetAllBalances(gomock.Any(), testFeeNativeAcc.GetAddress()).Return(sdk.NewCoins(amount))
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), types.FeeWasmName, types.ModuleName, sdk.Coins{amount}).Return(nil)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToModule(gomock.Any(), types.FeeNativeName, types.ModuleName, sdk.Coins{amount}).Return(nil)
	s.wasmKeeper.EXPECT().GetContractInfo(gomock.Any(), contractAddress).Return(&contractInfo)

	_ = s.dKeeper.AppendBuilderList(s.ctx, builder)

	s.dKeeper.AllocateTokens(s.ctx)
}
