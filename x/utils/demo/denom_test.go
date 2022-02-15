package denom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

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
			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(0))),
			d: "0 to 0",
		},
		{
			i: sdk.NewDecCoins(sdk.NewDecCoin(Cent, sdk.NewInt(1))),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
		{
			i: sdk.NewDecCoins(sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("1.1"))),
			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(10))),
			d: "1.1 shrp to 1shrp and 10 cent",
		},
		{
			i: cent1,
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
		{
			i: shrp101,
			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(1))),
			d: "1 cent to 1 cent",
		},
	}
	for _, tc := range tcs {
		r := ShrpDecToCoins(tc.i)
		require.Equal(t, tc.o, r, tc.d)
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
				sdk.NewCoin(ShrP, sdk.NewInt(0)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(0)),
			),
			o: sdk.NewCoins(),
			d: "0 - 0 = 0",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(9)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(9)),
			),
			o: sdk.NewCoins(),
			d: "9.99 - 9.99 = 0",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(5)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(4)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(1)),
			),
			d: "5 - 4 = 1",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(5)),
				sdk.NewCoin(Cent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(4)),
				sdk.NewCoin(Cent, sdk.NewInt(3)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(1)),
				sdk.NewCoin(Cent, sdk.NewInt(1)),
			),
			d: "5.04 - 4.03 = 1.01",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(5)),
				sdk.NewCoin(Cent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(4)),
				sdk.NewCoin(Cent, sdk.NewInt(5)),
			),
			o: sdk.NewCoins(
				sdk.NewCoin(Cent, sdk.NewInt(99)),
			),
			d: "5 shrp 4 cent - 4 shrp 5 cent = 0 shrp 99 cent",
		},
		{
			x: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(5)),
				sdk.NewCoin(Cent, sdk.NewInt(4)),
			),
			y: sdk.NewCoins(
				sdk.NewCoin(ShrP, sdk.NewInt(5)),
				sdk.NewCoin(Cent, sdk.NewInt(6)),
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

func TestNormalizeCoins(t *testing.T) {
	type testCase struct {
		i sdk.Coins
		o sdk.Coin
		d string
	}
	rate := sdk.NewDec(200)
	tcs := []testCase{
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(0)), sdk.NewCoin(Cent, sdk.NewInt(0))),
			o: sdk.NewCoin(Base, sdk.NewInt(0)),
			d: "0.0 shrp -> 0 base",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(0))),
			o: sdk.NewCoin(Base, sdk.NewInt(200)),
			d: "1.0 shrp -> 200 base",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(0)), sdk.NewCoin(Cent, sdk.NewInt(99))),
			o: sdk.NewCoin(Base, sdk.NewInt(198)),
			d: "0.99 shrp -> 198 base",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(99))),
			o: sdk.NewCoin(Base, sdk.NewInt(398)),
			d: "1.99 shrp -> 398 base",
		},
	}
	for _, tc := range tcs {
		r, _ := NormalizeCoins(sdk.NewDecCoinsFromCoins(tc.i...), &rate)
		require.Equal(t, tc.o, r, tc.d)
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
			i: sdk.NewCoin(Base, sdk.NewInt(1)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(1))),
			d: "1 base -> 1 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(Base, sdk.NewInt(3)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(2))),
			d: "3 base -> 2 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(Base, sdk.NewInt(4)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(2))),
			d: "4 base -> 2 cent",
		},
	}
	for _, tc := range tcs {
		dr := ToDecShrPCoin(sdk.NewDecCoinsFromCoins(tc.i), rate)
		r := ShrpDecToCoins(sdk.NewDecCoins(dr))
		require.Equal(t, tc.o, r, tc.d)
	}
}

func TestToDisplayCoins(t *testing.T) {
	type testCase struct {
		i sdk.Coins
		o sdk.DecCoins
		d string
	}
	tcs := []testCase{
		{
			i: sdk.NewCoins(),
			o: sdk.NewDecCoins(),
			d: "0 -> 0",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Shr, sdk.NewInt(1)), sdk.NewCoin(ShrP, sdk.NewInt(1))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.NewDec(1)), sdk.NewDecCoinFromDec(ShrP, sdk.NewDec(1))),
			d: "1shr 1shrp -> 1shr 1shrp",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(10)), sdk.NewCoin(Cent, sdk.NewInt(10))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("0.0000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("0.1"))),
			d: "10base 1cent -> 0.0000001shr 0.01shrp",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(210000010)), sdk.NewCoin(Cent, sdk.NewInt(131))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("2.1000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("1.31"))),
			d: "210000010base 131cent -> 2.1000001shr 1.31shrp",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Shr, sdk.NewInt(1)), sdk.NewCoin(Base, sdk.NewInt(210000010)), sdk.NewCoin(ShrP, sdk.NewInt(2)), sdk.NewCoin(Cent, sdk.NewInt(131))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("3.1000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("3.31"))),
			d: "1shr 210000010base 2shrp 131cent -> 3.1000001shr 3.31shrp",
		},
	}

	for _, tc := range tcs {
		r := ToDisplayCoins(tc.i)
		require.Equal(t, tc.o, r, tc.d)
	}
}
