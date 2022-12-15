package types

const (
	// ModuleName defines the module name
	ModuleName = "distributionx"

	// StoreKey defines the primary module store key
	StoreKey = "X" + ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_distributionx"

	// FeeWasmName fee pool name for wasm transactions
	FeeWasmName = "fee_collector_wasm"

	// FeeNativeName fee pool name for native transactions
	FeeNativeName = "fee_collector_native"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	BuilderListKey      = "BuilderList/value/"
	BuilderListCountKey = "BuilderList/count/"
)
