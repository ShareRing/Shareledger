package types

import (
	"math"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type AdjustmentCoins struct {
	Sub sdk.Coins
	Add sdk.Coins
}

func GetCostShrpForShr(currentShrp sdk.Coins, needShr sdk.Int, rate float64) (cost AdjustmentCoins, err error) {
	neededShrpF := float64(needShr.Int64()) / rate
	neededShrp, err := ParseShrpCoinsFloat(neededShrpF)
	if err != nil {
		return
	}
	newBalance, err := SubShrpCoins(currentShrp, neededShrp)
	if err != nil {
		return
	}
	cost = AdjustmentCoins{
		Sub: sdk.NewCoins(),
		Add: sdk.NewCoins(),
	}
	zeroI := sdk.NewInt(0)
	if v := currentShrp.AmountOf(DenomSHRP).Sub(newBalance.AmountOf(DenomSHRP)); v.GT(zeroI) {
		cost.Sub = cost.Sub.Add(sdk.NewCoin(DenomSHRP, v))
	}
	if v := currentShrp.AmountOf(DenomCent).Sub(newBalance.AmountOf(DenomCent)); !v.Equal(zeroI) {
		if v.LT(zeroI) {
			cost.Add = cost.Add.Add(sdk.NewCoin(DenomCent, v.Abs()))
		} else {
			cost.Sub = cost.Sub.Add(sdk.NewCoin(DenomCent, v))
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

	oldCents := currentCoins.AmountOf(DenomCent)
	addedCents := addedCoins.AmountOf(DenomCent)
	totalCents := oldCents.Add(addedCents)

	ac.Add = sdk.NewCoins()
	ac.Sub = sdk.NewCoins()
	// convert cent to shrp
	ac.Add = ac.Add.Add(sdk.NewCoin(DenomSHRP, sdk.NewInt(totalCents.Int64()/100)))

	newCent := sdk.NewInt(totalCents.Int64() % 100)
	if oldCents.GT(newCent) {
		ac.Sub = ac.Sub.Add(sdk.NewCoin(DenomCent, oldCents.Sub(newCent)))
	} else {
		ac.Add = ac.Add.Add(sdk.NewCoin(DenomCent, newCent.Sub(oldCents)))
	}

	return
}

// SubShrpCoins return x - y.
// x,y: Coins which can contain [SHRP, cent]
// z: result in coins
// err:
//	coins is not in valid form,
// 	sdkerrors.ErrInsufficientFunds: negative result.
func SubShrpCoins(x sdk.Coins, y sdk.Coins) (z sdk.Coins, err error) {
	if err = x.Validate(); err != nil {
		return
	}

	xI := x.AmountOf(DenomSHRP).Mul(sdk.NewInt(100)).Add(x.AmountOf(DenomCent))
	yI := y.AmountOf(DenomSHRP).Mul(sdk.NewInt(100)).Add(y.AmountOf(DenomCent))

	zI := xI.Sub(yI)
	if zI.LT(sdk.NewInt(0)) {
		err = sdkerrors.ErrInsufficientFunds
		return
	}
	shrp := sdk.NewInt(zI.Int64() / 100)
	cent := sdk.NewInt(zI.Int64() - zI.Int64()/100*100)

	z = sdk.NewCoins(
		sdk.NewCoin(DenomSHRP, shrp),
		sdk.NewCoin(DenomCent, cent),
	)
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
		cent, err = strconv.ParseInt(strNumbers[1], 10, 64)
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
		sdk.NewCoin(DenomSHRP, sdk.NewInt(shrp)),
		sdk.NewCoin(DenomCent, sdk.NewInt(cent)),
	), nil
}

// ParseShrpCoinsFloat parse float to shrp coins: shrp and cent.
// Always round up.
// E.g: 1.000000000000001 => 1.01
func ParseShrpCoinsFloat(v float64) (coins sdk.Coins, err error) {
	if v < 0 {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "Negative Coins are not accepted")
		return
	}

	v = math.Ceil(v*100) / 100
	shrp := int64(v)
	cent := int64(v*100) - shrp*100

	return sdk.NewCoins(
		sdk.NewCoin(DenomSHRP, sdk.NewInt(shrp)),
		sdk.NewCoin(DenomCent, sdk.NewInt(cent)),
	), nil
}

func ParseShrCoinsStr(s string) (coins sdk.Coins, err error) {
	v, ok := sdk.NewIntFromString(s)
	if !ok {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, s)
		return
	}
	if v.LT(sdk.NewInt(0)) {
		err = sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, s)
		return
	}
	coins = sdk.NewCoins(sdk.NewCoin(DenomSHR, v))
	return
}
