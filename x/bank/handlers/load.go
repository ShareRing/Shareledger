package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	Err "github.com/sharering/shareledger/x/bank/error"
	"github.com/sharering/shareledger/x/bank/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

//--------------------------------
// Handler for the message

func HandleMsgLoad(am auth.AccountMapper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		loadMsg, ok := msg.(messages.MsgLoad)
		if !ok {
			return sdkTypes.NewResult(sdk.NewError(Err.BankCodespace, Err.MsgMailedFormBank, "MsgLoad is malformed").Result())
		}

		// IMPORTANT
		// TODO: require a list of limited accounts which are priviledged to load coins

		signer := auth.GetSigner(ctx)

		// Only reserve is allowed to execute this function
		if !utils.IsValidReserve(signer.GetAddress()) {
			return sdkTypes.NewResult(sdk.ErrInternal(fmt.Sprintf(constants.RES_RESERVE_ONLY)).Result())
		}

		// Credit the account
		var resT sdk.Result

		if resT = handleTo(ctx, am, loadMsg.Account, loadMsg.Amount); !resT.IsOK() {
			return sdkTypes.NewResult(resT)
		}

		event := sdk.NewEvent(
			EventTypeLoad,
			sdk.NewAttribute(AttributeAccountAddress, loadMsg.Account.String()),
			sdk.NewAttribute(AttributeAmount, loadMsg.Amount.String()),
		)
		ctx.EventManager().EmitEvent(event)
		return sdkTypes.Result{
			Result: sdk.Result{
				Log:    resT.Log,
				Data:   resT.Data,
				Events: ctx.EventManager().Events(),
			},
		}
	}
}
