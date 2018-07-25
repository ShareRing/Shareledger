package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank/messages"
)

//------------------------------------------------------------------
// Handler for the message

// Handle MsgSend.
// NOTE: msg.From, msg.To, and msg.Amount were already validated
// in ValidateBasic().
func HandleMsgSend(key *sdk.KVStoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		sendMsg, ok := msg.(messages.MsgSend)

		fmt.Println("Received %v", sendMsg)

		if !ok {
			// Create custom error message and return result
			// Note: Using unreserved error codespace
			return sdk.NewError(2, 1, "MsgSend is malformed").Result()
		}

		// Load the store.
		store := ctx.KVStore(key)

		// Debit from the sender.
		var resF sdk.Result
		var resT sdk.Result
		if resF = handleFrom(store, sendMsg.From, sendMsg.Amount); !resF.IsOK() {
			return resF
		}

		// Credit the receiver.
		if resT = handleTo(store, sendMsg.To, sendMsg.Amount); !resT.IsOK() {
			return resT
		}
		fmt.Println("Result Log:", resF.Log, "-", resT.Log)
		res := fmt.Sprintf("From:%v - To:%v", resF.Log, resT.Log)
		// Return a success (Code 0).
		// Add list of key-value pair descriptors ("tags").
		return sdk.Result{
			Log:  res,
			Data: append(resF.Data, resT.Data...),
			Tags: sendMsg.Tags(),
		}
	}
}

// Convenience Handlers
func handleFrom(store sdk.KVStore, from sdk.Address, amt types.Coin) sdk.Result {
	// Get sender account from the store.
	accBytes := store.Get(from)
	//if accBytes == nil {
	// Account was not added to store. Return the result of the error.
	//return sdk.NewError(2, 101, "Account not added to store").Result()
	//}

	// Unmarshal the JSON account bytes.
	var acc types.AppAccount
	if accBytes != nil {
		err := json.Unmarshal(accBytes, &acc)

		if err != nil {
			// InternalError
			return sdk.ErrInternal("Error when deserializing account").Result()
		}
	} else {
		acc = types.AppAccount{
			Coins: types.NewCoin("SHR", 0),
		}
	}

	// Deduct msg amount from sender account.
	senderCoins := acc.Coins.Minus(amt)

	// If any coin has negative amount, return insufficient coins error.
	if !senderCoins.IsNotNegative() {
		return sdk.ErrInsufficientCoins("Insufficient coins in account").Result()
	}

	fmt.Println("Deduce From:", acc.Coins.Amount, " ", senderCoins.Amount)
	// Set acc coins to new amount.
	acc.Coins = senderCoins

	// Encode sender account.
	naccBytes, nerr := json.Marshal(acc)
	if nerr != nil {
		return sdk.ErrInternal("Account encoding error").Result()
	}

	// Update store with updated sender account
	store.Set(from, naccBytes)
	return sdk.Result{Data: naccBytes,
		Log: strconv.FormatInt(acc.Coins.Amount, 10)}
	//Log: "Result" }
}

func handleTo(store sdk.KVStore, to sdk.Address, amt types.Coin) sdk.Result {
	// Add msg amount to receiver account
	accBytes := store.Get(to)
	var acc types.AppAccount
	if accBytes == nil {
		// Receiver account does not already exist, create a new one.
		acc = types.AppAccount{}
	} else {
		// Receiver account already exists. Retrieve and decode it.
		err := json.Unmarshal(accBytes, &acc)
		if err != nil {
			return sdk.ErrInternal("Account decoding error").Result()
		}
	}

	// Add amount to receiver's old coins
	receiverCoins := acc.Coins.Plus(amt)

	// Update receiver account
	acc.Coins = receiverCoins
	fmt.Println("After Plus:", acc.Coins.Amount)

	// Encode receiver account
	accBytes, err := json.Marshal(acc)
	if err != nil {
		return sdk.ErrInternal("Account encoding error").Result()
	}

	// Update store with updated receiver account
	store.Set(to, accBytes)
	return sdk.Result{Log: strconv.FormatInt(acc.Coins.Amount, 10)}
}
