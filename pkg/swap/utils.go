package swap

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func buildTypedData(signSchemaData apitypes.TypedData, params SwapParams) (apitypes.TypedData, error) {
	txIds := make([]interface{}, 0, len(params.TransactionIds))
	destinations := make([]interface{}, 0, len(params.DestAddrs))
	amounts := make([]interface{}, 0, len(params.Amounts))
	for i := 0; i < len(params.TransactionIds); i++ {
		txIds = append(txIds, (*math.HexOrDecimal256)(params.TransactionIds[i]))
		destinations = append(destinations, params.DestAddrs[i].Hex())
		amounts = append(amounts, (*math.HexOrDecimal256)(params.Amounts[i]))
	}
	signSchemaData.Message = apitypes.TypedDataMessage{
		"ids":     txIds,
		"tos":     destinations,
		"amounts": amounts,
	}
	return signSchemaData, nil
}
