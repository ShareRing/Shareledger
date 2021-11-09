package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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

	DefaultParamspace = ModuleName

	MAX_ID_LEN      = 64
	MAX_ID_IN_BATCH = 10
)

const (
	QueryPathAddress = "address"
	QueryPathId      = "id"
)

// ID events
const (
	EventCreateID       = "create_id"
	EventCreateIDBatch  = "create_id_batch"
	EventUpdateID       = "update_id"
	EventReplaceIDOwner = "replce_id_owner"
)

const (
	EventAttrIssuer = "issuer"
	EventAttrId     = "id"
	EventAttrOwner  = "owner"
)

var (
	IdAddressStatePrefix = []byte{0x1}
	IdStatePrefix        = []byte{0x2}
)

// ID message types
const (
	TypeMsgCreateID       = "create_id"
	TypeMsgCreateIDBatch  = "create_id_batch"
	TypeMsgUpdateID       = "update_id"
	TypeMsgReplaceIdOwner = "replace_id_owner"
)

func IdAddressStateStoreKey(addr sdk.AccAddress) []byte {
	return append(IdAddressStatePrefix, addr.Bytes()...)
}
func IdStateStoreKey(id string) []byte {
	return append(IdStatePrefix, []byte(id)...)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
