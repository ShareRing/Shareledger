package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sharering/shareledger/x/distributionx/types"
)

func validParams() types.Params {
	tp := types.DefaultParams()
	tp.DevPoolAccount = "shareledger18l057pgdpuccl7u4uf8nh69xnjpz7az28c0gvk"
	return tp
}

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is not valid",
			genState: types.DefaultGenesis(),
			valid:    false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: validParams(),
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
				Params: validParams(),
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
				Params: validParams(),
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
				Params: validParams(),
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
				Params: validParams(),
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
