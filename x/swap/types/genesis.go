package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Requests: []Request{},
		Schemas:  []Schema{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in request
	requestIdMap := make(map[uint64]bool)
	requestCount := gs.GetRequestCount()
	for _, elem := range gs.Requests {
		if _, ok := requestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for request")
		}
		if elem.Id >= requestCount {
			return fmt.Errorf("request id should be lower or equal than the last id")
		}
		requestIdMap[elem.Id] = true
	}

	// Check for duplicated index in format
	formatIndexMap := make(map[string]struct{})

	for _, elem := range gs.Schemas {
		index := string(FormatKey(elem.Network))
		if _, ok := formatIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for format")
		}
		formatIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
