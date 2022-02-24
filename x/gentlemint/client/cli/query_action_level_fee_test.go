package cli_test

import (
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"strconv"
	"testing"

	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithActionLevelFeeObjects(t *testing.T, n int) (*network.Network, []types.ActionLevelFee) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		state.ActionLevelFeeList = append(state.ActionLevelFeeList, types.ActionLevelFee{
			Action: strconv.Itoa(i),
		})
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ActionLevelFeeList
}
