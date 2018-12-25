package fee

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/exchange"
)

func NewFeeHandler(am auth.AccountMapper, exchangeKey *sdk.KVStoreKey) sdk.FeeHandler {
	return func(
		ctx sdk.Context,
		result sdk.Result,
	) (_ sdk.Result, abort bool) {
		// Several tx don't return fee

		if result.FeeDenom == "" && result.FeeAmount == 0 {
			return result, false
		}

		txFee := types.NewCoin(result.FeeDenom, result.FeeAmount)

		keeper := bank.NewKeeper(am)

		signer := auth.GetSigner(ctx).GetAddress()

		// abort due to fee has invalid denom or negative amount
		if !(txFee.HasValidDenom() && txFee.IsNotNegative()) {
			return sdk.ErrInternal(fmt.Sprintf(constants.INVALID_TX_FEE, txFee)).Result(),
				true
		}

		signerCoins := keeper.GetCoins(ctx, signer)

		// if Account is less than txFee
		if signerCoins.LT(txFee) {

			deltaCoins := signerCoins.Minus(txFee)
			deltaCoin := deltaCoins.GetCoin(txFee.Denom).Neg()

			exchangeKeeper := exchange.NewKeeper(exchangeKey, keeper)

			err := exchangeKeeper.BuyCoin(
				ctx,
				signer,
				utils.StringToAddress(constants.DEFAULT_RESERVE),
				constants.EXCHANGABLE_FEE_DENOM,
				result.FeeDenom,
				deltaCoin.Amount, // only buy the difference
			)

			if err != nil {
				return sdk.ErrInternal(fmt.Sprintf(constants.INSUFFICIENT_BALANCE, err)).Result(),
					true
			}

		}

		// Subtract fee to tx signer
		_, err := keeper.SubtractCoin(ctx, signer, txFee)

		// Insufficient coin
		if err != nil {
			return sdk.ErrInternal(fmt.Sprintf(constants.INSUFFICIENT_BALANCE, err)).Result(),
				true
		}

		// if everything succeed, original result
		return result, false
	}

}
