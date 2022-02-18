package types_test

import (
	"testing"

	"github.com/sharering/shareledger/x/gentlemint/types"
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
				ExchangeRate: &types.ExchangeRate{
					Rate: "200",
				},
				LevelFeeList: []types.LevelFee{
					{
						Level: "0",
					},
					{
						Level: "1",
					},
				},
				ActionLevelFeeList: []types.ActionLevelFee{
					{
						Action: "0",
					},
					{
						Action: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated levelFee",
			genState: &types.GenesisState{
				LevelFeeList: []types.LevelFee{
					{
						Level: "0",
					},
					{
						Level: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated actionLevelFee",
			genState: &types.GenesisState{
				ActionLevelFeeList: []types.ActionLevelFee{
					{
						Action: "0",
					},
					{
						Action: "0",
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
