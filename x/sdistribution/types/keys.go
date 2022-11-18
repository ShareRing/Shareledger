package types

const (
	// ModuleName defines the module name
	ModuleName = "sdistribution"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sdistribution"

	// FeeCollectorName the root string for the fee collector account address
	FeeCollectorName = "sdistribution_fee_collector"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
