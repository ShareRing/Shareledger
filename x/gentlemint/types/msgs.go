package types

import (
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type AdjustmentCoins struct {
	Sub sdk.Coins
	Add sdk.Coins
}

//func PShrToDecShrp(pshr sdk.Coin, rate sdk.Dec) sdk.DecCoin {
//	shrp := sdk.NewDec(pshr.Amount.Int64()).Quo(rate)
//	return sdk.NewDecCoinFromDec(denom.ShrP, shrp)
//}

//func PShrToShrp(pshr sdk.Coin, rate sdk.Dec) (coin sdk.Coins) {
//	shrp := PShrToDecShrp(pshr, rate)
//	return denom.ShrpDecToCoins(shrp.Amount)
//}

//func DecCoinsToPShr(coins sdk.DecCoins, rate sdk.Dec) (coin sdk.Coin) {
//	shrp := ShrpDecCoinsToCoins(coins)
//	mixedCoins := sdk.NewCoins(sdk.NewCoin(denom.PShr, coins.AmountOf(denom.PShr).TruncateInt()))
//	mixedCoins = mixedCoins.Add(shrp...)
//	return CoinsToPShr(mixedCoins, rate)
//}

//func CoinsToPShr(coins sdk.Coins, rate sdk.Dec) (coin sdk.Coin) {
//	shrpDec := sdk.NewDec(coins.AmountOf(denom.ShrP).Int64()).Add(sdk.NewDec(coins.AmountOf(denom.Cent).Int64()).Quo(sdk.NewDec(100)))
//	coin = sdk.NewCoin(denom.PShr, shrpDec.Mul(rate).TruncateInt().Add(coins.AmountOf(denom.PShr)))
//	return coin
//}

//func ShrpDecCoinsToCoins(shrp sdk.DecCoins) (coin sdk.Coins) {
//	shrpD := shrp.AmountOf(denom.ShrP).Add(shrp.AmountOf(denom.Cent).Quo(sdk.NewDec(100)))
//	return ShrpDecToCoins(shrpD)
//}

//func ShrpDecToCoins(shrp sdk.Dec) (coin sdk.Coins) {
//	return sdk.NewCoins(
//		sdk.NewCoin(denom.ShrP, shrp.TruncateInt()),
//		sdk.NewCoin(denom.Cent, shrp.Sub(shrp.TruncateDec()).MulInt(sdk.NewInt(100)).Ceil().TruncateInt()),
//	)
//}

func GetCostShrpForPShr(currentShrp sdk.Coins, needShr sdk.Int, rate sdk.Dec) (cost AdjustmentCoins, err error) {
	neededDecShrp := denom.ToDecShrPCoin(sdk.NewDecCoins(sdk.NewDecCoin(denom.PShr, needShr)), rate)
	neededShrp := denom.ShrpDecToCoins(sdk.NewDecCoins(neededDecShrp))
	if err != nil {
		return
	}
	newBalance, err := denom.SubShrpCoins(currentShrp, neededShrp)
	if err != nil {
		return
	}
	cost = AdjustmentCoins{
		Sub: sdk.NewCoins(),
		Add: sdk.NewCoins(),
	}
	zeroI := sdk.NewInt(0)
	if v := currentShrp.AmountOf(denom.ShrP).Sub(newBalance.AmountOf(denom.ShrP)); v.GT(zeroI) {
		cost.Sub = cost.Sub.Add(sdk.NewCoin(denom.ShrP, v))
	}
	if v := currentShrp.AmountOf(denom.Cent).Sub(newBalance.AmountOf(denom.Cent)); !v.Equal(zeroI) {
		if v.LT(zeroI) {
			cost.Add = cost.Add.Add(sdk.NewCoin(denom.Cent, v.Abs()))
		} else {
			cost.Sub = cost.Sub.Add(sdk.NewCoin(denom.Cent, v))
		}
	}
	return
}

func AddShrpCoins(currentCoins sdk.Coins, addedCoins sdk.Coins) (ac AdjustmentCoins, err error) {
	if err = currentCoins.Validate(); err != nil {
		return
	}
	if err = addedCoins.Validate(); err != nil {
		return
	}

	oldCents := currentCoins.AmountOf(denom.Cent)
	addedCents := addedCoins.AmountOf(denom.Cent)
	totalCents := oldCents.Add(addedCents)

	ac.Add = sdk.NewCoins()
	ac.Sub = sdk.NewCoins()
	// convert cent to shrp
	ac.Add = ac.Add.Add(sdk.NewCoin(denom.ShrP, addedCoins.AmountOf(denom.ShrP)))
	ac.Add = ac.Add.Add(sdk.NewCoin(denom.ShrP, sdk.NewInt(totalCents.Int64()/100)))

	newCent := sdk.NewInt(totalCents.Int64() % 100)
	if oldCents.GT(newCent) {
		ac.Sub = ac.Sub.Add(sdk.NewCoin(denom.Cent, oldCents.Sub(newCent)))
	} else {
		ac.Add = ac.Add.Add(sdk.NewCoin(denom.Cent, newCent.Sub(oldCents)))
	}

	return
}

// ParseShrpCoinsStr return shrp and cent coins.
// only get 2 decimals to cent without rouding.
func ParseShrpCoinsStr(s string) (coins sdk.Coins, err error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}
	if f < 0 {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
		return
	}

	strNumbers := strings.Split(s, ".")
	var shrp, cent int64
	shrp, err = strconv.ParseInt(strNumbers[0], 10, 64)
	if err != nil {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "parsing got %+v", err)
		return
	}
	if len(strNumbers) > 1 {
		centStr := strNumbers[1]
		if len(centStr) == 1 {
			centStr = strNumbers[1] + "0" // cover case x.1 => x.10
		}
		cent, err = strconv.ParseInt(centStr, 10, 64)
		if err != nil {
			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "parsing got %+v", err)
			return
		}
		if cent > 99 {
			err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cent value, %v, should be less than 100", cent)
			return
		}
	}
	return sdk.NewCoins(
		sdk.NewCoin(denom.ShrP, sdk.NewInt(shrp)),
		sdk.NewCoin(denom.Cent, sdk.NewInt(cent)),
	), nil
}

func ParsePShrCoinsStr(s string) (coins sdk.Coins, err error) {
	v, ok := sdk.NewIntFromString(s)
	if !ok {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, s)
		return
	}
	if v.LT(sdk.NewInt(0)) {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, s)
		return
	}
	coins = sdk.NewCoins(sdk.NewCoin(denom.PShr, v))
	return
}
