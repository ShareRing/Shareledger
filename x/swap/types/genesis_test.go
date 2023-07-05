package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/x/swap/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Requests: []types.Request{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RequestCount: 2,
				Batches: []types.Batch{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				BatchCount: 2,
				Schemas: []types.Schema{
					{
						Network: "0",
					},
					{
						Network: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated request",
			genState: &types.GenesisState{
				Requests: []types.Request{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid request count",
			genState: &types.GenesisState{
				Requests: []types.Request{
					{
						Id: 1,
					},
				},
				RequestCount: 0,
			},
			valid: false,
		},

		{
			desc: "duplicated format",
			genState: &types.GenesisState{
				Schemas: []types.Schema{
					{
						Network: "0",
					},
					{
						Network: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
