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
			o: sdk.NewCoin(PShr, sdk.NewInt(0)),
			d: "0.0 shrp -> 0 pshr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(0))),
			o: sdk.NewCoin(PShr, sdk.NewInt(200)),
			d: "1.0 shrp -> 200 pshr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(0)), sdk.NewCoin(Cent, sdk.NewInt(99))),
			o: sdk.NewCoin(PShr, sdk.NewInt(198)),
			d: "0.99 shrp -> 198 pshr",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(Cent, sdk.NewInt(99))),
			o: sdk.NewCoin(PShr, sdk.NewInt(398)),
			d: "1.99 shrp -> 398 pshr",
		},
	}
	for _, tc := range tcs {
		r := NormalizeCoins(sdk.NewDecCoinsFromCoins(tc.i...), rate)
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
			i: sdk.NewCoin(PShr, sdk.NewInt(1)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(1))),
			d: "1 pshr -> 1 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(PShr, sdk.NewInt(3)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(2))),
			d: "3 pshr -> 2 cent (should round up when not even)",
		},
		{
			i: sdk.NewCoin(PShr, sdk.NewInt(4)),
			o: sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(2))),
			d: "4 pshr -> 2 cent",
		},
	}
	for _, tc := range tcs {
		dr := ToDecShrPCoin(sdk.NewDecCoinsFromCoins(tc.i), rate)
		r := ShrpDecToCoins(sdk.NewDecCoins(dr))
		require.Equal(t, tc.o, r, tc.d)
	}
}
