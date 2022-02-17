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

//// ParseShrpCoinsStr return shrp and cent coins.
//// only get 2 decimals to cent without rouding.
//func ParseShrpCoinsStr(s string) (coins sdk.Coins, err error) {
//	f, err := strconv.ParseFloat(s, 64)
//	if err != nil {
//		return
//	}
//	if f < 0 {
//		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
//		return
//	}
//
//	strNumbers := strings.Split(s, ".")
//	var shrp, cent int64
//	shrp, err = strconv.ParseInt(strNumbers[0], 10, 64)
//	if err != nil {
//		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "parsing got %+v", err)
//		return
//	}
//	if len(strNumbers) > 1 {
//		centStr := strNumbers[1]
//		if len(centStr) == 1 {
//			centStr = strNumbers[1] + "0" // cover case x.1 => x.10
//		}
//		cent, err = strconv.ParseInt(centStr, 10, 64)
//		if err != nil {
//			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "parsing got %+v", err)
//			return
//		}
//		if cent > 99 {
//			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cent value, %v, should be less than 100", cent)
//			return
//		}
//	}
//	return sdk.NewCoins(
//		sdk.NewCoin(ShrP, sdk.NewInt(shrp)),
//		sdk.NewCoin(BaseUSD, sdk.NewInt(cent)),
//	), nil
//}

//// NormalizeToBaseCoins convert all dec coins into 2 base coins, Base and BaseUSD, in shareledger
//func NormalizeToBaseCoins(coins sdk.DecCoins) (baseCoins sdk.Coins, err error) {
//	if err = coins.Validate(); err != nil {
//		return
//	}
//	baseCoins = sdk.NewCoins(
//		sdk.NewCoin(Base, coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent)).
//			Add(coins.AmountOf(Base)).
//			TruncateInt()),
//		sdk.NewCoin(BaseUSD, coins.AmountOf(ShrP).Mul(sdk.NewDec(USDExponent)).
//			Add(coins.AmountOf(Base)).
//			TruncateInt()),
//	)
//	return
//}

//// NormalizeToBaseUSDCoin convert all coins to BaseUSD coin
//// if there is any amount of Shr, Base, rateUSDSHR should be required
//func NormalizeToBaseUSDCoin(coins sdk.DecCoins, rateUSDSHR *sdk.Dec) (baseUSD sdk.Coin, err error) {
//	if err = coins.Validate(); err != nil {
//		return
//	}
//
//	shrD := sdk.NewDecCoinFromDec(Shr, coins.AmountOf(Base).Quo(sdk.NewDec(ShrExponent)).Add(coins.AmountOf(Shr)))
//	shrP := sdk.NewDecCoinFromDec(ShrP, coins.AmountOf(ShrP).Add(coins.AmountOf(BaseUSD).Quo(sdk.NewDec(USDExponent))))
//
//
//	if !shrD.IsZero() {
//		if rateUSDSHR == nil || rateUSDSHR.LTE(sdk.NewDec(0)) {
//			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "usd rate is required and should not less than or equal 0")
//			return
//		}
//		//convert to Base rate
//		uBaseRate = *rateUSDSHR
//	}
//	shrD = shrD.Add(shrP.Mu)
//
//	baseUSD = sdk.NewCoin(BaseUSD,
//			coins.AmountOf(BaseUSD).
//			Add(coins.AmountOf(ShrP).Mul(sdk.NewDec(USDExponent))).
//		)
//	return
//}

//// NormalizeToBaseCoin convert all coins to Base coin
//// if there is any amount of ShrP, BaseUSD, rateUSDSHR should be required
//func NormalizeToBaseCoin(coins sdk.DecCoins, rateUSDSHR *sdk.Dec) (sdk.Coin, error) {
//	if err := coins.Validate(); err != nil {
//		return sdk.Coin{}, err
//	}
//
//	shrpDec := coins.AmountOf(ShrP).
//		Add(coins.AmountOf(BaseUSD).Quo(sdk.NewDec(USDExponent)))
//	uBaseRate := sdk.NewDec(1)
//
//	if shrpDec.GT(sdk.NewDec(0)) {
//		if rateUSDSHR == nil || rateUSDSHR.LTE(sdk.NewDec(0)) {
//			return sdk.Coin{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, fmt.Sprintf("%v rate to %v is required", ShrP, Shr))
//		}
//		uBaseRate = rateUSDSHR.Mul(sdk.NewDec(ShrExponent))
//	}
//
//	coin := sdk.NewCoin(Base,
//		shrpDec.Mul(uBaseRate).
//			Add(coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent))).
//			Add(coins.AmountOf(Base)).
//			TruncateInt())
//	return coin, nil
//}

//// ToDecShrPCoin convert all coins' types to ShrP dec coin
//func ToDecShrPCoin(coins sdk.DecCoins, usdRate sdk.Dec) sdk.DecCoin {
//	shrpDec := coins.AmountOf(ShrP).
//		Add(coins.AmountOf(BaseUSD).Quo(sdk.NewDec(USDExponent)))
//
//	base := sdk.NewDecCoinFromDec(Base,
//		shrpDec.Mul(usdRate).
//			Add(coins.AmountOf(Base)).
//			Add(coins.AmountOf(Shr).Mul(sdk.NewDec(ShrExponent))))
//
//	return sdk.NewDecCoinFromDec(ShrP, shrpDec.Add(base.Amount.Quo(usdRate)))
//}
//
//// ShrpDecToCoins convert shrp dec coins to int coins which contains shrp and cent denom
//func ShrpDecToCoins(dCoins sdk.DecCoins) (coin sdk.Coins) {
//	shrp := dCoins.AmountOf(ShrP).Add(dCoins.AmountOf(BaseUSD).Quo(sdk.NewDec(USDExponent)))
//	return sdk.NewCoins(
//		sdk.NewCoin(ShrP, shrp.TruncateInt()),
//		sdk.NewCoin(BaseUSD, shrp.Sub(shrp.TruncateDec()).MulInt(sdk.NewInt(USDExponent)).Ceil().TruncateInt()),
//	)
//}

//// SubShrpCoins return x - y.
//// x,y: Coins which can contain [ShrP, cent]
//// z: result in coins
//// err:
////	coins is not in valid form,
//// 	sdkerrors.ErrInsufficientFunds: negative result.
//func SubShrpCoins(x sdk.Coins, y sdk.Coins) (z sdk.Coins, err error) {
//	if err = x.Validate(); err != nil {
//		return
//	}
//
//	xI := x.AmountOf(ShrP).Mul(sdk.NewInt(USDExponent)).Add(x.AmountOf(BaseUSD))
//	yI := y.AmountOf(ShrP).Mul(sdk.NewInt(USDExponent)).Add(y.AmountOf(BaseUSD))
//
//	zI := xI.Sub(yI)
//	if zI.IsNegative() {
//		err = sdkerrors.ErrInsufficientFunds
//		return
//	}
//	shrp := sdk.NewInt(zI.Int64() / USDExponent)
//	cent := sdk.NewInt(zI.Int64() - zI.Int64()/USDExponent*USDExponent)
//
//	z = sdk.NewCoins(
//		sdk.NewCoin(ShrP, shrp),
//		sdk.NewCoin(BaseUSD, cent),
//	)
//	return
//
//}
