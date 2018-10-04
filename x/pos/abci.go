package pos

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/pos/keeper"
	abci "github.com/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.Validator {
	var valUpdates []abci.Validator
	//TODO: return updated Validators list
	return valUpdates
}
