package ante

import (
	assetmoduletypes "github.com/sharering/shareledger/x/asset/types"
	bookingtypes "github.com/sharering/shareledger/x/booking/types"
	gentleminttypes "github.com/sharering/shareledger/x/gentlemint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"math"
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
		feeTx, ok := tx.(sdk.FeeTx)
		if !ok {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		}
		shrTXFee := feeTx.GetFee().AmountOf(gentleminttypes.DenomSHR)
		shrRequiredFee := requiredFees.AmountOf(gentleminttypes.DenomSHR)
		if shrRequiredFee.GT(shrTXFee) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s shr required: %s shr", shrTXFee, shrRequiredFee)
		}
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
	shrAmt = int64(math.Ceil(fiatFee * exRate))
	return sdk.NewCoins(sdk.NewCoin(gentleminttypes.DenomSHR, sdk.NewInt(shrAmt)))
}
