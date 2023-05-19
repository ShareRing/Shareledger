package electoral

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sharering/shareledger/x/electoral/simulation"
	"github.com/sharering/shareledger/x/electoral/types"
)

// GenerateGenesisState creates a randomized GenState of the module
func (am AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.MustGenRandGenesis(simState)
}

// ProposalContents doesn't return any content functions for governance proposals
func (am AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.NewWeightedOperations(simState, am.keeper, am.gk, am.ak, am.bk)
}
