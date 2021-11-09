package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(ids []*ID) *GenesisState {
	return &GenesisState{
		IDs: ids,
	}
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

// // UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
// func (g GenesisState) UnpackInterfaces(c codecAnyUnpacker) error {
// 	for i := range g.Validators {
// 		if err := g.Validators[i].UnpackInterfaces(c); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
