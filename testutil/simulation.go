package testutil

import (
	"fmt"
	"math/rand"

	"github.com/sharering/shareledger/x/swap/types"
	"github.com/thanhpk/randstr"
)

func RandBool(r *rand.Rand) bool {
	return r.Intn(2) == 1
}

func RandRate(r *rand.Rand) int64 {
	return r.Int63()
}

// RandNetwork generate the random network eht or bsc
func RandNetwork(r *rand.Rand) string {
	if RandBool(r) {
		return types.NetworkNameEthereum
	}
	return types.NetworkNameBinanceSmartChain
}

// RandEthAddress generate the random ethereum base address
func RandEthAddress() string {
	return fmt.Sprintf("0x%s", randstr.Hex(40))
}

func RandPick[T any](r *rand.Rand, src []T) T {
	idx := r.Intn(len(src))
	return src[idx]
}

func RandERC20Event(r *rand.Rand) []*types.TxEvent {
	num := r.Intn(5-1) + 1
	txs := make([]*types.TxEvent, num)
	for i := 0; i < len(txs); i++ {
		txs[i] = &types.TxEvent{
			TxHash:   randstr.Hex(45),
			Sender:   fmt.Sprintf("0x%s", randstr.Hex(40)),
			LogIndex: uint64(r.Intn(20)),
		}
	}
	return txs
}
