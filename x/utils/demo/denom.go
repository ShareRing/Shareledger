package denom

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

func CheckSupport(denom string) error {
	if _, found := supportedDenoms[denom]; !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supported", denom)
	}
	return nil
}

func CheckSupportedCoins(dCoins sdk.DecCoins, coins sdk.Coins) error {
	for _, c := range dCoins {
		if _, found := supportedDenoms[c.Denom]; !found {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supported", c)
		}
	}
	for _, c := range coins {
		if _, found := supportedDenoms[c.Denom]; !found {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supported", c)
		}
	}
	return nil
}

const (
	Base = "nshr"
	Shr  = "shr"

	ShrP    = "shrp"
	BaseUSD = "cent"
)

var (
	USDExponent = int64(100)
	ShrExponent = int64(math.Pow(10, 9))

	OneShr          = sdk.NewCoins(sdk.NewCoin(Base, sdk.NewInt(ShrExponent)))
	OneUSD          = sdk.NewCoins(sdk.NewCoin(ShrP, sdk.NewInt(1)))
	OneHundredCents = sdk.NewCoins(sdk.NewCoin(BaseUSD, sdk.NewInt(100)))
)

var supportedDenoms = map[string]struct{}{
	Base:    {},
	BaseUSD: {},
	Shr:     {},
	ShrP:    {},
}

func buildDenomUnits(baseName string, usdRate sdk.Dec) (map[string]sdk.Dec, error) {
	switch baseName {
	case Base: // Base
		return map[string]sdk.Dec{
			Base:    sdk.NewDec(1),
			Shr:     sdk.NewDec(ShrExponent),
			BaseUSD: usdRate.Mul(sdk.NewDec(ShrExponent)).Quo(sdk.NewDec(USDExponent)),
			ShrP:    usdRate.Mul(sdk.NewDec(ShrExponent)),
		}, nil
	case BaseUSD: // BaseUSD
		return map[string]sdk.Dec{
			BaseUSD: sdk.NewDec(1),
			ShrP:    sdk.NewDec(USDExponent),
			Base:    sdk.NewDec(USDExponent).Quo(sdk.NewDec(ShrExponent).Mul(usdRate)),
			Shr:     sdk.NewDec(USDExponent).Quo(usdRate),
		}, nil
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "provided %v base denom is not supported. Supported Denoms: %v, %v", baseName, Base, BaseUSD)
	}
}

func NormalizeToBaseCoins(dcoins sdk.DecCoins, roundUp bool) (sdk.Coins, error) {
	if err := dcoins.Validate(); err != nil {
		return nil, err
	}
	if err := CheckSupportedCoins(dcoins, nil); err != nil {
		return nil, err
	}
	// there is no need to set usd rate, since we only need to convert BaseUSD and Base
	base, err := NormalizeToBaseCoin(Base, sdk.NewDecCoins(
		sdk.NewDecCoinFromDec(Shr, dcoins.AmountOf(Shr)),
		sdk.NewDecCoinFromDec(Base, dcoins.AmountOf(Base))), sdk.NewDec(1), roundUp)
	if err != nil {
		return nil, err
	}
	baseUSD, err := NormalizeToBaseCoin(BaseUSD, sdk.NewDecCoins(
		sdk.NewDecCoinFromDec(ShrP, dcoins.AmountOf(ShrP)),
		sdk.NewDecCoinFromDec(BaseUSD, dcoins.AmountOf(BaseUSD))), sdk.NewDec(1), roundUp)
	if err != nil {
		return nil, err
	}

	return sdk.NewCoins(
		base, baseUSD,
	), err
}

func NormalizeToBaseCoin(baseName string, dcoins sdk.DecCoins, usdRate sdk.Dec, roundUp bool) (coin sdk.Coin, err error) {
	coin = sdk.NewCoin(baseName, sdk.NewInt(0))
	if err = dcoins.Validate(); err != nil {
		return
	}
	baseUnits, err := buildDenomUnits(baseName, usdRate)
	if err != nil {
		return
	}
	destUnit, found := baseUnits[baseName]
	if !found {
		err = fmt.Errorf("%v not found in base units map", baseName)
		return
	}
	dValue := sdk.NewDec(0)
	for _, c := range dcoins {
		srcUnit, found := baseUnits[c.Denom]
		if !found {
			err = fmt.Errorf("%v not found in base units map", srcUnit)
			return
		}
		dValue = dValue.Add(c.Amount.Mul(srcUnit).Quo(destUnit))
	}
	iv := dValue.TruncateInt()
	if roundUp {
		iv = dValue.Ceil().TruncateInt()
	}
	coin = sdk.NewCoin(baseName, iv)
	return
}

// ToDisplayCoins convert coins to display coins which are Shr and ShrP
func ToDisplayCoins(coins sdk.Coins) sdk.DecCoins {
	shr := sdk.NewDecCoinFromDec(Shr,
		sdk.NewDec(coins.AmountOf(Base).Int64()).QuoInt64(ShrExponent).
			Add(coins.AmountOf(Shr).ToDec()))

	shrP := sdk.NewDecCoinFromDec(ShrP,
		sdk.NewDec(coins.AmountOf(BaseUSD).Int64()).QuoInt64(USDExponent).
			Add(coins.AmountOf(ShrP).ToDec()))

	return sdk.NewDecCoins(shr, shrP)
}

func To(coins sdk.DecCoins, dest string, usdRate sdk.Dec) (coin sdk.DecCoin, err error) {
	if err = coins.Validate(); err != nil {
		return
	}
	if err = CheckSupport(dest); err != nil {
		return
	}
	baseUnits, err := buildDenomUnits(Base, usdRate)
	if err != nil {
		return
	}

	destUnit, found := baseUnits[dest]
	if !found {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supporter", dest)
		return
	}

	vd := sdk.NewDec(0)
	for _, c := range coins {
		srcUnit, found := baseUnits[c.Denom]
		if !found {
			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supporter", c.Denom)
			return
		}
		vd = vd.Add(c.Amount.Quo(srcUnit).Mul(destUnit))
	}
	coin = sdk.NewDecCoinFromDec(dest, vd)
	return
}
