package posTypes

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
)

// After every block, the validator keeps a commission and accumulate the reward since a Reward Distribution.
// A Reward Distribution happens when
// * A new validator is added
// * A validator withdraws Reward
// * Validator perform Reward Distribution
// To implement this, a Validator Distribution has the following attributes:
// * LatestBlock = Latest Block that performs Reward Distribution
// * Reward Accumulation = Reward received by this validator since Latest Block
// * Commission = Commission kept by this Validator

type ValidatorDistInfo struct {
	ValidatorAddr    sdk.Address `json:"validator_addr"`    // Validator Address
	RewardAccum      types.Coin  `json:"reward_accum"`      // Block Reward accumulation since BlockHeight, excluding commission
	Commission       types.Coin  `json:"commision"`         // total commission
	WithdrawalHeight int64       `json:"withdrawal_height"` // Latest blockheight that performs reward distribution
	ValidatorReward  types.Coin  `json:"validator_reward"`  // Validator reward accumulation not yet witdrawed
}

// NewValidatorDistInfo - return new ValidatorDistInfo
func NewValidatorDistInfo(
	validatorAddress sdk.Address,
	currentHeight int64,
) ValidatorDistInfo {
	return ValidatorDistInfo{
		ValidatorAddr:    validatorAddress,
		RewardAccum:      types.NewZeroPOSCoin(),
		Commission:       types.NewZeroPOSCoin(),
		WithdrawalHeight: currentHeight,
		ValidatorReward:  types.NewZeroPOSCoin(),
	}
}

func MustMarshalValidatorDist(
	cdc *wire.Codec, vdi ValidatorDistInfo,
) []byte {
	return cdc.MustMarshalBinary(vdi)
}

func UnmarshalValidatorDist(
	cdc *wire.Codec, value []byte,
) (
	vdi ValidatorDistInfo, err error,
) {
	err = cdc.UnmarshalBinary(value, &vdi)
	if err != nil {
		return
	}

	return vdi, nil
}

func MustUnmarshalValidatorDist(
	cdc *wire.Codec, value []byte,
) ValidatorDistInfo {
	vdi, err := UnmarshalValidatorDist(cdc, value)
	if err != nil {
		panic(err)
	}

	return vdi
}

func (vdi ValidatorDistInfo) HumanReadableString() string {
	resp := ""
	resp += fmt.Sprintf("Validator Address: %x\n", vdi.ValidatorAddr)
	resp += fmt.Sprintf("Reward Accum: %s\n", vdi.RewardAccum.String())
	resp += fmt.Sprintf("Commission: %s\n", vdi.Commission.String())
	resp += fmt.Sprintf("WithdrawalHeight: %d\n", vdi.WithdrawalHeight)
	resp += fmt.Sprintf("ValidatorReward: %s\n", vdi.ValidatorReward.String())

	return resp
}
