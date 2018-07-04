package handlers

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/bank/messages"
	"github.com/sharering/shareledger/types"

	"encoding/json"
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


		accBytes := store.Get(checkMsg.Account)

		var acc types.AppAccount
		if accBytes != nil {
			err := json.Unmarshal(accBytes, &acc)

			if err != nil {
				// InternalError
				return sdk.ErrInternal("Error when deserializing account").Result()
			}
		} else {
			acc = types.AppAccount{
				Coins: types.NewCoin("SHR", 0 ),
			}
		}

		if acc.Coins.Denom == checkMsg.Denom {
			return sdk.Result{
				Log: fmt.Sprintf("%v", acc.Coins.Amount),
				Tags: checkMsg.Tags(),
			}
		} else {
			return sdk.Result{
				Log: fmt.Sprintf("This account doensn't have this Coin"),
				Tags: checkMsg.Tags(),
			}
		}

	}
}