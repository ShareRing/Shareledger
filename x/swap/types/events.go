package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

const (
	AttributeValueCategory = ModuleName

	EventTypeApproveRequests     = "approve_requests"
	EventTypeRejectRequests      = "reject_requests"
	EventTypeCancelRequest       = "cancel_request"
	EventTypeCompleteBatch       = "complete_batch"
	EventTypeCancelBatch         = "cancel_batch"
	EventTypeCreateRequest       = "create_request"
	EventTypeChangeRequestStatus = "change_request_status"

	EventTypeDeposit  = "deposit"
	EventTypeWithdraw = "withdraw"

	EventAttrSwapTotalAmount = "total_amount"
	EventAttrSwapFee         = "fee"
	EventAttrSwapDestAddr    = "dest_addr"
	EventAttrSwapSrcAddr     = "src_addr"
	EventAttrSwapDestNetwork = "dest_network"
	EventAttrSwapSrcNetwork  = "src_network"
	EventAttrSwapId          = "request_id"
	EventAttrSwapIds         = "request_ids"
	EventAttrSwapType        = "type"
	EventAttrSwapCreator     = "creator"
	EventAttrBatchId         = "batch_id"
	EventAttrBatchIds        = "batch_ids"
	EventAttrStatus          = "status"

	EventAttrWithdrawAmount   = "amount"
	EventAttrDepositAmount    = EventAttrWithdrawAmount
	EventAttrWithdrawReceiver = "receiver"
)

// NewApproveRequestsEvent  constructs a new approve sdk.Event
func NewApproveRequestsEvent(creator string, batchId uint64, reqID []string, totalAmount sdk.DecCoins) sdk.Event {
	return sdk.NewEvent(
		EventTypeApproveRequests,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrBatchId, fmt.Sprintf("%v", batchId)),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
		sdk.NewAttribute(EventAttrSwapTotalAmount, totalAmount.String()),
	)
}

// NewCreateRequestsEvent  constructs a new reject list of  a swap request sdk.Event
func NewCreateRequestsEvent(creator string, reqID uint64, totalAmount, fee sdk.Coins, srcAddr, srcNet, destAddr, destNet string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreateRequest,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapId, fmt.Sprintf("%v", reqID)),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrSwapTotalAmount, totalAmount.String()),
		sdk.NewAttribute(EventAttrSwapFee, fee.String()),
		sdk.NewAttribute(EventAttrSwapSrcAddr, srcAddr),
		sdk.NewAttribute(EventAttrSwapDestAddr, destAddr),
		sdk.NewAttribute(EventAttrSwapSrcNetwork, srcNet),
		sdk.NewAttribute(EventAttrSwapDestNetwork, destNet),
	)
}

// NewRejectRequestsEvent  constructs a new reject sdk.Event
func NewRejectRequestsEvent(creator string, reqID []string, totalAmount sdk.Coin) sdk.Event {
	return sdk.NewEvent(
		EventTypeRejectRequests,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrBatchId, ""),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
		sdk.NewAttribute(EventAttrSwapTotalAmount, totalAmount.String()),
	)
}

// NewApproveInEvent  constructs a new approve in sdk.Event
func NewApproveInEvent(creator string, reqIDs []string, totalAmount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		EventTypeApproveRequests,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrBatchId, ""),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqIDs, ",")),
		sdk.NewAttribute(EventAttrSwapTotalAmount, totalAmount.String()),
	)
}

// NewCancelRequestEvent  constructs a new cancel sdk.Event
func NewCancelRequestEvent(creator string, reqID []string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCancelRequest,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
	)
}

// NewCompleteBatchEvent  constructs a new complete sdk.Event
func NewCompleteBatchEvent(creator string, batchID uint64, reqID []string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCompleteBatch,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
		sdk.NewAttribute(EventAttrBatchId, fmt.Sprintf("%v", batchID)),
	)
}

// NewCancelBatchEvent  constructs a new cancel sdk.Event
func NewCancelBatchEvent(creator string, batchIDs, reqID []string) sdk.Event {
	return sdk.NewEvent(
		EventTypeCancelBatch,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapCreator, creator),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
		sdk.NewAttribute(EventAttrBatchIds, strings.Join(batchIDs, ",")),
	)
}

// NewChangeRequestStatusesEvent  constructs a new change status sdk.Event
func NewChangeRequestStatusesEvent(reqID []string, status string) sdk.Event {
	return sdk.NewEvent(
		EventTypeChangeRequestStatus,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(EventAttrSwapIds, strings.Join(reqID, ",")),
		sdk.NewAttribute(EventAttrStatus, status),
	)
}
