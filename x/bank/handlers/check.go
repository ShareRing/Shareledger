package handlers

import (
	//"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank/messages"
)

//------------------------------------------------------------------
// Handler for the message

// Handle MsgSend.
// NOTE: msg.From, msg.To, and msg.Amount were already validated
// in ValidateBasic().
func HandleMsgCheck(am auth.AccountMapper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		checkMsg, ok := msg.(messages.MsgCheck)

		if !ok {
			// Create custom error message and return result
			// Note: Using unreserved error codespace
			return sdk.NewError(2, 1, "MsgCheck is malformed").Result()
		}

		account := am.GetAccount(ctx, checkMsg.Account)
		if account != nil {
			return sdk.Result{
				Log:  account.String(),
				Tags: checkMsg.Tags(),
			}
		} else {
			shrAcc := auth.NewSHRAccountWithAddress(checkMsg.Account)
			account = shrAcc
			return sdk.Result{
				Log:  account.String(),
				Tags: checkMsg.Tags(),
			}
		}

	}
}
