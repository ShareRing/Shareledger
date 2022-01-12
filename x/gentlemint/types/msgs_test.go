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
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(0))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(0))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(0))),
			},
			d: "1 shrp + 1 shrp = 2 shrp -> add: 1 shrp, sub 0 cent",
		},
		{
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(2))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			},
			d: "1.5 shrp + 1.5 shrp = 3 shrp -> add: 2 shrp, sub 50 cent",
		},
		{
			ic: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(50))),
			ia: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(70))),
			oa: AdjustmentCoins{
				Add: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(2))),
				Sub: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(30))),
			},
			d: "1.5 shrp + 1.7 shrp = 3.2 shrp -> add: 1 shrp, sub 30 cent",
		},
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
	rate := sdk.NewDec(200)
	type testCase struct {
		iCurrent sdk.Coins
		iNeed    sdk.Int
		iRate    sdk.Dec
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
		{
			iCurrent: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(2)),
				sdk.NewCoin(DenomCent, sdk.NewInt(10)),
			),
			iRate: rate,
			iNeed: sdk.NewInt(30000000),
			oCost: AdjustmentCoins{
				Sub: nil,
				Add: nil,
			},
			oError: sdkerrors.ErrInsufficientFunds,
			d:      "insufficient funds",
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
				sdk.NewCoin(DenomSHRP, sdk.NewInt(9)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(DenomSHRP, sdk.NewInt(9)),
			),
			o: sdk.NewCoins(),
			d: "9.99 - 9.99 = 0",
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
				sdk.NewCoin(DenomCent, sdk.NewInt(10)),
			),
			d: "1.1 shrp -> 1shrp 10 cent",
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

func TestShrToShrp(t *testing.T) {
	rate := sdk.NewDec(200)
	type testCase struct {
		i sdk.Coin
		o sdk.Coins
		d string
	}
	tcs := []testCase{
		{
			i: sdk.NewCoin(DenomSHR, sdk.NewInt(1)),
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(1))),
			d: "1 shr -> 1 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(DenomSHR, sdk.NewInt(3)),
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(2))),
			d: "3 shr -> 2 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(DenomSHR, sdk.NewInt(4)),
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(2))),
			d: "4 shr -> 2 cent",
		},
	}
	for _, tc := range tcs {
		r := ShrToShrp(tc.i, rate)
		require.Equal(t, tc.o, r, tc.d)
	}
}

func TestShrpDecCoinsToCoins(t *testing.T) {
	type testCase struct {
		i sdk.DecCoins
		o sdk.Coins
		d string
	}
	cent1, _ := sdk.ParseDecCoins("1cent")
	shrp101, _ := sdk.ParseDecCoins("1shrp,1cent")
	tcs := []testCase{
		{
			i: sdk.NewDecCoins(),
			o: sdk.NewCoins(),
			d: "0 to 0",
		},
		{
			i: sdk.NewDecCoins(sdk.NewDecCoin(DenomCent, sdk.NewInt(1))),
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
		{
			i: sdk.NewDecCoins(sdk.NewDecCoinFromDec(DenomSHRP, sdk.MustNewDecFromStr("1.1"))),
			o: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(10))),
			d: "1.1 shrp to 1shrp and 10 cent",
		},
		{
			i: cent1,
			o: sdk.NewCoins(sdk.NewCoin(DenomCent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
		{
			i: shrp101,
			o: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
	}
	for _, tc := range tcs {
		r := ShrpDecCoinsToCoins(tc.i)
		require.Equal(t, tc.o, r, tc.d)
	}
}

func TestShrpToShr(t *testing.T) {
	type testCase struct {
		i sdk.Coins
		o sdk.Coin
		d string
	}
	rate := sdk.NewDec(200)
	tcs := []testCase{
		{
			i: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(0)), sdk.NewCoin(DenomCent, sdk.NewInt(0))),
			o: sdk.NewCoin(DenomSHR, sdk.NewInt(0)),
			d: "0.0 -> 0 shr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(0))),
			o: sdk.NewCoin(DenomSHR, sdk.NewInt(200)),
			d: "1.0 -> 200 shr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(0)), sdk.NewCoin(DenomCent, sdk.NewInt(99))),
			o: sdk.NewCoin(DenomSHR, sdk.NewInt(198)),
			d: "0.99 -> 198 shr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(DenomSHRP, sdk.NewInt(1)), sdk.NewCoin(DenomCent, sdk.NewInt(99))),
			o: sdk.NewCoin(DenomSHR, sdk.NewInt(398)),
			d: "1.99 -> 398 shr",
		},
	}
	for _, tc := range tcs {
		r := ShrpToShr(tc.i, rate)
		require.Equal(t, tc.o, r, tc.d)
	}
}
