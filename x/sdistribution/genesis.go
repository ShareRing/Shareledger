package sdistribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/sdistribution/keeper"
	"github.com/sharering/shareledger/x/sdistribution/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the reward
	for _, elem := range genState.RewardList {
		k.SetReward(ctx, elem)
	}
	// Set all the builderCount
	for _, elem := range genState.BuilderCountList {
		k.SetBuilderCount(ctx, elem)
	}
	// Set all the builderList
	for _, elem := range genState.BuilderListList {
		k.SetBuilderList(ctx, elem)
	}

	// Set builderList count
	k.SetBuilderListCount(ctx, genState.BuilderListCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RewardList = k.GetAllReward(ctx)
	genesis.BuilderCountList = k.GetAllBuilderCount(ctx)
	genesis.BuilderListList = k.GetAllBuilderList(ctx)
	genesis.BuilderListCount = k.GetBuilderListCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}