package types

const (
	// ModuleName defines the module name
	ModuleName = "document"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_document"

	TypeMsgCreateDoc  = "create_doc"
	TypeMsgCreateDocs = "create_docs"
	TypeMsgUpdateDoc  = "update_doc"
	TypeMsgRevokeDoc  = "revoke_doc"

	DocDetailKeyPrefix = "DocumentDetail/"
	DocBasicKeyPrefix  = "DocumentBasic/"

	DocRevokeFlag = 0xffff
	Separator     = "/"

	MAX_LEN       = 64
	MAX_LEN_BATCH = 20
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
