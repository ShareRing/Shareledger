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

	DefaultParamspace = ModuleName

	MAX_LEN       = 64
	MAX_LEN_BATCH = 10
)

const (
	TypeMsgCreateDoc        = "create_doc"
	TypeMsgCreateDocInBatch = "create_doc_batch"
	TypeMsgUpdateDoc        = "update_doc"
	TypeMsgRevokeDoc        = "revoke_doc"
)

const (
	QueryByProof  = "proof"
	QueryByHolder = "holder"
)
const (
	EventTypeCreateDoc = "create_doc"
	EventTypeUpdateDoc = "update_doc"
	EventTypeRevokeDoc = "revoke_doc"

	EventAttrHolder = "holder"
	EventAttrProof  = "proof"
	EventAttrIssuer = "issuer"
	EventAttrData   = "data"
)

var (
	StateKeySep     = "|"
	DocDetailPrefix = []byte{0x1}
	DocBasicPrefix  = []byte{0x2}
	DocRevokeFlag   = 0xffff
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
