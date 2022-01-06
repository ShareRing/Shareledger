package types

const (
	// ModuleName defines the module name
	ModuleName = "id"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_id"

	// ID message types
	TypeMsgCreateID       = "create_id"
	TypeMsgCreateIDs      = "create_ids"
	TypeMsgUpdateID       = "update_id"
	TypeMsgReplaceIdOwner = "replace_id_owner"

	MAX_ID_LEN      = 64
	MAX_ID_IN_BATCH = 20

	AddressKeyPrefix = "Address/"
	IDKeyPrefix      = "ID/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
