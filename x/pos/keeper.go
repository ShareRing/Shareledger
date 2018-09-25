package pos

import (
	"bitbucket/shareringvn/cosmos-sdk/wire"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/bank"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *wire.Codec
	coinKeeper bank.Keeper

	// codespace
	codespace sdk.CodespaceType
}
