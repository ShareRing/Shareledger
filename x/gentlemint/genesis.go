package gentlemint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.ExchangeRate != nil {
		k.SetExchangeRate(ctx, *genState.ExchangeRate)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all exchangeRate
	exchangeRate, found := k.GetExchangeRate(ctx)
	if found {
		genesis.ExchangeRate = &exchangeRate
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
