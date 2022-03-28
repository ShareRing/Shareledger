package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

const (
	// ModuleName defines the module name
	ModuleName = "gentlemint"

	ModuleNameAlias = "gm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gentlemint"
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
	RequiredBaseAmt              = sdk.NewCoin(denom.Base, sdk.NewInt(10*denom.ShrExponent))
	MaxBaseSupply                = sdk.NewInt(4396000000 * denom.ShrExponent)
	DefaultExchangeRateSHRPToSHR = sdk.NewDec(200)
)

var (
	AllowanceLoader = sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent)))
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
