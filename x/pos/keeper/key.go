package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
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
	// Last* values are const during a block.
	LastValidatorPowerKey = []byte{0x11} // prefix for each key to a validator index, for bonded validators
	LastTotalPowerKey     = []byte{0x12} // prefix for the total power
)

// gets the key for the validator with address
// VALUE: stake/types.Validator
func GetValidatorKey(operatorAddr sdk.AccAddress) []byte {
	return append(ValidatorsKey, operatorAddr.Bytes()...)
}

func GetValidatorDistKey(operatorAddr sdk.AccAddress) []byte {
	return append(ValidatorDistKey, operatorAddr.Bytes()...)
}

// gets the key for the validator with pubkey
// VALUE: validator operator address ([]byte)
func GetValidatorByConsAddrKey(addr sdk.AccAddress) []byte {
	return append(ValidatorsByConsAddrKey, addr.Bytes()...)
}

//______________________________________________________________________________

// gets the key for delegator bond with validator
// VALUE: stake/types.Delegation
func GetDelegationKey(delAddr sdk.AccAddress, valAddr sdk.AccAddress) []byte {
	return append(GetDelegationsKey(delAddr), valAddr.Bytes()...)
}

// gets the prefix for a delegator for all validators
func GetDelegationsKey(delAddr sdk.AccAddress) []byte {
	return append(DelegationKey, delAddr.Bytes()...)
}

// gets the prefix keyspace for the indexes of unbonding delegations for a validator
func GetUBDsByValIndexKey(valAddr sdk.AccAddress) []byte {
	return append(UnbondingDelegationByValIndexKey, valAddr.Bytes()...)
}

// gets the prefix for all unbonding delegations from a delegator
func GetUBDsKey(delAddr sdk.AccAddress) []byte {
	return append(UnbondingDelegationKey, delAddr.Bytes()...)
}

// gets the key for an unbonding delegation by delegator and validator addr
// VALUE: stake/types.UnbondingDelegation
func GetUBDKey(delAddr sdk.AccAddress, valAddr sdk.AccAddress) []byte {
	return append(
		GetUBDsKey(delAddr.Bytes()),
		valAddr.Bytes()...)
}

// gets the index-key for an unbonding delegation, stored by validator-index
// VALUE: none (key rearrangement used)
func GetUBDByValIndexKey(delAddr sdk.AccAddress, valAddr sdk.AccAddress) []byte {
	return append(GetUBDsByValIndexKey(valAddr), delAddr.Bytes()...)
}

// gets the prefix keyspace for redelegations from a delegator
func GetREDsKey(delAddr sdk.AccAddress) []byte {
	return append(RedelegationKey, delAddr.Bytes()...)
}

// gets the key for a redelegation
// VALUE: stake/types.RedelegationKey
func GetREDKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.AccAddress) []byte {
	return append(append(
		GetREDsKey(delAddr.Bytes()),
		valSrcAddr.Bytes()...),
		valDstAddr.Bytes()...)
}

// gets the index-key for a redelegation, stored by source-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValSrcIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.AccAddress) []byte {
	return append(append(
		GetREDsFromValSrcIndexKey(valSrcAddr),
		delAddr.Bytes()...),
		valDstAddr.Bytes()...)
}

// gets the index-key for a redelegation, stored by destination-validator-index
// VALUE: none (key rearrangement used)
func GetREDByValDstIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.AccAddress) []byte {
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
func GetREDsFromValSrcIndexKey(valSrcAddr sdk.AccAddress) []byte {
	return append(RedelegationByValSrcIndexKey, valSrcAddr.Bytes()...)
}

// gets the prefix keyspace for all redelegations redelegating towards a destination validator
func GetREDsToValDstIndexKey(valDstAddr sdk.AccAddress) []byte {
	return append(RedelegationByValDstIndexKey, valDstAddr.Bytes()...)
}

// gets the prefix keyspace for all redelegations redelegating towards a destination validator
// from a particular delegator
func GetREDsByDelToValDstIndexKey(delAddr sdk.AccAddress, valDstAddr sdk.AccAddress) []byte {
	return append(
		GetREDsToValDstIndexKey(valDstAddr),
		delAddr.Bytes()...)
}

// get the validator by power index.
// Power index is the key used in the power-store, and represents the relative
// power ranking of the validator.
// VALUE: validator operator address ([]byte)
func GetValidatorsByPowerIndexKey(validator posTypes.Validator, pool posTypes.Pool) []byte {
	// NOTE the address doesn't need to be stored because counter bytes must always be different
	return getValidatorPowerRank(validator)
}

// get the bonded validator index key for an operator address
func GetLastValidatorPowerKey(operator sdk.AccAddress) []byte {
	return append(LastValidatorPowerKey, operator...)
}

// get the power ranking of a validator
// NOTE the larger values are of higher value
// nolint: unparam
func getValidatorPowerRank(validator posTypes.Validator) []byte {

	potentialPower := validator.Tokens

	// todo: deal with cases above 2**64, ref https://github.com/cosmos/cosmos-sdk/issues/2439#issuecomment-427167556
	tendermintPower := potentialPower.RoundInt64()
	tendermintPowerBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(tendermintPowerBytes[:], uint64(tendermintPower))

	powerBytes := tendermintPowerBytes
	powerBytesLen := len(powerBytes)

	// key is of format prefix || powerbytes || heightBytes || counterBytes
	key := make([]byte, 1+powerBytesLen+8+2)

	key[0] = ValidatorsByPowerIndexKey[0]
	copy(key[1:powerBytesLen+1], powerBytes)

	// include heightBytes height is inverted (older validators first)
	binary.BigEndian.PutUint64(key[powerBytesLen+1:powerBytesLen+9], ^uint64(validator.BondHeight))
	// include counterBytes, counter is inverted (first txns have priority)
	binary.BigEndian.PutUint16(key[powerBytesLen+9:powerBytesLen+11], ^uint16(validator.BondIntraTxCounter))

	return key
}
