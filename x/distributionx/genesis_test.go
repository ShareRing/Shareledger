package distributionx_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/sharering/shareledger/testutil/keeper"
	"github.com/sharering/shareledger/testutil/nullify"
	"github.com/sharering/shareledger/x/distributionx"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func TestGenesis(t *testing.T) {
	params := types.DefaultParams()
	params.BuilderWindows = 12
	params.TxThreshold = 2222
	params.DevPoolAccount = "share1"

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

	// this line is used by starport scaffolding # genesis/test/assert
}
