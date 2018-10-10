package fee

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
)

func NewFeeHandler(am auth.AccountMapper) sdk.FeeHandler {
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
		if !(txFee.HasValidDenoms() && txFee.IsNotNegative()) {
			return sdk.ErrInternal(fmt.Sprintf(constants.INVALID_TX_FEE, txFee)).Result(),
				true
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
