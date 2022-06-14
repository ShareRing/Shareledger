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
	SwapStatusCanceled = "cancelled"
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
	SwapTypeIn  = "in"
	SwapTypeOut = "out"
)

const (
	TxnStatusSuccess = "SUCCESS"
	TxnStatusFail    = "FAIL"
)
