package handlers

import (
	"fmt"
	"strconv"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank/messages"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/constants"
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

	// Unmarshal the JSON account bytes.
	var acc types.AppAccount

	err := utils.Retrieve(store, from, &acc)
	if err != nil {
		return sdk.ErrInternal(utils.Format(constants.ERROR_STORE_RETRIEVAL,
											"AppAccount",
											constants.STORE_BANK)).Result()
	}

	// In case there is no associate account
	if acc == (types.AppAccount{}) {
		acc = types.NewDefaultAccount()
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


	err = utils.Store(store, from, acc)

	if err != nil {
		return sdk.ErrInternal(utils.Format(constants.ERROR_STORE_UPDATE,
											utils.ByteToString(from),
											constants.STORE_BANK)).Result()
	}

	return sdk.Result{Log: strconv.FormatInt(acc.Coins.Amount, 10)}
}

func handleTo(store sdk.KVStore, to sdk.Address, amt types.Coin) sdk.Result {
	// Add msg amount to receiver account
	var acc types.AppAccount

	err := utils.Retrieve(store, to, &acc)
	if err != nil {
		return sdk.ErrInternal(utils.Format(constants.ERROR_STORE_RETRIEVAL,
											utils.ByteToString(to),
											constants.STORE_BANK)).Result()
	}

	// In case there is no associate account
	if acc == (types.AppAccount{}) {
		acc = types.NewDefaultAccount()
	}

	// Add amount to receiver's old coins
	receiverCoins := acc.Coins.Plus(amt)

	// Update receiver account
	acc.Coins = receiverCoins
	fmt.Println("After Plus:", acc.Coins.Amount)

	err = utils.Store(store, to, acc)

	if err != nil {
		return sdk.ErrInternal(utils.Format(constants.ERROR_STORE_UPDATE,
											utils.ByteToString(to),
											constants.STORE_BANK)).Result()
	}

	return sdk.Result{Log: strconv.FormatInt(acc.Coins.Amount, 10)}
}
