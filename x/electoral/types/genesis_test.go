package types_test

import (
	"testing"

	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/require"
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
				AccStateList: []types.AccState{
					{
						Key: "0",
					},
					{
						Key: "1",
					},
				},
				Authority: &types.Authority{
					Address: "address",
				},
				Treasurer: &types.Treasurer{
					Address: "address",
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated accState",
			genState: &types.GenesisState{
				AccStateList: []types.AccState{
					{
						Key: "0",
					},
					{
						Key: "0",
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
