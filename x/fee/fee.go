package fee

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/utils"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/exchange"

	sdkTypes "github.com/sharering/shareledger/cosmos-wrapper/types"
)

type FeeHandler func(sdk.Context, sdkTypes.Result) (sdk.Result, bool)

func NewFeeHandler(am auth.AccountMapper, exchangeKey *sdk.KVStoreKey) FeeHandler {
	return func(
		ctx sdk.Context,
		result sdkTypes.Result,
	) (_ sdk.Result, abort bool) {
		// Several tx don't return fee

		if result.FeeDenom == "" && result.FeeAmount == 0 {
			// if everything succeed, original result
			result.Tags = result.Tags.
				AppendTag(FeeDenom, constants.FEE_DENOM).
				AppendTag(FeeAmount, strconv.FormatInt(int64(constants.NONE), 10)) 
			return result.CosmosResult(), false
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
		result.Tags = result.Tags.
			AppendTag(FeeDenom, result.FeeDenom).
			AppendTag(FeeAmount, strconv.FormatInt(result.FeeAmount, 10))

		return result.CosmosResult(), false
	}

}
