package constant

import sdk "github.com/cosmos/cosmos-sdk/types"

type DefaultLevel string

const HighFee = DefaultLevel("HIGHT")
const MediumFee = DefaultLevel("MEDIUM")
const LowFee = DefaultLevel("LOW")
const MinFee = DefaultLevel("MIN")

var DefaultFeeLevel = map[DefaultLevel]sdk.DecCoin{
	HighFee:   sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.05")),
	MediumFee: sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.03")),
	LowFee:    sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.02")),
	MinFee:    sdk.NewDecCoinFromDec("shrp", sdk.MustNewDecFromStr("0.01")),
}
