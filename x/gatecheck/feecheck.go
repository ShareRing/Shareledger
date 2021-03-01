package gatecheck

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/booking"
	"github.com/sharering/shareledger/x/gentlemint"
)

const (
	FEE_LEVEL_HIGH   = "HIGH"
	HIGHFEE          = 0.05
	FEE_LEVEL_MEDIUM = "MEDIUM"
	MEDIUMFEE        = 0.03
	FEE_LEVEL_LOW    = "LOW"
	LOWFEE           = 0.02
	MINFEE           = 0.01
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
	feeLevel := getFeeLevel(msg)
	if feeLevel == SKIP_CHECK_LEVEL {
		return next(ctx, tx, simulate)
	}
	exRateStr := cfd.gk.GetExchangeRate(ctx)
	exRate, err := strconv.ParseFloat(exRateStr, 64)
	if err != nil {
		return ctx, err
	}
	requiredFees := getRequiredFees(feeLevel, exRate)

	feeTx, _ := tx.(ante.FeeTx)
	fee := feeTx.GetFee()
	if requiredFees.IsAnyGT(fee) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", fee, requiredFees)
	}
	return next(ctx, tx, simulate)
}

func getFeeLevel(msg sdk.Msg) string {
	switch msg.Type() {
	case asset.TypeAssetCreateMsg, booking.TypeBookMsg:
		return FEE_LEVEL_HIGH
	case asset.TypeAssetUpdateMsg, booking.TypeBookCompleteMsg:
		return FEE_LEVEL_MEDIUM
	case asset.TypeAssetDeleteMsg, "send", gentlemint.TypesSendSHRP, gentlemint.TypeSendSHR:
		return FEE_LEVEL_LOW
	default:
		return SKIP_CHECK_LEVEL
	}
}

func getRequiredFees(feeLevel string, exRate float64) sdk.Coins {
	var shrAmt int64
	var fiatFee float64
	switch feeLevel {
	case FEE_LEVEL_HIGH:
		fiatFee = HIGHFEE
	case FEE_LEVEL_MEDIUM:
		fiatFee = MEDIUMFEE
	case FEE_LEVEL_LOW:
		fiatFee = LOWFEE
	default:
		fiatFee = MINFEE
	}
	shrAmt = int64(fiatFee*exRate) + 1
	return sdk.NewCoins(sdk.NewCoin("shr", sdk.NewInt(shrAmt)))
}
