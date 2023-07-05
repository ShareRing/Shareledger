package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/sharering/shareledger/x/gentlemint/types"
)

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.ActionLevelFeeKeyPrefix)):

			aLFeeA := types.ActionLevelFee{}
			aLFeeB := types.ActionLevelFee{}

			cdc.MustUnmarshal(kvA.Value, &aLFeeA)
			cdc.MustUnmarshal(kvB.Value, &aLFeeB)
			return fmt.Sprintf("Action Level Fee A: %s\nB:%s", aLFeeA, aLFeeB)
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.LevelFeeKeyPrefix)):
			lFeeA := types.LevelFee{}
			lFeeB := types.LevelFee{}

			cdc.MustUnmarshal(kvA.Value, &lFeeA)
			cdc.MustUnmarshal(kvB.Value, &lFeeB)

			return fmt.Sprintf("Level Fee A: %s\nB:%s", lFeeA, lFeeA)
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.ExchangeRateKey)):
			exchangeRateA := types.ExchangeRate{}
			exchangeRateB := types.ExchangeRate{}

			cdc.MustUnmarshal(kvA.Value, &exchangeRateA)
			cdc.MustUnmarshal(kvB.Value, &exchangeRateB)

			return fmt.Sprintf("Exchange Rate A: %s\nB:%s", exchangeRateA, exchangeRateB)
		default:
			panic(fmt.Sprintf("Key: %s unsuported", kvA.String()))
		}

	}
}
