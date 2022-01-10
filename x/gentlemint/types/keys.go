package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "gentlemint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gentlemint"
)

const (
	DenomSHR  = "shr"
	DenomSHRP = "shrp"
	DenomCent = "cent"
)

var (
	RequiredSHRAmt      = sdk.NewInt(10)
	MaxSHRSupply        = sdk.NewInt(4396000000)
	DefaultExchangeRate = float64(200)
)

type ShrpStatus string

var (
	StatusSHRPLoaderActived   ShrpStatus = "actived"
	StatusSHRPLoaderInactived ShrpStatus = "inactived"
)

type Status string

var (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

var (
	OneShr          = sdk.NewCoins(sdk.NewCoin(DenomSHR, sdk.NewInt(1)))
	OneShrP         = sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)))
	OneHundredCents = sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(100)))
	FeeLoadSHRP     = OneShr
	AllowanceLoader = sdk.NewCoins(sdk.NewCoin(DenomSHR, sdk.NewInt(20)))
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
	ExchangeRateKey = "exchange_shrp_to_shr"
)
