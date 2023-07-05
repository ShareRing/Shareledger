package simulation

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, []byte(types.BuilderCountKeyPrefix)):
			var countA, countB types.BuilderCount
			cdc.MustUnmarshal(kvA.Value, &countA)
			cdc.MustUnmarshal(kvB.Value, &countB)

			return fmt.Sprintf("BuilderCount A: %v /n B: %v", countA, countB)
		case bytes.Equal(kvA.Key, []byte(types.RewardKeyPrefix)):
			var rewardA, rewardB types.Reward
			cdc.MustUnmarshal(kvA.Value, &rewardA)
			cdc.MustUnmarshal(kvB.Value, &rewardB)
			return fmt.Sprintf("Reward A: %v /n B: %v", rewardA, rewardB)
		case bytes.Equal(kvA.Key, []byte(types.BuilderListKey)):
			var builderListA, builderListB types.BuilderList
			cdc.MustUnmarshal(kvA.Value, &builderListA)
			cdc.MustUnmarshal(kvB.Value, &builderListB)
			return fmt.Sprintf("BuilderList A: %v /n B: %v", builderListA, builderListB)
		case bytes.Equal(kvA.Key, []byte(types.BuilderListCountKey)):
			var builderListCountA, builderListCountB uint64
			builderListCountA = binary.BigEndian.Uint64(kvA.Value)
			builderListCountB = binary.BigEndian.Uint64(kvB.Value)
			return fmt.Sprintf("BuilderListCount A: %v /n B: %v", builderListCountA, builderListCountB)
		default:
			panic(fmt.Sprintf("invalid swap key prefix %X", kvA.Key[:]))
		}
	}
}
