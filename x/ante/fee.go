package ante

import (
	assetmoduletypes "github.com/ShareRing/Shareledger/x/asset/types"
	bookingtypes "github.com/ShareRing/Shareledger/x/booking/types"
	gentleminttypes "github.com/ShareRing/Shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	FEE_LEVEL_HIGH   = "HIGH"
	HIGHFEE          = 0.05 // usd
	FEE_LEVEL_MEDIUM = "MEDIUM"
	MEDIUMFEE        = 0.03
	FEE_LEVEL_LOW    = "LOW"
	LOWFEE           = 0.02
	MINFEE           = 0.01
	SKIP_CHECK_LEVEL = "SKIP"
)

type CheckFeeDecorator struct {
	gk GentlemintKeeper
}

func NewCheckFeeDecorator(gk GentlemintKeeper) CheckFeeDecorator {
	return CheckFeeDecorator{
		gk: gk,
	}
}

func (cfd CheckFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	msg := msgs[0]
	feeLevel := getFeeLevel(msg)
	if feeLevel != SKIP_CHECK_LEVEL {
		exRate := cfd.gk.GetExchangeRateF(ctx)
		requiredFees := getRequiredFees(feeLevel, exRate)
		ctx = ctx.WithMinGasPrices(sdk.NewDecCoinsFromCoins(requiredFees...))
	}
	return next(ctx, tx, simulate)
}

func getFeeLevel(msg sdk.Msg) string {
	switch msg.(type) {
	case *bookingtypes.MsgBook, *assetmoduletypes.MsgCreate:
		return FEE_LEVEL_HIGH
	case *bookingtypes.MsgComplete, *assetmoduletypes.MsgUpdate:
		return FEE_LEVEL_MEDIUM
	case *assetmoduletypes.MsgDelete, *gentleminttypes.MsgSendShrp, *gentleminttypes.MsgSendShr, *banktypes.MsgSend:
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
