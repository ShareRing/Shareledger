package simulation_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/x/swap/simulation"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func TestNewDecodeStore(t *testing.T) {
	cdc := app.MakeTestEncodingConfig().Codec
	dec := simulation.NewDecodeStore(cdc)

	requestA := types.Request{
		Id:          2,
		SrcAddr:     "shareledger1r5adtd7fe7j8jnl8y3jf243xa6axstjwmpgyjt",
		DestAddr:    "0xxxxxxxxxx",
		SrcNetwork:  "shareledger",
		DestNetwork: "bsc",
		Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(10000)),
		Fee:         sdk.NewCoin(denom.Base, sdk.NewInt(100)),
		Status:      "pending",
	}
	requestB := types.Request{
		Id:          4,
		SrcAddr:     "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4",
		DestAddr:    "0xxxxxxxxxx",
		SrcNetwork:  "shareledger",
		DestNetwork: "bsc",
		Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(2200)),
		Fee:         sdk.NewCoin(denom.Base, sdk.NewInt(200)),
		Status:      "pending",
	}

	requestABz, err := cdc.Marshal(&requestA)
	require.NoError(t, err)

	requestBBz, err := cdc.Marshal(&requestB)
	require.NoError(t, err)

	batchA := types.Batch{
		Id:         12,
		Signature:  "b4b0359e0a5acaa5e7133cdc4dfa8ecf",
		RequestIds: []uint64{34, 23},
		Status:     "pending",
		Network:    "bsc",
	}

	batchB := types.Batch{
		Id:         12,
		Signature:  "2d051cfcba9dda43a1f7fef67723ff94",
		RequestIds: []uint64{54, 34, 23},
		Status:     "pending",
		Network:    "eth",
	}

	batchABz, err := cdc.Marshal(&batchA)
	require.NoError(t, err)

	batchBBz, err := cdc.Marshal(&batchB)
	require.NoError(t, err)

	batchCountAbz := make([]byte, 8)
	binary.BigEndian.PutUint64(batchCountAbz, 23)

	batchCountBbz := make([]byte, 8)
	binary.BigEndian.PutUint64(batchCountBbz, 24)

	pastTxEventA := types.PastTxEvent{
		SrcAddr:  "src1",
		DestAddr: "dest1",
	}

	pastTxEventB := types.PastTxEvent{
		SrcAddr:  "src2",
		DestAddr: "dest2",
	}

	pastTxEventABz, err := cdc.Marshal(&pastTxEventA)
	require.NoError(t, err)
	pastTxEventBBz, err := cdc.Marshal(&pastTxEventB)
	require.NoError(t, err)

	tests := []struct {
		name        string
		kvA, kvB    kv.Pair
		expectedLog string
		wantPanic   bool
	}{
		{
			name: "Request A-B",
			kvA: kv.Pair{
				Key:   types.KeyPrefix(types.RequestKey("pending")),
				Value: requestABz,
			},
			kvB: kv.Pair{
				Key:   types.KeyPrefix(types.RequestKey("pending")),
				Value: requestBBz,
			},
			wantPanic:   false,
			expectedLog: fmt.Sprintf("RequestA: %v\nRequestB: %v", requestA, requestB),
		},
		{
			name: "Batch A-B",
			kvA: kv.Pair{
				Key:   types.KeyPrefix(types.BatchKey),
				Value: batchABz,
			},
			kvB: kv.Pair{
				Key:   types.KeyPrefix(types.BatchKey),
				Value: batchBBz,
			},
			wantPanic:   false,
			expectedLog: fmt.Sprintf("BatchA: %v\nBatchB: %v", batchA, batchB),
		},
		{
			name: "RequestCount A-B",
			kvA: kv.Pair{
				Key:   types.KeyPrefix(types.BatchCountKey),
				Value: batchCountAbz,
			},
			kvB: kv.Pair{
				Key:   types.KeyPrefix(types.BatchCountKey),
				Value: batchCountBbz,
			},
			wantPanic:   false,
			expectedLog: fmt.Sprintf("BatchCountA: %v\nBatchCountB: %v", 23, 24),
		},
		{
			name: "Past transaction event A-b",
			kvA: kv.Pair{
				Key:   append(types.KeyPrefix(types.PastTxEventsKeyPrefix), types.PastTxEventKey("hash1", 2)...),
				Value: pastTxEventABz,
			},
			kvB: kv.Pair{
				Key:   append(types.KeyPrefix(types.PastTxEventsKeyPrefix), types.PastTxEventKey("hash1", 2)...),
				Value: pastTxEventBBz,
			},
			wantPanic:   false,
			expectedLog: fmt.Sprintf("PastTXEventA: %v\nPastTXEventB: %v", pastTxEventA, pastTxEventB),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				require.Panics(t, func() { dec(tt.kvA, tt.kvB) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(tt.kvA, tt.kvB), tt.name)
			}
		})
	}
}
