package types

const (
	IDTypeTransaction = "transaction"
	IDTypeBatch       = "batch"
)

const (
	NetworkNameShareLedger = "SLP3"
)

// pending -> approved|rejected -> processing -> done
const (
	SwapStatusPending    = "pending"
	SwapStatusApproved   = "approved"
	SwapStatusReject     = "rejected"
	SwapStatusProcessing = "processing"
	SwapStatusDone       = "done"
)

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
