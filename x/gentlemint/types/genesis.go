package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ExchangeRate:       nil,
		LevelFeeList:       []LevelFee{},
		ActionLevelFeeList: []ActionLevelFee{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	// Check for duplicated index in levelFee
	levelFeeIndexMap := make(map[string]struct{})

	for _, elem := range gs.LevelFeeList {
		index := string(LevelFeeKey(elem.Level))
		if _, ok := levelFeeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for levelFee")
		}
		levelFeeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionLevelFee
	actionLevelFeeIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionLevelFeeList {
		index := string(ActionLevelFeeKey(elem.Action))
		if _, ok := actionLevelFeeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionLevelFee")
		}
		actionLevelFeeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
