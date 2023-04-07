package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/sharering/shareledger/x/electoral/types"
)

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, []byte(types.AuthorityKey)):
			var authorityA types.Authority
			var authorityB types.Authority

			err := cdc.Unmarshal(kvA.Value, &authorityA)
			if err != nil {
				panic(err)
			}
			err = cdc.Unmarshal(kvB.Value, &authorityB)
			if err != nil {
				panic(err)
			}
			return fmt.Sprintf("Authority A: %v\nB: %v", authorityA, authorityB)
		case bytes.Equal(kvA.Key, []byte(types.TreasurerKey)):
			var treasureA types.Treasurer
			var treasureB types.Treasurer

			err := cdc.Unmarshal(kvA.Value, &treasureA)
			if err != nil {
				panic(err)
			}
			err = cdc.Unmarshal(kvB.Value, &treasureB)
			if err != nil {
				panic(err)
			}
			return fmt.Sprintf("Treasure A: %s\nB: %v", treasureA, treasureB)
		case bytes.Equal(kvA.Key, types.KeyPrefix(types.AccStateKeyPrefix)):
			var accA types.AccState
			var accB types.AccState

			err := cdc.Unmarshal(kvA.Value, &accA)
			if err != nil {
				panic(err)
			}
			err = cdc.Unmarshal(kvB.Value, &accB)
			if err != nil {
				panic(err)
			}
			return fmt.Sprintf("AccState A: %s\nB: %v", accA, accB)
		default:
			panic(fmt.Sprintf("Key %s not support", string(kvA.Key)))
		}
	}
}
