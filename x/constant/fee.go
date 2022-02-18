package constant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
)

type DefaultLevel string

const HighFee = DefaultLevel("high")
const MediumFee = DefaultLevel("medium")
const LowFee = DefaultLevel("low")
const MinFee = DefaultLevel("min")
const NoFee = DefaultLevel("zero")

var DefaultFeeLevel = map[DefaultLevel]sdk.DecCoin{
	HighFee:   sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.05")),
	MediumFee: sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.03")),
	LowFee:    sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.02")),
	MinFee:    sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.01")),
	NoFee:     sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0")),
}
