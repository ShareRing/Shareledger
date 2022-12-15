package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RewardList:       []Reward{},
		BuilderCountList: []BuilderCount{},
		BuilderListList:  []BuilderList{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in reward
	rewardIndexMap := make(map[string]struct{})

	for _, elem := range gs.RewardList {
		index := string(RewardKey(elem.Index))
		if _, ok := rewardIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for reward")
		}
		rewardIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in builderCount
	builderCountIndexMap := make(map[string]struct{})

	for _, elem := range gs.BuilderCountList {
		index := string(BuilderCountKey(elem.Index))
		if _, ok := builderCountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for builderCount")
		}
		builderCountIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in builderList
	builderListIdMap := make(map[uint64]bool)
	builderListCount := gs.GetBuilderListCount()
	for _, elem := range gs.BuilderListList {
		if _, ok := builderListIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for builderList")
		}
		if elem.Id >= builderListCount {
			return fmt.Errorf("builderList id should be lower or equal than the last id")
		}
		builderListIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
