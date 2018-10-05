package keeper

import (
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	bank "github.com/sharering/shareledger/x/bank"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *wire.Codec
	bankKeeper bank.Keeper

	// codespace
	codespace sdk.CodespaceType
}

func NewKeeper(posKey sdk.StoreKey, bk bank.Keeper, cdc *wire.Codec) Keeper {
	keeper := Keeper{
		storeKey:   posKey,
		cdc:        cdc,
		bankKeeper: bk,
	}
	return keeper
}
