package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccStateList: []AccState{},
		Authority:    nil,
		Treasurer:    nil,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in accState
	accStateIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccStateList {
		index := string(AccStateKey(IndexKeyAccState(elem.Key)))
		if _, ok := accStateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for accState")
		}
		accStateIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
