package types

import "fmt"

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
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	RequestCountKey = "Request-count-"
)

func RequestKey(status string) string {
	return fmt.Sprintf("Request-%s-value-", status)
}

const (
	BatchKey      = "Batch-value-"
	BatchCountKey = "Batch-count-"
)
