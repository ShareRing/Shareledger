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
	DemonSHR     = "shr"
	DenomSHRP    = "shrp"
	DenomCent    = "cent"
	AuthorityKey = "A"
)

var (
	RequiredSHRAmt = sdk.NewInt(10)
	MaxSHRSupply   = sdk.NewInt(4396000000)
)

var (
	OneShr          = sdk.NewCoins(sdk.NewCoin(DemonSHR, sdk.NewInt(1)))
	OneShrP         = sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)))
	OneHundredCents = sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(100)))
	FeeLoadSHRP     = OneShr
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
