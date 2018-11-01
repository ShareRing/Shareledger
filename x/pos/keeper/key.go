package keeper

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
)

// TODO remove some of these prefixes once have working multistore

//nolint
var (
	// Keys for store prefixes
	ParamKey                         = []byte{0x00} // key for parameters relating to staking
	PoolKey                          = []byte{0x01} // key for the staking pools
	ValidatorsKey                    = []byte{0x02} // prefix for each key to a validator
	ValidatorsByConsAddrKey          = []byte{0x03} // prefix for each key to a validator index, by pubkey
	ValidatorsBondedIndexKey         = []byte{0x04} // prefix for each key to a validator index, for bonded validators
	ValidatorsByPowerIndexKey        = []byte{0x05} // prefix for each key to a validator index, sorted by power
	IntraTxCounterKey                = []byte{0x06} // key for intra-block tx index
	DelegationKey                    = []byte{0x07} // key for a delegation
	UnbondingDelegationKey           = []byte{0x08} // key for an unbonding-delegation
	UnbondingDelegationByValIndexKey = []byte{0x09} // prefix for each key for an unbonding-delegation, by validator operator
	RedelegationKey                  = []byte{0x0A} // key for a redelegation
	RedelegationByValSrcIndexKey     = []byte{0x0B} // prefix for each key for an redelegation, by source validator operator
	RedelegationByValDstIndexKey     = []byte{0x0C} // prefix for each key for an redelegation, by destination validator operator
	ValidatorDistKey                 = []byte{0x0D} // prefix for each key for validator distribution information
)

// gets the key for the validator with address
// VALUE: stake/types.Validator
func GetValidatorKey(operatorAddr sdk.Address) []byte {
	return append(ValidatorsKey, operatorAddr.Bytes()...)
}

func GetValidatorDistKey(operatorAddr sdk.Address) []byte {
	return append(ValidatorDistKey, operatorAddr.Bytes()...)
}

//______________________________________________________________________________

// gets the key for delegator bond with validator
// VALUE: stake/types.Delegation
func GetDelegationKey(delAddr sdk.Address, valAddr sdk.Address) []byte {
	return append(GetDelegationsKey(delAddr), valAddr.Bytes()...)
}

// gets the prefix for a delegator for all validators
func GetDelegationsKey(delAddr sdk.Address) []byte {
	return append(DelegationKey, delAddr.Bytes()...)
}
