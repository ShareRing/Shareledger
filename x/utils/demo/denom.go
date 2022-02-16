package denom

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

const (
	Base = "nshr"
	Shr  = "shr"

	BaseUSD = "cent"
	ShrP    = "shrp"
	Cent    = "cent"
)

var (
	USDExponent = int64(100)
	ShrExponent = int64(math.Pow(10, 9))

	OneShr          = sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(ShrExponent)))
	OneUSD          = sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)))
	OneHundredCents = sdk.NewCoins(sdk.NewCoin(Cent, sdk.NewInt(100)))
)

// ToDisplayCoins convert coins to display coins which are SHR and SHRP
func ToDisplayCoins(coins sdk.Coins) sdk.DecCoins {
	shr := sdk.NewDecCoinFromDec(Shr,
		sdk.NewDec(coins.AmountOf(Base).Int64()).QuoInt64(ShrExponent).
			Add(coins.AmountOf(Shr).ToDec()))
	shrP := sdk.NewDecCoinFromDec(ShrP,
		sdk.NewDec(coins.AmountOf(Cent).Int64()).QuoInt64(USDExponent).
			Add(coins.AmountOf(ShrP).ToDec()))

	return sdk.NewDecCoins(shr, shrP)
}

// NormalizeBaseCoins convert all dec coins into 2 base coins, Base and BaseUSD, in shareledger
func NormalizeBaseCoins(coins sdk.DecCoins) (baseCoins sdk.Coins, err error) {
	if err = coins.Validate(); err != nil {
		return
	}
	baseCoins = sdk.NewCoins(
		sdk.NewCoin(Base, coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent)).
			Add(coins.AmountOf(Base)).
			TruncateInt()),
		sdk.NewCoin(BaseUSD, coins.AmountOf(ShrP).Mul(sdk.NewDec(USDExponent)).
			Add(coins.AmountOf(Base)).
			TruncateInt()),
	)
	return
}

// NormalizeCoins convert all coins to base coin, shrp
// if there is any amount of SHRP, usdRate should be required
func NormalizeCoins(coins sdk.DecCoins, usdRate *sdk.Dec) (sdk.Coin, error) {
	shrpDec := coins.AmountOf(ShrP).
		Add(coins.AmountOf(Cent).Quo(sdk.NewDec(USDExponent)))
	uRate := sdk.NewDec(1)
	if shrpDec.GT(sdk.NewDec(0)) {
		if usdRate == nil {
			return sdk.Coin{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, fmt.Sprintf("shrp rate to %v is required", Base))
		}
		uRate = *usdRate
	}

	coin := sdk.NewCoin(Base,
		shrpDec.Mul(uRate).
			Add(coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent))).
			Add(coins.AmountOf(Base)).
			TruncateInt())
	return coin, nil
}

// ToDecShrPCoin convert all coins' types to shrp dec coin
func ToDecShrPCoin(coins sdk.DecCoins, usdRate sdk.Dec) sdk.DecCoin {
	shrpDec := coins.AmountOf(ShrP).
		Add(coins.AmountOf(Cent).Quo(sdk.NewDec(USDExponent)))

	base := sdk.NewDecCoinFromDec(Base,
		shrpDec.Mul(usdRate).
			Add(coins.AmountOf(Base)).
			Add(coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent))))

	return sdk.NewDecCoinFromDec(ShrP, shrpDec.Add(base.Amount.Quo(usdRate)))
}

// ShrpDecToCoins convert shrp dec coins to int coins which contains shrp and cent denom
func ShrpDecToCoins(dCoins sdk.DecCoins) (coin sdk.Coins) {
	shrp := dCoins.AmountOf(ShrP).Add(dCoins.AmountOf(Cent).Quo(sdk.NewDec(USDExponent)))
	return sdk.NewCoins(
		sdk.NewCoin(ShrP, shrp.TruncateInt()),
		sdk.NewCoin(Cent, shrp.Sub(shrp.TruncateDec()).MulInt(sdk.NewInt(USDExponent)).Ceil().TruncateInt()),
	)
}

// SubShrpCoins return x - y.
// x,y: Coins which can contain [ShrP, cent]
// z: result in coins
// err:
//	coins is not in valid form,
// 	sdkerrors.ErrInsufficientFunds: negative result.
func SubShrpCoins(x sdk.Coins, y sdk.Coins) (z sdk.Coins, err error) {
	if err = x.Validate(); err != nil {
		return
	}

	xI := x.AmountOf(ShrP).Mul(sdk.NewInt(USDExponent)).Add(x.AmountOf(Cent))
	yI := y.AmountOf(ShrP).Mul(sdk.NewInt(USDExponent)).Add(y.AmountOf(Cent))

	zI := xI.Sub(yI)
	if zI.IsNegative() {
		err = sdkerrors.ErrInsufficientFunds
		return
	}
	shrp := sdk.NewInt(zI.Int64() / USDExponent)
	cent := sdk.NewInt(zI.Int64() - zI.Int64()/USDExponent*USDExponent)

	z = sdk.NewCoins(
		sdk.NewCoin(ShrP, shrp),
		sdk.NewCoin(Cent, cent),
	)
	return

}
