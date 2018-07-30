package handlers

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/x/bank/messages"

)

//------------------------------------------------------------------
// Handler for the message

// Handle MsgSend.
// NOTE: msg.From, msg.To, and msg.Amount were already validated
// in ValidateBasic().
func HandleMsgCheck(key *sdk.KVStoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		checkMsg, ok := msg.(messages.MsgCheck)

		fmt.Printf("Received %v\n", checkMsg)

		if !ok {
			// Create custom error message and return result
			// Note: Using unreserved error codespace
			return sdk.NewError(2, 1, "MsgCheck is malformed").Result()
		}

		// Load the store.
		store := ctx.KVStore(key)

		var acc types.AppAccount
		err := utils.Retrieve(store, checkMsg.Account, &acc)
		if err != nil {
			return sdk.ErrInternal(utils.Format(constants.ERROR_STORE_RETRIEVAL,
												utils.ByteToString(checkMsg.Account),
												"bank")).Result()
		}

		// If no acc found
		if acc == (types.AppAccount{}) {
			acc = types.NewDefaultAccount()
		}

		if acc.Coins.Denom == checkMsg.Denom {
			return sdk.Result{
				Log:  fmt.Sprintf("%v", acc.Coins.Amount),
				Tags: checkMsg.Tags(),
			}
		} else {
			return sdk.Result{
				Log:  fmt.Sprintf("This account doensn't have this Coin"),
				Tags: checkMsg.Tags(),
			}
		}

	}
}
