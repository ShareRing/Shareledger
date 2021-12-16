package types

const (
	// ModuleName defines the module name
	ModuleName = "electoral"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_electoral"
)

type Status string

var (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

var (
	// Event enroll acc op
	EventTypeEnrollAccOp = "enroll_account_operator"
	EventTypeRevokeAccOp = "revoke_account_operator"
	// Event Type
	EventTypeEnrollIdSigner = "enroll_id_signer"
	EventTypeRevokeIdSigner = "revoke_id_signer"

	EventTypeRevokeDocIssuer = "revoke_doc_issuer"
	EventTypeEnrollDocIssuer = "enroll_doc_issuer"
	// Attr
	EventAttrAddress = "address"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AuthorityKey = "Authority-value-"
)

const (
	TreasurerKey = "Treasurer-value-"
)
