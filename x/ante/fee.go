package ante

//const (
//	FEE_LEVEL_HIGH   = "HIGH"
//	HIGHFEE          = 0.05 // usd
//	FEE_LEVEL_MEDIUM = "MEDIUM"
//	MEDIUMFEE        = 0.03
//	FEE_LEVEL_LOW    = "LOW"
//	LOWFEE           = 0.02
//	MINFEE           = 0.01
//	SKIP_CHECK_LEVEL = "SKIP"
//)
//
//type CheckFeeDecorator struct {
//	gk GentlemintKeeper
//}
//
//func NewCheckFeeDecorator(gk GentlemintKeeper) CheckFeeDecorator {
//	return CheckFeeDecorator{
//		gk: gk,
//	}
//}
//
//func (cfd CheckFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
//	feeTx, ok := tx.(sdk.FeeTx)
//	if !ok {
//		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
//	}
//	msgs := tx.GetMsgs()
//	requiredFees := sdk.NewCoins()
//	for _, msg := range msgs {
//		feeLevel := getFeeLevel(msg)
//		if feeLevel != SKIP_CHECK_LEVEL {
//			exRate := cfd.gk.GetExchangeRateF(ctx)
//			requiredFees.Add(getRequiredFees(feeLevel, exRate)...)
//		}
//	}
//	shrTXFee := feeTx.GetFee().AmountOf(gentleminttypes.DenomSHR)
//	shrRequiredFee := requiredFees.AmountOf(gentleminttypes.DenomSHR)
//	if shrRequiredFee.GT(shrTXFee) {
//		return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "got: %s shr required: %s shr", shrTXFee, shrRequiredFee)
//	}
//	return next(ctx, tx, simulate)
//}
//
//func getFeeLevel(msg sdk.Msg) string {
//	switch msg.(type) {
//	case *bookingtypes.MsgCreateBooking, *assetmoduletypes.MsgCreateAsset:
//		return FEE_LEVEL_HIGH
//	case *bookingtypes.MsgCompleteBooking, *assetmoduletypes.MsgUpdateAsset:
//		return FEE_LEVEL_MEDIUM
//	case *assetmoduletypes.MsgDeleteAsset, *gentleminttypes.MsgSendShrp, *gentleminttypes.MsgSendShr, *banktypes.MsgSend:
//		return FEE_LEVEL_LOW
//	default:
//		return SKIP_CHECK_LEVEL
//	}
//}
//
//func getRequiredFees(feeLevel string, exRate float64) sdk.Coins {
//	var shrAmt int64
//	var fiatFee float64
//	switch feeLevel {
//	case FEE_LEVEL_HIGH:
//		fiatFee = HIGHFEE
//	case FEE_LEVEL_MEDIUM:
//		fiatFee = MEDIUMFEE
//	case FEE_LEVEL_LOW:
//		fiatFee = LOWFEE
//	default:
//		fiatFee = MINFEE
//	}
//	shrAmt = int64(math.Ceil(fiatFee * exRate))
//	return sdk.NewCoins(sdk.NewCoin(gentleminttypes.DenomSHR, sdk.NewInt(shrAmt)))
//}
