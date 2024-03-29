package denom

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
	"math/big"
)

func CheckSupport(denom string) error {
	if _, found := supportedDenoms[denom]; !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v is not supported", denom)
	}
	return nil
}

func IsShrOrBase(dCoin sdk.DecCoin) bool {
	if dCoin.Denom == Base || dCoin.Denom == Shr {
		return true
	}
	return false
}

func CheckFeeSupportedCoin(dCoin sdk.DecCoin) error {
	if dCoin.Denom != ShrP {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "%v denomination is not supported. Only support %v", dCoin.Denom, ShrP)
	}
	return nil
}

func CheckSupportedCoins(dCoins sdk.DecCoins, coins sdk.Coins) error {
	if len(dCoins) == 0 && len(coins) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no coins")
	}
	var err error
	if len(dCoins) > 0 {
		err = dCoins.Validate()
	}
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coins string was not supported. Format should be {amount0}{denomination},...,{amountN}{denominationN}")
	}
	if len(coins) > 0 {
		err = coins.Validate()
	}
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coins string was not supported. Format should be {amount0}{denomination},...,{amountN}{denominationN}")
	}
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

func getBaseDenomUnits(usdRate sdk.Dec) map[string]sdk.Dec {
	return map[string]sdk.Dec{
		Base:    sdk.NewDec(1),
		Shr:     sdk.NewDec(ShrExponent),
		BaseUSD: usdRate.Mul(sdk.NewDec(ShrExponent)).Quo(sdk.NewDec(USDExponent)),
		ShrP:    usdRate.Mul(sdk.NewDec(ShrExponent)),
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
	baseDecCoins, err := To(dcoins, baseName, usdRate)
	if err != nil {
		return
	}
	coin = sdk.NewCoin(baseName, baseDecCoins.Amount.TruncateInt())
	if roundUp {
		coin = sdk.NewCoin(baseName, baseDecCoins.Amount.Ceil().TruncateInt())
	}
	return
}

// ToDisplayCoins convert coins to display coins which are Shr and ShrP
func ToDisplayCoins(coins sdk.Coins) sdk.DecCoins {
	shr := sdk.NewDecCoins(
		sdk.NewDecCoin(Base, coins.AmountOf(Base)),
		sdk.NewDecCoin(Shr, coins.AmountOf(Shr)),
	)
	dShr, _ := To(shr, Shr, sdk.NewDec(1))

	shrp := sdk.NewDecCoins(
		sdk.NewDecCoin(BaseUSD, coins.AmountOf(BaseUSD)),
		sdk.NewDecCoin(ShrP, coins.AmountOf(ShrP)),
	)
	dShrP, _ := To(shrp, ShrP, sdk.NewDec(1))

	return sdk.NewDecCoins(dShr, dShrP)
}

// ExponentToBase convert value to base coin with given exponent
func ExponentToBase(value sdk.Dec, exponent int) (coin sdk.Coin) {
	baseCoin := value.Quo(sdk.NewDec(int64(math.Pow10(exponent)))).Mul(sdk.NewDec(ShrExponent))
	return sdk.NewCoin(Base, baseCoin.TruncateInt())
}

// ShrCoinsToExponent only shr and nshr are allowed. others will be ignored.
func ShrCoinsToExponent(coins sdk.DecCoins, exponent int, roundUp bool) (total *big.Int, err error) {
	baseCoins, err := NormalizeToBaseCoins(coins, roundUp)
	if err != nil {
		return nil, err
	}
	totalBaseCoin := baseCoins.AmountOf(Base)
	return totalBaseCoin.Mul(sdk.NewInt(int64(math.Pow10(exponent)))).Quo(sdk.NewInt(ShrExponent)).BigInt(), nil
}

func To(coins sdk.DecCoins, dest string, usdRate sdk.Dec) (coin sdk.DecCoin, err error) {
	if err = coins.Validate(); err != nil {
		return
	}
	if err = CheckSupport(dest); err != nil {
		return
	}
	baseUnits := getBaseDenomUnits(usdRate)
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
		vd = vd.Add(c.Amount.Mul(srcUnit).Quo(destUnit))
	}
	coin = sdk.NewDecCoinFromDec(dest, vd)
	return
}

func OnlyBaseDenom(cs sdk.Coins) sdk.Coin {
	for _, c := range cs {
		if c.Denom == Base {
			return c
		}
	}
	return sdk.Coin{}
}
