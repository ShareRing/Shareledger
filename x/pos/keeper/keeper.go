package keeper

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *wire.Codec
	//coinKeeper bank.Keeper

	// codespace
	codespace sdk.CodespaceType
}
