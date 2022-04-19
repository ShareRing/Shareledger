package types_test

import (
	"testing"

	"github.com/sharering/shareledger/x/swap/types"
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

				IdList: []types.Id{
					{
						IDType: "0",
					},
					{
						IDType: "1",
					},
				},
				RequestList: []types.Request{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				RequestCount: 2,
				BatchList: []types.Batch{
	{
		Id: 0,
	},
	{
		Id: 1,
	},
},
BatchCount: 2,
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated id",
			genState: &types.GenesisState{
				IdList: []types.Id{
					{
						IDType: "0",
					},
					{
						IDType: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated request",
			genState: &types.GenesisState{
				RequestList: []types.Request{
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
				RequestList: []types.Request{
					{
						Id: 1,
					},
				},
				RequestCount: 0,
			},
			valid: false,
		},
		{
	desc:     "duplicated batch",
	genState: &types.GenesisState{
		BatchList: []types.Batch{
			{
				Id: 0,
			},
			{
				Id: 0,
			},
		},
	},
	valid:    false,
},
{
	desc:     "invalid batch count",
	genState: &types.GenesisState{
		BatchList: []types.Batch{
			{
				Id: 1,
			},
		},
		BatchCount: 0,
	},
	valid:    false,
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
