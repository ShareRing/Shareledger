package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sharering/shareledger/x/electoral/keeper"
	"github.com/sharering/shareledger/x/electoral/types"
	"math/rand"
)

const (
	opWeightEnrollRelayer     = "op_weight_msg_enroll_relayer"
	opWeightEnrollAccountOp   = "op_weight_msg_enroll_accountOp"
	opWeightEnrollAprover     = "op_weight_msg_enroll_approver"
	opWeightEnrollDocIssuer   = "op_weight_msg_enroll_doc_issuer"
	opWeightEnrollIdSigner    = "op_weight_msg_enroll_id_signer"
	opWeightEnrollLoader      = "op_weight_msg_enroll_loader"
	opWeightEnrollSwapManager = "op_weight_msg_enroll_swap_manager"
	opWeightEnrollVoter       = "op_weight_msg_enroll_voter"

	opWeightRevokeRelayer     = "op_weight_msg_revoke_relayer"
	opWeightRevokeAccountOp   = "op_weight_msg_revoke_accountOp"
	opWeightRevokeApprover    = "op_weight_msg_revoke_approver"
	opWeightRevokeDocIssuer   = "op_weight_msg_revoke_doc_issuer"
	opWeightRevokeIdSigner    = "op_weight_msg_revoke_id_signer"
	opWeightRevokeLoader      = "op_weight_msg_revoke_loader"
	opWeightRevokeSwapManager = "op_weight_msg_revoke_swap_manager"
	opWeightRevokeVoter       = "op_weight_msg_revoke_voter"

	defaultWeightErEnroll int = 100
	defaultWeightErRevoke int = 70
)

func NewWeightedOperations(simState module.SimulationState, k keeper.Keeper, gk keeper.GentlemintKeeper, ak types.AccountKeeper, bk types.BankKeeper) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightErRelayer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollRelayer, &weightErRelayer, nil,
		func(_ *rand.Rand) {
			weightErRelayer = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErRelayer,
		SimulateEnrollRelayer(k, gk, ak, bk),
	))

	var weightRevokeRelayer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeRelayer, &weightRevokeRelayer, nil,
		func(_ *rand.Rand) {
			weightRevokeRelayer = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeRelayer,
		SimulateRevokeRelayer(k, gk, ak, bk),
	))
	//Account Operator

	var weightErOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollAccountOp, &weightErOperator, nil,
		func(_ *rand.Rand) {
			weightErOperator = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErOperator,
		SimulateEnrollAccountOperator(k, gk, ak, bk),
	))

	var weightRevokeOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeAccountOp, &weightRevokeOperator, nil,
		func(_ *rand.Rand) {
			weightRevokeOperator = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeOperator,
		SimulateRevokeAccountOperator(k, gk, ak, bk),
	))

	// Approver
	var weightErApprover int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollAprover, &weightErApprover, nil,
		func(_ *rand.Rand) {
			weightErApprover = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErApprover,
		SimulateEnrollApprover(k, gk, ak, bk),
	))

	var weightRevokeApprover int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeApprover, &weightRevokeApprover, nil,
		func(_ *rand.Rand) {
			weightRevokeApprover = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeApprover,
		SimulateRevokeApprover(k, gk, ak, bk),
	))

	// Doc Issuer
	var weightErDocIssuer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollDocIssuer, &weightErDocIssuer, nil,
		func(_ *rand.Rand) {
			weightErDocIssuer = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErDocIssuer,
		SimulateEnrollDocIssuer(k, gk, ak, bk),
	))

	var weightRevokeDocIssuer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeDocIssuer, &weightRevokeDocIssuer, nil,
		func(_ *rand.Rand) {
			weightRevokeDocIssuer = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeDocIssuer,
		SimulateRevokeDocIssuer(k, gk, ak, bk),
	))

	// ID signer
	var weightErIdSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollIdSigner, &weightErIdSigner, nil,
		func(_ *rand.Rand) {
			weightErIdSigner = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErIdSigner,
		SimulateEnrollIdSigner(k, gk, ak, bk),
	))

	var weightRevokeIdSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeIdSigner, &weightRevokeIdSigner, nil,
		func(_ *rand.Rand) {
			weightRevokeIdSigner = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeIdSigner,
		SimulateRevokeIdSigner(k, gk, ak, bk),
	))

	// Loader
	var weightErLoader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollLoader, &weightErLoader, nil,
		func(_ *rand.Rand) {
			weightErLoader = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErLoader,
		SimulateEnrollIdSigner(k, gk, ak, bk),
	))

	var weightRevokeLoader int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeLoader, &weightRevokeLoader, nil,
		func(_ *rand.Rand) {
			weightRevokeLoader = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeLoader,
		SimulateRevokeLoader(k, gk, ak, bk),
	))

	// Swap manager
	var weightErSwapManger int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollSwapManager, &weightErSwapManger, nil,
		func(_ *rand.Rand) {
			weightErSwapManger = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErSwapManger,
		SimulateEnrollSwapManager(k, gk, ak, bk),
	))

	var weightRevokeSwapManager int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeSwapManager, &weightRevokeSwapManager, nil,
		func(_ *rand.Rand) {
			weightRevokeSwapManager = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeSwapManager,
		SimulateRevokeSwapManager(k, gk, ak, bk),
	))

	// Voter
	var weightErVoter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightEnrollVoter, &weightErVoter, nil,
		func(_ *rand.Rand) {
			weightErVoter = defaultWeightErEnroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightErVoter,
		SimulateEnrollSwapManager(k, gk, ak, bk),
	))

	var weightRevokeVoter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightRevokeVoter, &weightRevokeVoter, nil,
		func(_ *rand.Rand) {
			weightRevokeVoter = defaultWeightErRevoke
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightRevokeVoter,
		SimulateRevokeSwapManager(k, gk, ak, bk),
	))

	return operations
}
