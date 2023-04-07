package distributionx_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"testing"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	params := types.DefaultParams()
	params.BuilderWindows = 12
	params.TxThreshold = 2222
	params.DevPoolAccount = "share1"

	params.NativeDevelopment = sdk.NewDec(34)
	params.NativeValidator = sdk.NewDec(14)
	params.WasmValidator = sdk.NewDec(43)
	params.WasmMasterBuilder = sdk.NewDec(55)
	params.WasmContractAdmin = sdk.NewDec(55)
	params.WasmDevelopment = sdk.NewDec(33)

	genesisState := types.GenesisState{
		Params: params,
		RewardList: []types.Reward{
			{
				Index:  "0",
				Amount: sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(12333))),
			},
			{
				Amount: sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(3000))),
			},
		},
		BuilderCountList: []types.BuilderCount{
			{
				Index: "0",
				Count: 3,
			},
			{
				Index: "1",
				Count: 44,
			},
		},
		BuilderListList: []types.BuilderList{
			{
				Id:              0,
				ContractAddress: "shareledgerxxxxx",
			},
			{
				Id:              1,
				ContractAddress: "shareledgerxxx_333",
			},
		},
		BuilderListCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DistributionxKeeper(t)
	distributionx.InitGenesis(ctx, *k, genesisState)
	got := distributionx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RewardList, got.RewardList)
	require.ElementsMatch(t, genesisState.BuilderCountList, got.BuilderCountList)
	require.ElementsMatch(t, genesisState.BuilderListList, got.BuilderListList)
	require.Equal(t, genesisState.BuilderListCount, got.BuilderListCount)
	require.Equal(t, genesisState.Params.NativeValidator.String(), got.Params.NativeValidator.String())
	require.Equal(t, genesisState.Params.NativeDevelopment.String(), got.Params.NativeDevelopment.String())
	require.Equal(t, genesisState.Params.WasmValidator.String(), got.Params.WasmValidator.String())
	require.Equal(t, genesisState.Params.WasmMasterBuilder.String(), got.Params.WasmMasterBuilder.String())
	require.Equal(t, genesisState.Params.WasmContractAdmin.String(), got.Params.WasmContractAdmin.String())
	require.Equal(t, genesisState.Params.WasmDevelopment.String(), got.Params.WasmDevelopment.String())

	// this line is used by starport scaffolding # genesis/test/assert
}
