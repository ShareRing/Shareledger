package denom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

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
