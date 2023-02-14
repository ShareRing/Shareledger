package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	type Test struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}
	testCases := []Test{
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
		{
			desc: "empty min gas price",
			genState: &types.GenesisState{
				Params: types.Params{
					MinimumGasPrices: sdk.DecCoins{},
				},
			},
			valid: true,
		},
		{
			desc: "no min gas price",
			genState: &types.GenesisState{
				Params: types.Params{
					MinimumGasPrices: nil,
				},
			},
			valid: true,
		},
		{
			desc: "negative min gas price mount",
			genState: &types.GenesisState{
				Params: types.Params{
					MinimumGasPrices: sdk.DecCoins{
						sdk.DecCoin{
							Denom:  "XXX",
							Amount: sdk.NewDec(-1),
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "valid min gas price mount",
			genState: &types.GenesisState{
				Params: types.Params{
					MinimumGasPrices: sdk.DecCoins{
						sdk.DecCoin{
							Denom:  "SHR",
							Amount: sdk.NewDec(100),
						},
						sdk.DecCoin{
							Denom:  "ABC",
							Amount: sdk.NewDec(1),
						},
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range testCases {
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
