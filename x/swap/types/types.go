package types

const (
	IDTypeTransaction = "transaction"
	IDTypeBatch       = "batch"
)

const (
	NetworkNameShareLedger = "SLP3"
)

// pending -> approved|rejected -> processing -> done
// this value should be synced with the map below
const (
	SwapStatusPending    = "pending"
	SwapStatusApproved   = "approved"
	SwapStatusReject     = "rejected"
	SwapStatusProcessing = "processing"
	SwapStatusDone       = "done"
)

var SupportedSwapStatuses = map[string]struct{}{
	SwapStatusPending:    {},
	SwapStatusApproved:   {},
	SwapStatusReject:     {},
	SwapStatusProcessing: {},
	SwapStatusDone:       {},
}

func SwapStatusSupported(status string) bool {
	_, found := SupportedSwapStatuses[status]
	return found
}

const (
	AttributeValueCategory   = ModuleName
	EventTypeSwapAmount      = "amount"
	EventTypeSwapFee         = "fee"
	EventTypeSwapId          = "id"
	EventTypeSwapDestAddr    = "dest_addr"
	EventTypeSwapSrcAddr     = "src_addr"
	EventTypeSwapDestNetwork = "dest_network"
	EventTypeSwapSrcNetwork  = "src_network"
)
