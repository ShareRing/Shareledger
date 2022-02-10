package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

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
	DenomPSHR = "pshr"
	DenomSHRP = "shrp"
	DenomCent = "cent"
)

var (
	RequiredPSHRAmt     = sdk.NewInt(10 * denom.ShrExponent)
	MaxPSHRSupply       = sdk.NewInt(4396000000 * denom.ShrExponent)
	DefaultExchangeRate = sdk.NewDec(200).Mul(sdk.NewDec(denom.ShrExponent))
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
	OneShr          = sdk.NewCoins(sdk.NewCoin(DenomPSHR, sdk.NewInt(denom.ShrExponent)))
	OneShrP         = sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)))
	OneHundredCents = sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(100)))
	FeeLoadSHRP     = OneShr
	AllowanceLoader = sdk.NewCoins(sdk.NewCoin(DenomPSHR, sdk.NewInt(20*denom.ShrExponent)))
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
	ExchangeRateKey = "exchangeRate"
)
