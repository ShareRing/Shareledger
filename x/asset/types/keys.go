package types

const (
	// ModuleName defines the module name
	ModuleName = "asset"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_asset"

	TypeAssetCreateMsg = "asset_create"
	TypeAssetUpdateMsg = "asset_update"
	TypeAssetDeleteMsg = "asset_delete"

	UUIDKeyPrefix = "UUID/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
