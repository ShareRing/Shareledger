package swap

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sharering/shareledger/testutil/sample"
	swapsimulation "github.com/sharering/shareledger/x/swap/simulation"
	"github.com/sharering/shareledger/x/swap/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = swapsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgOut = "op_weight_msg_out"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOut int = 100

	opWeightMsgApprove = "op_weight_msg_approve"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApprove int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	swapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&swapGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	swapParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyOutFee), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(swapParams.OutFee))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgOut int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOut, &weightMsgOut, nil,
		func(_ *rand.Rand) {
			weightMsgOut = defaultWeightMsgOut
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOut,
		swapsimulation.SimulateMsgOut(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApprove int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApprove, &weightMsgApprove, nil,
		func(_ *rand.Rand) {
			weightMsgApprove = defaultWeightMsgApprove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApprove,
		swapsimulation.SimulateMsgApprove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
