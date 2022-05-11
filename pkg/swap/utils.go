package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	swaptypes "github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"math/big"
)

func BuildTypedData(signSchemaData apitypes.TypedData, requests []swaptypes.Request) (apitypes.TypedData, error) {
	txIds := make([]interface{}, 0, len(requests))
	destinations := make([]interface{}, 0, len(requests))
	amounts := make([]interface{}, 0, len(requests))
	for _, tx := range requests {
		txIds = append(txIds, (*math.HexOrDecimal256)(new(big.Int).SetUint64(tx.Id)))
		destinations = append(destinations, tx.DestAddr)
		bCoin, err := denom.NormalizeToBaseCoins(sdk.NewDecCoinsFromCoins(*tx.Amount), false)
		if err != nil {
			return signSchemaData, err
		}
		amounts = append(amounts, (*math.HexOrDecimal256)(big.NewInt(bCoin.AmountOf(denom.Base).Int64())))
	}
	signSchemaData.Message = apitypes.TypedDataMessage{
		"ids":     txIds,
		"tos":     destinations,
		"amounts": amounts,
	}
	return signSchemaData, nil
}
