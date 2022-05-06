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

	opWeightMsgDeposit = "op_weight_msg_deposit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeposit int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	opWeightMsgCreateFormat = "op_weight_msg_format"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFormat int = 100

	opWeightMsgUpdateFormat = "op_weight_msg_format"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFormat int = 100

	opWeightMsgDeleteFormat = "op_weight_msg_format"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFormat int = 100

	opWeightMsgCancel = "op_weight_msg_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancel int = 100

	opWeightMsgReject = "op_weight_msg_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReject int = 100

	opWeightMsgIn = "op_weight_msg_in"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIn int = 100

	opWeightMsgApproveIn = "op_weight_msg_approve_in"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveIn int = 100

	opWeightMsgUpdateBatch = "op_weight_msg_update_batch"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBatch int = 100

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
		FormatList: []types.Format{
		{
			Creator: sample.AccAddress(),
Network: "0",
},
		{
			Creator: sample.AccAddress(),
Network: "1",
},
	},
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

	var weightMsgDeposit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeposit, &weightMsgDeposit, nil,
		func(_ *rand.Rand) {
			weightMsgDeposit = defaultWeightMsgDeposit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeposit,
		swapsimulation.SimulateMsgDeposit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		swapsimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = defaultWeightMsgCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancel,
		swapsimulation.SimulateMsgCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReject, &weightMsgReject, nil,
		func(_ *rand.Rand) {
			weightMsgReject = defaultWeightMsgReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReject,
		swapsimulation.SimulateMsgReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgIn, &weightMsgIn, nil,
		func(_ *rand.Rand) {
			weightMsgIn = defaultWeightMsgIn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIn,
		swapsimulation.SimulateMsgIn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveIn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveIn, &weightMsgApproveIn, nil,
		func(_ *rand.Rand) {
			weightMsgApproveIn = defaultWeightMsgApproveIn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveIn,
		swapsimulation.SimulateMsgApproveIn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateFormat int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateFormat, &weightMsgCreateFormat, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFormat = defaultWeightMsgCreateFormat
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFormat,
		swapsimulation.SimulateMsgCreateFormat(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFormat int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateFormat, &weightMsgUpdateFormat, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFormat = defaultWeightMsgUpdateFormat
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFormat,
		swapsimulation.SimulateMsgUpdateFormat(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFormat int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteFormat, &weightMsgDeleteFormat, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFormat = defaultWeightMsgDeleteFormat
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFormat,
		swapsimulation.SimulateMsgDeleteFormat(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBatch int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBatch, &weightMsgUpdateBatch, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBatch = defaultWeightMsgUpdateBatch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBatch,
		swapsimulation.SimulateMsgUpdateBatch(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
