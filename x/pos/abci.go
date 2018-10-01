package pos

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, k Keeper) []abci.Validator {
	var valUpdates []abci.Validator
	//TODO: return updated Validators list
	return valUpdates
}
