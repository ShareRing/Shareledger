package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ExchangeRate:       nil,
		LevelFeeList:       []LevelFee{},
		ActionLevelFeeList: []ActionLevelFee{},
		Params:             DefaultParams(),
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

	if err := gs.Params.ValidateBasic(); err != nil {
		return fmt.Errorf("failed validate genesis params: %w", err)
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

// GetGenesisStateFromAppState returns x/gentlemint GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
