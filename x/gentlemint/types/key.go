package types

const (
	// ModuleName is the name of the module
	ModuleName = "gentlemint"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey is the module name router key
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	IdSignerPrefix  = "IdSigner"
	DocIssuerPrefix = "DocIssuer"
	AccOpPrefix     = "AccOp"

	EventTypeEnrollAccOp = "enroll_account_operator"
	EventTypeRevokeAccOp = "revoke_account_operator"

	EventTypeEnrollDocIssuer = "enroll_doc_issuer"
	EventTypeRevokeDocIssuer = "revoke_doc_issuer"

	EventTypeEnrollIdSigner = "enroll_id_signer"
	EventTypeRevokeIdSigner = "revoke_id_signer"

	EventAttrAddress = "address"
)
