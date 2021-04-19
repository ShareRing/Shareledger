package gatecheck

import (
	"github.com/ShareRing/modules/document"
	doctypes "github.com/ShareRing/modules/document/types"
	"github.com/ShareRing/modules/id"
	idtypes "github.com/ShareRing/modules/id/types"
	shareringUtils "github.com/ShareRing/modules/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/booking"
	"github.com/sharering/shareledger/x/gentlemint"
)

const (
	FEE_LEVEL_HIGH = "HIGH"
	// HIGHFEE          = 0.05
	FEE_LEVEL_MEDIUM = "MEDIUM"
	// MEDIUMFEE        = 0.03
	FEE_LEVEL_LOW = "LOW"
	// LOWFEE           = 0.02
	// MINFEE           = 0.01
	SKIP_CHECK_LEVEL = "SKIP"
)

type CheckFeeDecorator struct {
	gk gentlemint.Keeper
}

func NewCheckFeeDecorator(gk gentlemint.Keeper) CheckFeeDecorator {
	return CheckFeeDecorator{
		gk: gk,
	}
}

func (cfd CheckFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	msg := msgs[0]
	feeLevel, multiplier := getFeeLevel(msg)
	if feeLevel == SKIP_CHECK_LEVEL {
		return next(ctx, tx, simulate)
	}
	exRate := cfd.gk.GetExchangeRate(ctx)
	// exRate, err := strconv.ParseFloat(exRateStr, 64)
	// if err != nil {
	// 	return ctx, err
	// }

	requiredFees := getRequiredFees(feeLevel, exRate, int64(multiplier))

	feeTx, _ := tx.(ante.FeeTx)
	fee := feeTx.GetFee()
	if requiredFees.IsAnyGT(fee) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", fee, requiredFees)
	}
	return next(ctx, tx, simulate)
}

func getFeeLevel(msg sdk.Msg) (string, int) {
	switch msg.Type() {
	case asset.TypeAssetCreateMsg, booking.TypeBookMsg, document.TypeMsgCreateDoc, document.TypeMsgRevokeDoc:
		return FEE_LEVEL_HIGH, 1
	case asset.TypeAssetUpdateMsg, booking.TypeBookCompleteMsg, document.TypeMsgUpdateDoc, id.TypeMsgCreateID, id.TypeMsgUpdateID, id.TypeMsgReplaceIdOwner:
		return FEE_LEVEL_MEDIUM, 1
	case asset.TypeAssetDeleteMsg, "send", gentlemint.TypesSendSHRP, gentlemint.TypeSendSHR:
		return FEE_LEVEL_LOW, 1
	case document.TypeMsgCreateDocInBatch:
		m := msg.(doctypes.MsgCreateDocBatch)
		return FEE_LEVEL_HIGH, len(m.Holder)
	case id.TypeMsgCreateIDBatch:
		m := msg.(idtypes.MsgCreateIdBatch)
		return FEE_LEVEL_HIGH, len(m.Id)
	default:
		return SKIP_CHECK_LEVEL, 1
	}
}

// Calculate the SHR amount based on the fee level and the exchange rate
func getRequiredFees(feeLevel string, exRate sdk.Int, multiplier int64) sdk.Coins {
	var fiatFee sdk.Int
	switch feeLevel {
	case FEE_LEVEL_HIGH:
		fiatFee = shareringUtils.HIGHFEE
	case FEE_LEVEL_MEDIUM:
		fiatFee = shareringUtils.MEDIUMFEE
	case FEE_LEVEL_LOW:
		fiatFee = shareringUtils.LOWFEE
	default:
		fiatFee = shareringUtils.MINFEE
	}

	shrAmt := fiatFee.Mul(exRate)
	return sdk.NewCoins(sdk.NewCoin("shr", shrAmt.Mul(sdk.NewInt(multiplier))))
}
