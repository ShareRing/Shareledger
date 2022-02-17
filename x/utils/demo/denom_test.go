package denom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestShrpDecCoinsToCoins(t *testing.T) {
//	type testCase struct {
//		i sdk.DecCoins
//		o sdk.Coins
//		d string
//	}
//	cent1, _ := sdk.ParseDecCoins("1cent")
//	shrp101, _ := sdk.ParseDecCoins("1shrp,1cent")
//	tcs := []testCase{
//		{
//			i: sdk.NewDecCoins(),
//			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(0))),
//			d: "0 to 0",
//		},
//		{
//			i: sdk.NewDecCoins(sdk.NewDecCoin(BaseUSD, sdk.NewInt(1))),
//			o: sdk.NewCoins(sdk.NewCoin(BaseUSD, sdk.NewInt(1))),
//			d: "1 cent to 1 cent",
//		},
//		{
//			i: sdk.NewDecCoins(sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("1.1"))),
//			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(BaseUSD, sdk.NewInt(10))),
//			d: "1.1 shrp to 1shrp and 10 cent",
//		},
//		{
//			i: cent1,
//			o: sdk.NewCoins(sdk.NewCoin(BaseUSD, sdk.NewInt(1))),
//			d: "1 cent to 1 cent",
//		},
//		{
//			i: shrp101,
//			o: sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(BaseUSD, sdk.NewInt(1))),
//			d: "1 cent to 1 cent",
//		},
//	}
//	for _, tc := range tcs {
//		r := ShrpDecToCoins(tc.i)
//		require.Equal(t, tc.o, r, tc.d)
//	}
//}

//func TestSubShrpCoins(t *testing.T) {
//	// o, e := SubShrpCoins(x, y)
//	type testCase struct {
//		x sdk.Coins
//		y sdk.Coins
//		o sdk.Coins
//		e error
//		d string
//	}
//	testCases := []testCase{
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(0)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(0)),
//			),
//			o: sdk.NewCoins(),
//			d: "0 - 0 = 0",
//		},
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(9)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(9)),
//			),
//			o: sdk.NewCoins(),
//			d: "9.99 - 9.99 = 0",
//		},
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(5)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(4)),
//			),
//			o: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(1)),
//			),
//			d: "5 - 4 = 1",
//		},
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(5)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(4)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(4)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(3)),
//			),
//			o: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(1)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(1)),
//			),
//			d: "5.04 - 4.03 = 1.01",
//		},
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(5)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(4)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(4)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(5)),
//			),
//			o: sdk.NewCoins(
//				sdk.NewCoin(BaseUSD, sdk.NewInt(99)),
//			),
//			d: "5 shrp 4 cent - 4 shrp 5 cent = 0 shrp 99 cent",
//		},
//		{
//			x: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(5)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(4)),
//			),
//			y: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(5)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(6)),
//			),
//			e: sdkerrors.ErrInsufficientFunds,
//			d: "5 shrp 4 cent - 5 shrp 6 cent = error",
//		},
//	}
//
//	for _, tc := range testCases {
//		v, err := SubShrpCoins(tc.x, tc.y)
//		require.ElementsMatch(t, tc.o, v, tc.d)
//		require.Equal(t, tc.e, err, tc.d)
//	}
//}

func TestNormalizeToBaseCoin(t *testing.T) {

}

func TestNormalizeCoins(t *testing.T) {
	type testCase struct {
		i  sdk.DecCoins
		ib string
		ic bool
		o  sdk.Coin
		d  string
	}
	rate := sdk.NewDec(200)
	tcs := []testCase{
		{
			i:  sdk.NewDecCoinsFromCoins(sdk.NewCoin(ShrP, sdk.NewInt(0)), sdk.NewCoin(BaseUSD, sdk.NewInt(0))),
			ib: Base,
			o:  sdk.NewCoin(Base, sdk.NewInt(0*ShrExponent)),
			d:  "0.0 shrp -> 0 ushr",
		},
		{
			i:  sdk.NewDecCoinsFromCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)), sdk.NewCoin(BaseUSD, sdk.NewInt(99))),
			ib: Base,
			o:  sdk.NewCoin(Base, sdk.NewInt(398*ShrExponent)),
			d:  "1.99 shrp -> 398*ShrExponent ushr",
		},
		{
			i:  sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("2.99"))),
			ib: Base,
			o:  sdk.NewCoin(Base, sdk.NewInt(2990000000)),
			d:  "2.99 shr -> 2.99 * ShrExponent ushr",
		},
		{
			i: sdk.NewDecCoins(
				sdk.NewDecCoin(ShrP, sdk.NewInt(1)),
				sdk.NewDecCoin(BaseUSD, sdk.NewInt(99)),
				sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("2.99")),
			),
			ib: Base,
			o:  sdk.NewCoin(Base, sdk.NewInt(400990000000)),
			d:  "1.99 shrp 2.99 shr -> (2.99 + 398) * ShrExponent ushr",
		},
		{
			i: sdk.NewDecCoins(
				sdk.NewDecCoin(ShrP, sdk.NewInt(1)),
				sdk.NewDecCoin(BaseUSD, sdk.NewInt(99)),
				sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("2.99")),
			),
			ib: BaseUSD,
			o:  sdk.NewCoin(BaseUSD, sdk.NewInt(200)),
			d:  "1.99 shrp 2.99 shr -> 200 cent (from 200.495) - round down by default",
		},
	}
	for _, tc := range tcs {
		r, _ := NormalizeToBaseCoin(tc.ib, tc.i, rate, tc.ic)
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
			i: sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(10)), sdk.NewCoin(BaseUSD, sdk.NewInt(10))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("0.00000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("0.1"))),
			d: "10base 1cent -> 0.00000001shr 0.01shrp",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(2100000010)), sdk.NewCoin(BaseUSD, sdk.NewInt(131))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("2.10000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("1.31"))),
			d: "2100000010base 131cent -> 2.10000001shr 1.31shrp",
		},
		{
			i: sdk.NewCoins(sdk.NewCoin(Shr, sdk.NewInt(1)), sdk.NewCoin(Base, sdk.NewInt(2100000010)), sdk.NewCoin(ShrP, sdk.NewInt(2)), sdk.NewCoin(BaseUSD, sdk.NewInt(131))),
			o: sdk.NewDecCoins(sdk.NewDecCoinFromDec(Shr, sdk.MustNewDecFromStr("3.10000001")), sdk.NewDecCoinFromDec(ShrP, sdk.MustNewDecFromStr("3.31"))),
			d: "1shr 2100000010base 2shrp 131cent -> 3.10000001shr 3.31shrp",
		},
	}

	for _, tc := range tcs {
		r := ToDisplayCoins(tc.i)
		require.Equal(t, tc.o, r, tc.d)
	}
}

//func TestParseShrpCoinsStr(t *testing.T) {
//	type testCase struct {
//		i  string
//		o  sdk.Coins
//		oe error
//		d  string
//	}
//	testCases := []testCase{
//		{
//			i: "0",
//			o: sdk.NewCoins(),
//			d: "0 shrp",
//		},
//		{
//			i: "0.01",
//			o: sdk.NewCoins(sdk.NewCoin(BaseUSD, sdk.NewInt(1))),
//			d: "0.01 shrp -> 1 cent",
//		},
//		{
//			i: "1.1",
//			o: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(1)),
//				sdk.NewCoin(BaseUSD, sdk.NewInt(10)),
//			),
//			d: "1.1 shrp -> 1shrp 10 cent",
//		},
//		{
//			i: "0.01",
//			o: sdk.NewCoins(
//				sdk.NewCoin(BaseUSD, sdk.NewInt(1)),
//			),
//			d: "0.01 shrp -> 1 cent",
//		},
//		{
//			i: "1",
//			o: sdk.NewCoins(
//				sdk.NewCoin(ShrP, sdk.NewInt(1)),
//			),
//			d: "1 -> 1 shrp",
//		},
//		{
//			i:  "-1",
//			oe: sdkerrors.ErrInvalidCoins,
//			d:  "negative -> err",
//		},
//		{
//			i:  "1.100",
//			oe: sdkerrors.ErrInvalidCoins,
//			d:  "100 cent -> err",
//		},
//	}
//	for i, tc := range testCases {
//		r, err := ParseShrpCoinsStr(tc.i)
//		require.Equal(t, tc.o, r, fmt.Sprintf("%s. test index %v", tc.d, i))
//		if tc.oe != nil {
//			require.NotNil(t, err, tc.d)
//			require.True(t, sdkerrors.IsOf(err, tc.oe), tc.d)
//		}
//	}
//}
