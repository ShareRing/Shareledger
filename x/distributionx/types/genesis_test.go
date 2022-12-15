package types_test

import (
	"testing"

	"github.com/sharering/shareledger/x/distributionx/types"
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

				RewardList: []types.Reward{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				BuilderCountList: []types.BuilderCount{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				BuilderListList: []types.BuilderList{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				BuilderListCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated reward",
			genState: &types.GenesisState{
				RewardList: []types.Reward{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated builderCount",
			genState: &types.GenesisState{
				BuilderCountList: []types.BuilderCount{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated builderList",
			genState: &types.GenesisState{
				BuilderListList: []types.BuilderList{
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
			desc: "invalid builderList count",
			genState: &types.GenesisState{
				BuilderListList: []types.BuilderList{
					{
						Id: 1,
					},
				},
				BuilderListCount: 0,
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
