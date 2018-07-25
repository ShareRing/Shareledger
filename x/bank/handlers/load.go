package handlers

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/bank/messages"
)

//--------------------------------
// Handler for the message

func HandleMsgLoad(key *sdk.KVStoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		loadMsg, ok := msg.(messages.MsgLoad)

		fmt.Printf("Received %v\n", loadMsg)

		if !ok {
			return sdk.NewError(2, 1, "MsgLoad is malformed").Result()
		}

		// Load the store
		store := ctx.KVStore(key)

		// Credit the account
		var resT sdk.Result
		if resT = handleTo(store, loadMsg.Account, loadMsg.Amount); !resT.IsOK() {
			return resT
		}

		return sdk.Result{
			Log:  resT.Log,
			Data: resT.Data,
			Tags: loadMsg.Tags(),
		}

	}
}
