package keeper

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
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

// gets the prefix keyspace for the indexes of unbonding delegations for a validator
func GetUBDsByValIndexKey(valAddr sdk.Address) []byte {
	return append(UnbondingDelegationByValIndexKey, valAddr.Bytes()...)
}

// gets the prefix for all unbonding delegations from a delegator
func GetUBDsKey(delAddr sdk.Address) []byte {
	return append(UnbondingDelegationKey, delAddr.Bytes()...)
}

// gets the key for an unbonding delegation by delegator and validator addr
// VALUE: stake/types.UnbondingDelegation
func GetUBDKey(delAddr sdk.Address, valAddr sdk.Address) []byte {
	return append(
		GetUBDsKey(delAddr.Bytes()),
		valAddr.Bytes()...)
}

// gets the index-key for an unbonding delegation, stored by validator-index
// VALUE: none (key rearrangement used)
func GetUBDByValIndexKey(delAddr sdk.Address, valAddr sdk.Address) []byte {
	return append(GetUBDsByValIndexKey(valAddr), delAddr.Bytes()...)
}

// gets the prefix keyspace for redelegations from a delegator
func GetREDsKey(delAddr sdk.Address) []byte {
	return append(RedelegationKey, delAddr.Bytes()...)
}

// gets the key for a redelegation
// VALUE: stake/types.RedelegationKey
func GetREDKey(delAddr sdk.Address, valSrcAddr, valDstAddr sdk.Address) []byte {
	return append(append(
		GetREDsKey(delAddr.Bytes()),
		valSrcAddr.Bytes()...),
		valDstAddr.Bytes()...)
}

// gets the index-key for a redelegation, stored by source-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValSrcIndexKey(delAddr sdk.Address, valSrcAddr, valDstAddr sdk.Address) []byte {
	return append(append(
		GetREDsFromValSrcIndexKey(valSrcAddr),
		delAddr.Bytes()...),
		valDstAddr.Bytes()...)
}

// gets the index-key for a redelegation, stored by destination-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValDstIndexKey(delAddr sdk.Address, valSrcAddr, valDstAddr sdk.Address) []byte {
	return append(append(
		GetREDsToValDstIndexKey(valDstAddr),
		delAddr.Bytes()...),
		valSrcAddr.Bytes()...)
}

// rearranges the ValSrcIndexKey to get the REDKey
func GetREDKeyFromValSrcIndexKey(IndexKey []byte) []byte {
	addrs := IndexKey[1:] // remove prefix bytes
	if len(addrs) != 3*types.ADDRESSLENGTH {
		panic("unexpected key length")
	}
	valSrcAddr := addrs[:types.ADDRESSLENGTH]
	delAddr := addrs[types.ADDRESSLENGTH : 2*types.ADDRESSLENGTH]
	valDstAddr := addrs[2*types.ADDRESSLENGTH:]

	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}

// rearranges the ValDstIndexKey to get the REDKey
func GetREDKeyFromValDstIndexKey(IndexKey []byte) []byte {
	addrs := IndexKey[1:] // remove prefix bytes
	if len(addrs) != 3*types.ADDRESSLENGTH {
		panic("unexpected key length")
	}
	valDstAddr := addrs[:types.ADDRESSLENGTH]
	delAddr := addrs[types.ADDRESSLENGTH : 2*types.ADDRESSLENGTH]
	valSrcAddr := addrs[2*types.ADDRESSLENGTH:]
	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}

// gets the prefix keyspace for all redelegations redelegating away from a source validator
func GetREDsFromValSrcIndexKey(valSrcAddr sdk.Address) []byte {
	return append(RedelegationByValSrcIndexKey, valSrcAddr.Bytes()...)
}

// gets the prefix keyspace for all redelegations redelegating towards a destination validator
func GetREDsToValDstIndexKey(valDstAddr sdk.Address) []byte {
	return append(RedelegationByValDstIndexKey, valDstAddr.Bytes()...)
}

// gets the prefix keyspace for all redelegations redelegating towards a destination validator
// from a particular delegator
func GetREDsByDelToValDstIndexKey(delAddr sdk.Address, valDstAddr sdk.Address) []byte {
	return append(
		GetREDsToValDstIndexKey(valDstAddr),
		delAddr.Bytes()...)
}
