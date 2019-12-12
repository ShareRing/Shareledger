package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	Err "github.com/sharering/shareledger/x/bank/error"
	"github.com/sharering/shareledger/x/bank/messages"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

//------------------------------------------------------------------
// Handler for the message

// Handle MsgTransferCoin.
// NOTE: msg.From, msg.To, and msg.Amount were already validated
// in ValidateBasic().
func HandleMsgTransferCoin(am auth.AccountMapper) sdkTypes.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdkTypes.Result {
		transferCoinMsg, ok := msg.(messages.MsgTransferCoin)
		if !ok {
			return sdkTypes.NewResult(sdk.NewError(Err.BankCodespace, Err.MsgMailedFormBank, "MsgTransferCoin is malformed").Result())
		}
		signer := auth.GetSigner(ctx)

		var resF sdk.Result
		var resT sdk.Result

		// From account is deduced
		if resF = handleFromTransferCoin(ctx, am, signer.GetAddress(), transferCoinMsg.Amount); !resF.IsOK() {
			return sdkTypes.NewResult(resF)
		}

		// Credit the receiver.
		if resT = handleToTransferCoin(ctx, am, transferCoinMsg.To, transferCoinMsg.Amount); !resT.IsOK() {
			return sdkTypes.NewResult(resT)
		}

		res := fmt.Sprintf("{\"from\":%v, \"to\":%v}", resF.Log, resT.Log)
		fee, denom := utils.GetMsgFee(msg)
		event := sdk.NewEvent(
			EventTypeTransferCoin,
			sdk.NewAttribute(AttributeToAddress, transferCoinMsg.To.String()),
			sdk.NewAttribute(AttributeAmount, transferCoinMsg.Amount.String()),
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

func handleFromTransferCoin(ctx sdk.Context, am auth.AccountMapper, from sdk.AccAddress, amt types.Coin) sdk.Result {
	acc := am.GetAccount(ctx, from)
	// In case there is no associate account
	if acc == nil {
		shrAcc := auth.NewSHRAccountWithAddress(from)
		acc = shrAcc
	}
	denom := amt.Denom
	err, feeCoin := getTransferCoinFee(denom)
	if err == nil {
		// Deduct msg amount from sender account.
		senderCoins := acc.GetCoins()
		senderCoinsAfter := senderCoins.Minus(amt)
		senderCoinsAfter = senderCoinsAfter.Minus(feeCoin)
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
	return sdk.ErrInvalidCoins("Transferring Invalid Coin. Abort").Result()
}

func handleToTransferCoin(ctx sdk.Context, am auth.AccountMapper, to sdk.AccAddress, amt types.Coin) sdk.Result {
	// Add msg amount to receiver account
	acc := am.GetAccount(ctx, to)
	// In case there is no associate account
	if acc == nil {
		shrAcc := auth.NewSHRAccountWithAddress(to)
		acc = shrAcc
	}
	denom := amt.Denom
	err, feeCoin := getTransferCoinFee(denom)
	if err == nil {

		// Add amount to receiver's old coins
		receiverCoins := acc.GetCoins()
		receiverCoinsAfter := receiverCoins.Plus(amt)

		// Update receiver account
		acc.SetCoins(receiverCoinsAfter)

		// Save to AccountMapper
		am.SetAccount(ctx, acc)

		// Add fee to feeCollector address
		feeCollectorAddress := utils.StringToAddress(constants.FEE_COLLECTOR)

		feeCollectorAcc := am.GetAccount(ctx, feeCollectorAddress)

		if feeCollectorAcc == nil {
			feeCollectorAcc = auth.NewSHRAccountWithAddress(feeCollectorAddress)
		}

		feeCollectorCoins := feeCollectorAcc.GetCoins()
		feeCollectorCoinsAfter := feeCollectorCoins.Plus(feeCoin)
		feeCollectorAcc.SetCoins(feeCollectorCoinsAfter)
		am.SetAccount(ctx, feeCollectorAcc)
		return sdk.Result{
			Log: acc.GetCoins().String(),
		}
	}
	return sdk.ErrInvalidCoins("Transferring Invalid Coin. Abort").Result()
}

func getTransferCoinFee(denom string) (error, types.Coin) {
	if denom == "SHR" {
		return nil, types.NewCoin("SHR", 4)
	} else if denom == "SHRP" {
		amount, err := types.NewDecFromStr("0.05")
		if err == nil {
			return nil, types.NewCoinFromDec("SHRP", amount)
		} else {
			return types.ErrInvalidCoins("Invalid Coin"), types.NewCoin("", 0)
		}
	}
	return types.ErrInvalidCoins("Invalid Coin"), types.NewCoin("", 0)
}
