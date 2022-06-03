package types

const (
	IDTypeTransaction = "transaction"
	IDTypeBatch       = "batch"
)

const (
	NetworkNameShareLedger       = "shareledger"
	NetworkNameEthereum          = "eth"
	NetworkNameBinanceSmartChain = "bsc"
)

// pending -> approved|rejected
// this value should be synced with the map below
const (
	SwapStatusPending  = "pending"
	SwapStatusApproved = "approved"
	SwapStatusRejected = "rejected"
	SwapStatusDone     = "done"
)

var SupportedSwapOutNetwork = map[string]struct{}{
	NetworkNameEthereum:          {},
	NetworkNameBinanceSmartChain: {},
}
var SupportedSwapStatuses = map[string]struct{}{
	SwapStatusPending:  {},
	SwapStatusApproved: {},
	SwapStatusRejected: {},
	SwapStatusDone:     {},
}

const (
	BatchStatusPending = "pending"
)

func SwapStatusSupported(status string) bool {
	_, found := SupportedSwapStatuses[status]
	return found
}

const (
	EventTypeSwapApprove         = "swap_approve"
	EventTypeSwapReject          = "swap_reject"
	EventTypeSwapOut             = "swap_out"
	EventTypeSwapIn              = "swap_int"
	EventTypeDeposit             = "swap_deposit"
	EventTypeWithDraw            = "swap_withdraw"
	EventTypeRequestChangeStatus = "swap_request_change_status"
	EventTypeRequestCancelStatus = "swap_cancel_request"

	EventTypeBatchDone   = "batch_done"
	EventTypeBatchCancel = "batch_cancel"
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

	EventTypeAttrBatchID = "batch_id"
)

const (
	TxnStatusSuccess = "SUCCESS"
	TxnStatusFail    = "FAIL"
)
