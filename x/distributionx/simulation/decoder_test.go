package simulation_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/x/distributionx/simulation"
	"github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func TestNewStoreDecoder(t *testing.T) {
	cdc := app.MakeTestEncodingConfig().Codec
	dec := simulation.NewStoreDecoder(cdc)

	builderCountA := types.BuilderCount{
		Index: "share1",
		Count: 2,
	}

	builderCountB := types.BuilderCount{
		Index: "share2",
		Count: 2,
	}
	builderCountABz := cdc.MustMarshal(&builderCountA)
	builderCountBBz := cdc.MustMarshal(&builderCountB)

	rewardA := types.Reward{
		Index:  "share111",
		Amount: []sdk.Coin{sdk.NewCoin(denom.Base, sdk.NewInt(10))},
	}

	rewardB := types.Reward{
		Index:  "share2332",
		Amount: []sdk.Coin{sdk.NewCoin(denom.Base, sdk.NewInt(100))},
	}
	rewardAbz := cdc.MustMarshal(&rewardA)
	rewardBbz := cdc.MustMarshal(&rewardB)

	builderListA := types.BuilderList{
		Id:              0,
		ContractAddress: "sharexxx",
	}

	builderListB := types.BuilderList{
		Id:              1,
		ContractAddress: "sharexxsx",
	}

	builderListABz := cdc.MustMarshal(&builderListA)
	builderListBBz := cdc.MustMarshal(&builderListB)

	builderListCountA := make([]byte, 8)
	builderListCountB := make([]byte, 8)

	binary.BigEndian.PutUint64(builderListCountA, 54)
	binary.BigEndian.PutUint64(builderListCountB, 54)

	tests := []struct {
		name        string
		kvA, kvB    kv.Pair
		expectedLog string
		wantPanic   bool
	}{
		{
			name: "Builder list count",
			kvA: kv.Pair{
				Key:   []byte(types.BuilderCountKeyPrefix),
				Value: builderCountABz,
			},
			kvB: kv.Pair{
				Key:   []byte(types.BuilderCountKeyPrefix),
				Value: builderCountBBz,
			},
			expectedLog: fmt.Sprintf("BuilderCount A: %v /n B: %v", builderCountA, builderCountB),
			wantPanic:   false,
		},
		{
			name: "Reward A-B",
			kvA: kv.Pair{
				Key:   []byte(types.RewardKeyPrefix),
				Value: rewardAbz,
			},
			kvB: kv.Pair{
				Key:   []byte(types.RewardKeyPrefix),
				Value: rewardBbz,
			},
			expectedLog: fmt.Sprintf("Reward A: %v /n B: %v", rewardA, rewardB),
		},
		{
			name: "Builder list A-B",
			kvA: kv.Pair{
				Key:   []byte(types.BuilderListKey),
				Value: builderListABz,
			},
			kvB: kv.Pair{
				Key:   []byte(types.BuilderListKey),
				Value: builderListBBz,
			},
			expectedLog: fmt.Sprintf("BuilderList A: %v /n B: %v", builderListA, builderListB),
		},
		{
			name: "Builder list count A-B",
			kvA: kv.Pair{
				Key:   []byte(types.BuilderListCountKey),
				Value: builderListCountA,
			},
			kvB: kv.Pair{
				Key:   []byte(types.BuilderListCountKey),
				Value: builderListCountB,
			},
			expectedLog: fmt.Sprintf("BuilderListCount A: %v /n B: %v", 54, 54),
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
