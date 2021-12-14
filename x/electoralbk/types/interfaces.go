package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type GentlemintKeeper interface {
	GetAuthorityAccount(ctx sdk.Context) string
}
