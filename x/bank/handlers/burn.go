package handlers

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	Err "github.com/sharering/shareledger/x/bank/error"
	"github.com/sharering/shareledger/x/bank/messages"
)

//--------------------------------
// Handler for the message

func HandleMsgBurn(am auth.AccountMapper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		burnMsg, ok := msg.(messages.MsgBurn)
		if !ok {
			return sdk.NewError(Err.BankCodespace, Err.MsgMailedFormBank, "MsgBurn is malformed").Result()
		}

		// IMPORTANT
		// TODO: require a list of limited accounts which are priviledged to load coins
		signer := auth.GetSigner(ctx)

		// Only reserve is allowed to execute this function
		if !utils.IsValidReserve(signer.GetAddress()) {
			return sdk.ErrInternal(fmt.Sprintf(constants.RES_RESERVE_ONLY)).Result()
		}

		if !bytes.Equal(signer.GetAddress(), burnMsg.Account) {
			return sdk.ErrInternal(fmt.Sprintf(constants.RES_OWN_ACCOUNT, burnMsg.Account, signer.GetAddress())).Result()
		}

		// Credit the account
		var resT sdk.Result

		if resT = handleFrom(ctx, am, burnMsg.Account, burnMsg.Amount); !resT.IsOK() {
			return resT
		}
		return sdk.Result{
			Log:  resT.Log,
			Data: resT.Data,
			Tags: burnMsg.Tags(),
		}

	}
}
