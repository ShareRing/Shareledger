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
	SwapStatusDone     = "done"
)

var SupportedSwapStatuses = map[string]struct{}{
	SwapStatusPending:  {},
	SwapStatusApproved: {},
	SwapStatusRejected: {},
	SwapStatusDone:     {},
}

const (
	BatchStatusPending = "pending"
	BatchStatusDone    = "done"
)

func SwapStatusSupported(status string) bool {
	_, found := SupportedSwapStatuses[status]
	return found
}

const (
	EventTypeSwapApprove         = "swap_approve"
	EventTypeSwapReject          = "swap_reject"
	EventTypeSwapOut             = "swap_out"
	EventTypeDeposit             = "swap_deposit"
	EventTypeWithDraw            = "swap_withdraw"
	EventTypeRequestChangeStatus = "swap_request_change_status"
	EventTypeRequestCancelStatus = "swap_cancel_request"
)

const (
	AttributeValueCategory                = ModuleName
	EventTypeSwapAmount                   = "amount"
	EventTypeSwapFee                      = "fee"
	EventTypeSwapDestAddr                 = "dest_addr"
	EventTypeSwapSrcAddr                  = "src_addr"
	EventTypeSwapDestNetwork              = "dest_network"
	EventTypeSwapSrcNetwork               = "src_network"
	EventTypeSwapId                       = "swap_id"
	EventTypeBatchId                      = "batch_id"
	EventTypeBatchTotal                   = "batch_total"
	EventTypeApproverAction               = "approver_action"
	EventTypeRejectArr                    = "reject_addr"
	EventTypeApproverAddr                 = "approver_addr"
	EventTypeDepositAmount                = EventTypeSwapAmount
	EventTypeDepositAddr                  = "sender"
	EventTypeWithdrawReceiver             = "receiver"
	EventTypeChangeRequestStatusNewStatus = "new_status"
)

const (
	TxnStatusSuccess = "SUCCESS"
	TxnStatusFail    = "FAIL"
)
