package types

const (
	IDTypeTransaction = "transaction"
	IDTypeBatch       = "batch"
)

const (
	NetworkNameShareLedger = "slp3"
)

// pending -> approved|rejected
// this value should be synced with the map below
const (
	SwapStatusPending  = "pending"
	SwapStatusApproved = "approved"
	SwapStatusRejected = "rejected"
)

var SupportedSwapStatuses = map[string]struct{}{
	SwapStatusPending:  {},
	SwapStatusApproved: {},
	SwapStatusRejected: {},
}

const (
	BatchStatusPending    = "pending"
	BatchStatusProcessing = "processing"
	BatchStatusDone       = "done"
)

func SwapStatusSupported(status string) bool {
	_, found := SupportedSwapStatuses[status]
	return found
}

const (
	EventTypeSwapApprove = "swap_approve"
	EventTypeSwapOut     = "swap_out"
	EventTypeDeposit     = "swap_deposit"
)

const (
	AttributeValueCategory   = ModuleName
	EventTypeSwapAmount      = "amount"
	EventTypeSwapFee         = "fee"
	EventTypeSwapDestAddr    = "dest_addr"
	EventTypeSwapSrcAddr     = "src_addr"
	EventTypeSwapDestNetwork = "dest_network"
	EventTypeSwapSrcNetwork  = "src_network"
	EventTypeSwapId          = "swap_id"
	EventTypeBatchId         = "batch_id"
	EventTypeBatchTotal      = "batch_total"
	EventTypeApproverAction  = "approver_action"
	EventTypeApproverAddr    = "approver_addr"
	EventTypeDepositAmount   = EventTypeSwapAmount
	EventTypeDepositAddr     = "sender"
)

const (
	TxnStatusSuccess = "SUCCESS"
	TxnStatusFail    = "FAIL"
)
