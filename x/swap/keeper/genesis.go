package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
)

// ImportRequest import data into keeper.
// This function will ignore all validation check, only use for GENESIS operations
func (k Keeper) ImportRequest(ctx sdk.Context, requests []types.Request) {
	stores := k.GetStoreRequestMap(ctx)
	for _, request := range requests {
		appendedValue := k.cdc.MustMarshal(&request)
		store, found := stores[request.Status]
		if !found {
			panic(fmt.Sprintf("the status of imported request is invalid. Value %+v", request))
		}
		store.Set(GetRequestIDBytes(request.Id), appendedValue)
	}
}
