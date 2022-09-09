package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// ModuleName defines the module name
	ModuleName = "swap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_swap"

	Seperator = "/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func RequestKey(status string) string {
	return fmt.Sprintf("Request-%s-value-", status)
}

const (
	RequestCountKey = "Request-count-"
	BatchKey        = "Batch-value-"
	BatchCountKey   = "Batch-count-"

	// SignSchemaKeyPrefix is the prefix to retrieve all Format
	SignSchemaKeyPrefix = "Schemas/"

	// PastTxEventsKeyPrefix is the prefix to retrieve all PastTxEvent
	PastTxEventsKeyPrefix = "PastTxEvents/"
)

// FormatKey returns the store key to retrieve a Format from the index fields
// network/ -> value
func FormatKey(
	network string,
) []byte {
	key := []byte{}

	key = append(key, []byte(Seperator)...)
	key = append([]byte(network), key...)

	return key
}

// PastTxEventKey returns the store key to retrieve a PastTxEvent from the index fields
// txhash/logindex/ -> value
func PastTxEventKey(
	txHash string,
	logIndex uint64,
) []byte {
	key := []byte{}

	key = append(key, []byte(Seperator)...)
	key = append(key, []byte(fmt.Sprintf("%d", logIndex))...)
	key = append(key, []byte(Seperator)...)
	key = append([]byte(txHash), key...)

	return key
}

// filter value by txhash
func PastTxEventByTxHashKey(txHash string) []byte {
	key := []byte{}

	key = append(key, []byte(Seperator)...)
	key = append([]byte(txHash), key...)

	return key
}
