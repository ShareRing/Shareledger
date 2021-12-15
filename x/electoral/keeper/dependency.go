package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type GentlemintKeeper interface {
	LoadAllowanceLoader(ctx sdk.Context, addr sdk.AccAddress) error
}
