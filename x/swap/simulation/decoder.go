package simulation

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/sharering/shareledger/x/swap/types"
)

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.SignSchemaKeyPrefix)):
			var (
				schemaA types.Schema
				schemaB types.Schema
			)
			err := cdc.Unmarshal(kvA.Value, &schemaA)
			panicIfErr(err)

			err = cdc.Unmarshal(kvB.Value, &schemaB)
			panicIfErr(err)
			return fmt.Sprintf("SchemaA: %v\nSchemaB: %v", schemaA, schemaB)
		case bytes.Equal(kvA.Key, []byte(types.RequestKey("pending"))),
			bytes.Equal(kvA.Key, []byte(types.RequestKey("approved"))):
			var (
				requestA types.Request
				requestB types.Request
			)
			err := cdc.Unmarshal(kvA.Value, &requestA)
			panicIfErr(err)
			err = cdc.Unmarshal(kvB.Value, &requestB)
			panicIfErr(err)
			return fmt.Sprintf("RequestA: %v\nRequestB: %v", requestA, requestB)
		case bytes.Equal(kvA.Key, []byte(types.BatchKey)):
			var (
				batchA types.Batch
				batchB types.Batch
			)
			err := cdc.Unmarshal(kvA.Value, &batchA)
			panicIfErr(err)
			err = cdc.Unmarshal(kvB.Value, &batchB)
			panicIfErr(err)
			return fmt.Sprintf("BatchA: %v\nBatchB: %v", batchA, batchB)
		case bytes.Equal(kvA.Key, []byte(types.BatchCountKey)):
			countA := binary.BigEndian.Uint64(kvA.Value)
			countB := binary.BigEndian.Uint64(kvB.Value)
			return fmt.Sprintf("BatchCountA: %v\nBatchCountB: %v", countA, countB)
		case bytes.Equal(kvA.Key, []byte(types.RequestCountKey)):
			countA := binary.BigEndian.Uint64(kvA.Value)
			countB := binary.BigEndian.Uint64(kvB.Value)
			return fmt.Sprintf("RequestCountA: %v\nRequestCountB: %v", countA, countB)
		case bytes.Equal(kvA.Key[:len([]byte(types.PastTxEventsKeyPrefix))], []byte(types.PastTxEventsKeyPrefix)):
			var (
				pastEventA types.PastTxEvent
				pastEventB types.PastTxEvent
			)
			err := cdc.Unmarshal(kvA.Value, &pastEventA)
			panicIfErr(err)
			err = cdc.Unmarshal(kvB.Value, &pastEventB)
			panicIfErr(err)
			return fmt.Sprintf("PastTXEventA: %v\nPastTXEventB: %v", pastEventA, pastEventB)
		default:
			return fmt.Sprintf("KVA_Key: %s \nKVB_Key: %s ", string(kvA.Key), string(kvB.Key))
		}
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
