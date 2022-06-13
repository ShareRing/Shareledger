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
var SupportedSwapStatusesStores = map[string]struct{}{
	SwapStatusPending:  {},
	SwapStatusApproved: {},
}

const (
	BatchStatusPending = "pending"
)

func SwapStatusSupported(status string) bool {
	_, found := SupportedSwapStatusesStores[status]
	return found
}

const (
	EventTypeSwapApprove         = "swap_approve"
	EventTypeSwapReject          = "swap_reject"
	EventTypeSwapCancel          = "swap_cancel"
	EventTypeSwapOut             = "swap_out"
	EventTypeSwapIn              = "swap_int"
	EventTypeDeposit             = "swap_deposit"
	EventTypeWithDraw            = "swap_withdraw"
	EventTypeRequestChangeStatus = "swap_request_change_status"

	EventTypeBatchDone   = "batch_done"
	EventTypeBatchCancel = "batch_cancel"
)

const (
	SwapTypeIn  = "in"
	SwapTypeOut = "out"
)

const (
	AttributeValueCategory   = ModuleName
	EventAttrSwapAmount      = "amount"
	EventAttrSwapFee         = "fee"
	EventAttrSwapDestAddr    = "dest_addr"
	EventAttrSwapSrcAddr     = "src_addr"
	EventAttrSwapDestNetwork = "dest_network"
	EventAttrSwapSrcNetwork  = "src_network"
	EventAttrSwapId          = "swap_id"
	EventAttrSwapIds         = "swap_ids"
	EventAttrSwapType        = "swap_type"

	EventAttrBatchId      = "batch_id"
	EventAttrBatchIds     = "batch_ids"
	EventAttrBatchTotal   = "batch_total"
	EventAttrBatchTxIDs   = "batch_txs"
	EventAttrBatchNetwork = "batch_network"

	EventAttrApproverAction   = "approver_action"
	EventAttrRejectArr        = "reject_addr"
	EventAttrApproverAddr     = "approver_addr"
	EventAttrDepositAmount    = "deposit_amount"
	EventAttrWithdrawAmount   = "withdraw_amount"
	EventAttrWithdrawReceiver = "withdraw_receiver"
	EventAttrCancelAddr       = "cancel_addr"
)

const (
	TxnStatusSuccess = "SUCCESS"
	TxnStatusFail    = "FAIL"
)
