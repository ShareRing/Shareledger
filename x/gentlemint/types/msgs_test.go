package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestAddShrpCoins(t *testing.T) {
	type testCase struct {
		ic sdk.Coins
		ia sdk.Coins
		oa AdjustmentCoins
		oe error
		d  string
	}
	testCases := []testCase{
		{
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			},
			d: "1.5 shrp + 0.5 shrp = 2 shrp -> add: 1 shrp, sub 5 cent",
		},
		{
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(70))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(30))),
			},
			d: "1.5 shrp + 0.7 shrp = 2.2 -> add: 1 shrp, sub: 3 cent",
		},
		{
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(550))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(70))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(6))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(530))),
			},
			d: "1.550 shrp + 0.7 shrp = 7.2 -> add: 6 shrp, sub: 530 cent",
		},
	}

	for _, tc := range testCases {
		r, e := AddShrpCoins(tc.ic, tc.ia)
		if tc.oe != nil {
			require.NotNil(t, e, tc.d)
			require.True(t, sdkerrors.IsOf(e, tc.oe), tc.d)
		} else {
			require.Equal(t, tc.oa, r, tc.d)
		}
	}
}

func TestGetCostShrpForShr(t *testing.T) {
	rate := float64(200)
	type testCase struct {
		iCurrent sdk.Coins
		iNeed    sdk.Int
		iRate    float64
		oCost    AdjustmentCoins
		oError   error
		d        string
	}
	testCases := []testCase{
		{
			iCurrent: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
			),
			iRate: rate,
			iNeed: sdk.NewInt(200),
			oCost: AdjustmentCoins{
				Sub: sdk.NewCoins(
					sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
				),
				Add: sdk.NewCoins(),
			},
			oError: nil,
			d:      "1 shrp buy 200 shr (1 shrp) -> cost -1 shrp",
		},
		{
			iCurrent: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
			),
			iRate: rate,
			iNeed: sdk.NewInt(200),
			oCost: AdjustmentCoins{
				Sub: sdk.NewCoins(
					sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
				),
				Add: sdk.NewCoins(),
			},
			oError: nil,
			d:      "1.1 shrp buy 200 shr (1 shrp) -> cost: -1 shrp",
		},
		{
			iCurrent: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(2)),
				sdk.NewCoin(DenomCent, sdk.NewInt(10)),
			),
			iRate: rate,
			iNeed: sdk.NewInt(300),
			oCost: AdjustmentCoins{
				Sub: sdk.NewCoins(
					sdk.NewCoin(DenomSHRP, sdk.NewInt(2)),
				),
				Add: sdk.NewCoins(
					sdk.NewCoin(DenomCent, sdk.NewInt(50)),
				),
			},
			oError: nil,
			d:      "2.10 shrp buy 300 shr (1.5 shrp)-> new: 0.6 shrp -> cost: - 2 shrp +50 cent",
		},
	}
	for i, tc := range testCases {
		if i == 2 {
			fmt.Println(i)
		}
		o, e := GetCostShrpForShr(tc.iCurrent, tc.iNeed, tc.iRate)
		require.Equal(t, tc.oCost.Sub, o.Sub, tc.d)
		require.Equal(t, tc.oCost.Add, o.Add, tc.d)
		if tc.oError != nil {
			require.NotNil(t, e, tc.d)
			require.True(t, sdkerrors.IsOf(e, tc.oError), tc.d)
		} else {
			require.Nil(t, e)
		}
	}

}

func TestSubShrpCoins(t *testing.T) {
	// o, e := SubShrpCoins(x, y)
	type testCase struct {
		x sdk.Coins
		y sdk.Coins
		o sdk.Coins
		e error
		d string
	}
	testCases := []testCase{
		{
			x: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(0)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(0)),
			),
			o: sdk.NewCoins(),
			d: "0 - 0 = 0",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(5)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(4)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
			),
			d: "5 - 4 = 1",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(5)),
				sdk.NewCoin(DenomCent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(4)),
				sdk.NewCoin(DenomCent, sdk.NewInt(3)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
				sdk.NewCoin(DenomCent, sdk.NewInt(1)),
			),
			d: "5.04 - 4.03 = 1.01",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(5)),
				sdk.NewCoin(DenomCent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(4)),
				sdk.NewCoin(DenomCent, sdk.NewInt(5)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(DenomCent, sdk.NewInt(99)),
			),
			d: "5 shrp 4 cent - 4 shrp 5 cent = 0 shrp 99 cent",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(5)),
				sdk.NewCoin(DenomCent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(5)),
				sdk.NewCoin(DenomCent, sdk.NewInt(6)),
			),
			e: sdkerrors.ErrInsufficientFunds,
			d: "5 shrp 4 cent - 5 shrp 6 cent = error",
		},
	}

	for _, tc := range testCases {
		v, err := SubShrpCoins(tc.x, tc.y)
		require.ElementsMatch(t, tc.o, v, tc.d)
		require.Equal(t, tc.e, err, tc.d)
	}
}

func TestParseShrpCoinsStr(t *testing.T) {
	type testCase struct {
		i  string
		o  sdk.Coins
		oe error
		d  string
	}
	testCases := []testCase{
		{
			i: "0",
			o: sdk.NewCoins(),
			d: "0 shrp",
		},
		{
			i: "0.01",
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(1))),
			d: "0.01 shrp -> 1 cent",
		},
		{
			i: "1.1",
			o: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
				sdk.NewCoin(DenomCent, sdk.NewInt(1)),
			),
			d: "1.1 shrp -> 1shrp 1 cent",
		},
		{
			i: "0.01",
			o: sdk.NewCoins(
				sdk.NewCoin(DenomCent, sdk.NewInt(1)),
			),
			d: "0.01 shrp -> 1 cent",
		},
		{
			i: "1",
			o: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(1)),
			),
			d: "1 -> 1 shrp",
		},
		{
			i:  "-1",
			oe: sdkerrors.ErrInvalidCoins,
			d:  "negative -> err",
		},
		{
			i:  "1.100",
			oe: sdkerrors.ErrInvalidCoins,
			d:  "100 cent -> err",
		},
	}
	for i, tc := range testCases {
		r, err := ParseShrpCoinsStr(tc.i)
		require.Equal(t, tc.o, r, fmt.Sprintf("%s. test index %v", tc.d, i))
		if tc.oe != nil {
			require.NotNil(t, err, tc.d)
			require.True(t, sdkerrors.IsOf(err, tc.oe), tc.d)
		}
	}
}
