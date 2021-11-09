package id

import (
	"github.com/ShareRing/Shareledger/x/id/keeper"
	"github.com/ShareRing/Shareledger/x/id/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type GenesisState struct {
// 	IDs []types.ID `json:"IDs" yaml:"IDs"`
// }

func NewGenesisState() types.GenesisState {
	return types.GenesisState{}
}

// TODO: Validate genesis data
func ValidateGenesis(data types.GenesisState) error {
	return nil
}

func DefaultGenesisState() types.GenesisState {
	return types.GenesisState{}
}

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, id := range genState.IDs {
		k.SetID(ctx, id)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	ids := []*types.ID{}

	cb := func(id types.ID) (stop bool) {
		ids = append(ids, &id)
		return false
	}

	k.IterateID(ctx, cb)

	return &types.GenesisState{
		IDs: ids,
	}
}
