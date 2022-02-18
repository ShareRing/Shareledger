package document

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/document/keeper"
	"github.com/sharering/shareledger/x/document/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for _, doc := range genState.Documents {
		k.SetDoc(ctx, doc)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	cb := func(doc types.Document) bool {
		genesis.Documents = append(genesis.Documents, &doc)
		return false
	}

	k.IterateDocs(ctx, cb)

	return genesis
}
