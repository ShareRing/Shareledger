package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	Err "github.com/sharering/shareledger/x/bank/error"
	"github.com/sharering/shareledger/x/bank/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

//------------------------------------------------------------------
// Handler for the message

// Handle MsgTransferShr.
// NOTE: msg.From, msg.To, and msg.Amount were already validated
// in ValidateBasic().
func HandleMsgTransferShr(am auth.AccountMapper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		transferShrMsg, ok := msg.(messages.MsgTransferShr)

		if !ok {
			// Create custom error message and return result
			// Note: Using unreserved error codespace
			return sdkTypes.NewResult(sdk.NewError(Err.BankCodespace, Err.MsgMailedFormBank, "MsgTransferShr is malformed").Result())
		}

		// Get signer from signatures
		signer := auth.GetSigner(ctx)

		// Debit from the sender.
		var resF sdk.Result
		var resT sdk.Result

		// From account is deduced from signature
		if resF = handleFromTransferShr(ctx, am, signer.GetAddress(), transferShrMsg.Amount); !resF.IsOK() {
			return sdkTypes.NewResult(resF)
		}

		// Credit the receiver.
		if resT = handleToTransferShr(ctx, am, transferShrMsg.To, transferShrMsg.Amount); !resT.IsOK() {
			return sdkTypes.NewResult(resT)
		}

		res := fmt.Sprintf("{\"from\":%v, \"to\":%v}", resF.Log, resT.Log)
		// Return a success (Code 0).
		// Add list of key-value pair descriptors ("tags").
		fee, denom := utils.GetMsgFee(msg)
		event := sdk.NewEvent(
			EventTypeTransferShr,
			sdk.NewAttribute(AttributeToAddress, transferShrMsg.To.String()),
			sdk.NewAttribute(AttributeAmount, transferShrMsg.Amount.String()),
			sdk.NewAttribute(AttributeEvent, ValueTransfered),
			sdk.NewAttribute(AttributeFromAddress, signer.GetAddress().String()),
		)
		ctx.EventManager().EmitEvent(event)
		return sdkTypes.Result{
			Result: sdk.Result{
				Log:    res,
				Data:   append(resF.Data, resT.Data...),
				Events: ctx.EventManager().Events(),
			},
			FeeAmount: fee,
			FeeDenom:  denom,
		}
	}
}

// Convenience Handlers
func handleFromTransferShr(ctx sdk.Context, am auth.AccountMapper, from sdk.AccAddress, amt types.Coin) sdk.Result {

	acc := am.GetAccount(ctx, from)

	// In case there is no associate account
	if acc == nil {
		shrAcc := auth.NewSHRAccountWithAddress(from)
		acc = shrAcc
	}

	// Deduct msg amount from sender account.
	senderCoins := acc.GetCoins()

	senderCoinsAfter := senderCoins.Minus(amt)

	// If any coin has negative amount, return insufficient coins error.
	if !senderCoinsAfter.IsNotNegative() {
		return sdk.ErrInsufficientCoins("Insufficient coins in account").Result()
	}

	// Set acc coins to new amount.
	acc.SetCoins(senderCoinsAfter)

	// Save to AccountMapper
	am.SetAccount(ctx, acc)

	return sdk.Result{Log: acc.GetCoins().String()}
}

func handleToTransferShr(ctx sdk.Context, am auth.AccountMapper, to sdk.AccAddress, amt types.Coin) sdk.Result {
	// Add msg amount to receiver account
	acc := am.GetAccount(ctx, to)

	// In case there is no associate account
	if acc == nil {
		shrAcc := auth.NewSHRAccountWithAddress(to)
		acc = shrAcc
	}

	// Add amount to receiver's old coins
	receiverCoins := acc.GetCoins()
	receiverCoinsAfter := receiverCoins.Plus(amt)

	// Update receiver account
	acc.SetCoins(receiverCoinsAfter)

	// Save to AccountMapper
	am.SetAccount(ctx, acc)

	// fmt.Println("acc.GetCoins().String()=", acc.GetCoins().String())
	return sdk.Result{
		Log: acc.GetCoins().String(),
	}
}

func getTransferShrFee() types.Coin {

}
