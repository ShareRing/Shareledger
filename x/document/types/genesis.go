package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(docs []*Document) *GenesisState {
	return &GenesisState{
		Documents: docs,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{}
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	return nil
}
