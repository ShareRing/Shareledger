package types

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

// status of a validator
type BondStatus byte

// nolint
const (
	Unbonded  BondStatus = 0x00
	Unbonding BondStatus = 0x01
	Bonded    BondStatus = 0x02
)

//BondStatusToString for pretty prints of Bond Status
func BondStatusToString(b BondStatus) string {
	switch b {
	case 0x00:
		return "Unbonded"
	case 0x01:
		return "Unbonding"
	case 0x02:
		return "Bonded"
	default:
		return ""
	}
}

// validator for a delegated proof of stake system
type Validator interface {
	GetMoniker() string          // moniker of the validator
	GetStatus() BondStatus       // status of the validator
	GetOwner() sdk.Address       // owner address to receive/return validators coins
	GetPubKey() PubKey           // validation pubkey
	GetPower() sdk.Rat           // validation power
	GetDelegatorShares() sdk.Rat // Total out standing delegator shares
	GetBondHeight() int64        // height in which the validator became active
}

// validator which fulfills abci validator interface for use in Tendermint
func ABCIValidator(v Validator) abci.Validator {
	var pubKey crypto.PubKeySecp256k1
	if pk, ok := v.GetPubKey().(PubKeySecp256k1); ok {

		copy(pubKey[:], pk[:65])

		return abci.Validator{
			PubKey: tmtypes.TM2PB.PubKey(pubKey),
			Power:  v.GetPower().Evaluate(),
		}
	} else {
		panic("PubKey is not of PubKeySecp256k1")
	}
}

// properties for the set of all validators
type ValidatorSet interface {
	// iterate through validator by owner-address, execute func for each validator
	IterateValidators(sdk.Context,
		func(index int64, validator Validator) (stop bool))

	// iterate through bonded validator by pubkey-address, execute func for each validator
	IterateValidatorsBonded(sdk.Context,
		func(index int64, validator Validator) (stop bool))

	Validator(sdk.Context, sdk.Address) Validator // get a particular validator by owner address
	TotalPower(sdk.Context) sdk.Rat               // total power of the validator set
	Slash(sdk.Context, PubKey, int64, sdk.Rat)    // slash the validator and delegators of the validator, specifying offence height & slash fraction
	Revoke(sdk.Context, PubKey)                   // revoke a validator
	Unrevoke(sdk.Context, PubKey)                 // unrevoke a validator
}

//_______________________________________________________________________________

// delegation bond for a delegated proof of stake system
type Delegation interface {
	GetDelegator() sdk.Address // delegator address for the bond
	GetValidator() sdk.Address // validator owner address for the bond
	GetBondShares() sdk.Rat    // amount of validator's shares
}

// properties for the set of all delegations for a particular
type DelegationSet interface {
	GetValidatorSet() ValidatorSet // validator set for which delegation set is based upon

	// iterate through all delegations from one delegator by validator-address,
	//   execute func for each validator
	IterateDelegations(ctx sdk.Context, delegator sdk.Address,
		fn func(index int64, delegation Delegation) (stop bool))
}
