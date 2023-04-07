package distributionx

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sharering/shareledger/testutil/sample"
	distributionxSimulation "github.com/sharering/shareledger/x/distributionx/simulation"
	"github.com/sharering/shareledger/x/distributionx/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = distributionxSimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgWithdrawReward = "op_weight_msg_withdraw_reward"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawReward int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	distributionxSimulation.RandomizedGenState(simState)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(str sdk.StoreDecoderRegistry) {
	str[types.StoreKey] = distributionxSimulation.NewStoreDecoder(am.cdc)
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgWithdrawReward int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawReward, &weightMsgWithdrawReward, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawReward = defaultWeightMsgWithdrawReward
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawReward,
		distributionxSimulation.SimulateMsgWithdrawReward(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
