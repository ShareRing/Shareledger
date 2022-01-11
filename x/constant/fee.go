package constant

import sdk "github.com/cosmos/cosmos-sdk/types"

type DefaultLevel string

const HighFee = DefaultLevel("high")
const MediumFee = DefaultLevel("medium")
const LowFee = DefaultLevel("low")
const MinFee = DefaultLevel("min")

var DefaultFeeLevel = map[DefaultLevel]sdk.DecCoin{
	HighFee:   sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.05")),
	MediumFee: sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.03")),
	LowFee:    sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.02")),
	MinFee:    sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.01")),
}
