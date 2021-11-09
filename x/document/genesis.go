package document

import (
	"github.com/ShareRing/Shareledger/x/document/keeper"
	"github.com/ShareRing/Shareledger/x/document/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGenesisState() types.GenesisState {
	return types.GenesisState{}
}

// TODO: Validate genesis data
func ValidateGenesis(data types.GenesisState) error {
	return nil
}

func DefaultGenesisState() *types.GenesisState {
	return &types.GenesisState{}
}

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, doc := range genState.Documents {
		k.SetDoc(ctx, doc)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	docs := []*types.Document{}

	cb := func(doc types.Document) (stop bool) {
		docs = append(docs, &doc)
		return false
	}

	k.IterateDocs(ctx, cb)

	return &types.GenesisState{
		Documents: docs,
	}
}
