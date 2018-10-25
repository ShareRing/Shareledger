package posTypes

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/types"
)

// Delegation represents the bond with tokens held by an account.  It is
// owned by one delegator, and is associated with the voting power of one
// pubKey.
type Delegation struct {
	DelegatorAddr sdk.Address `json:"delegator_addr"`
	ValidatorAddr sdk.Address `json:"validator_addr"`
	Shares        types.Dec   `json:"shares"`
	Height        int64       `json:"height"` // Last height bond updated
}

type delegationValue struct {
	Shares types.Dec
	Height int64
}

// aggregates of all delegations, unbondings and redelegations
/*
type DelegationSummary struct {
	Delegations          []Delegation          `json:"delegations"`
	UnbondingDelegations []UnbondingDelegation `json:"unbonding_delegations"`
	Redelegations        []Redelegation        `json:"redelegations"`
}
*/
// return the delegation without fields contained within the key for the store
func MustMarshalDelegation(cdc *wire.Codec, delegation Delegation) []byte {
	val := delegationValue{
		delegation.Shares,
		delegation.Height,
	}
	return cdc.MustMarshalBinary(val)
}

// return the delegation without fields contained within the key for the store
func MustUnmarshalDelegation(cdc *wire.Codec, key, value []byte) Delegation {
	delegation, err := UnmarshalDelegation(cdc, key, value)
	if err != nil {
		panic(err)
	}
	return delegation
}

// return the delegation without fields contained within the key for the store
func UnmarshalDelegation(cdc *wire.Codec, key, value []byte) (delegation Delegation, err error) {
	var storeValue delegationValue
	err = cdc.UnmarshalBinary(value, &storeValue)
	if err != nil {
		//err = fmt.Errorf("%v: %v", ErrNoDelegation(DefaultCodespace).Data(), err)
		return
	}

	addrs := key[1:] // remove prefix bytes
	if len(addrs) != 2*types.ADDRESSLENGTH {
		//err = fmt.Errorf("%v", ErrBadDelegationAddr(DefaultCodespace).Data())
		return
	}

	delAddr := sdk.Address(addrs[:types.ADDRESSLENGTH])
	valAddr := sdk.Address(addrs[types.ADDRESSLENGTH:])

	return Delegation{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
		Shares:        storeValue.Shares,
		Height:        storeValue.Height,
	}, nil
}

// ensure fulfills the sdk validator types
// var _ sdk.Delegation = Delegation{}

// nolint - for sdk.Delegation
func (b Delegation) GetDelegator() sdk.Address { return b.DelegatorAddr }
func (b Delegation) GetValidator() sdk.Address { return b.ValidatorAddr }
func (b Delegation) GetBondShares() types.Dec  { return b.Shares }

//Human Friendly pretty printer
func (b Delegation) HumanReadableString() (string, error) {

	resp := "Delegation \n"
	resp += fmt.Sprintf("Delegator: %s\n", b.DelegatorAddr.String())
	resp += fmt.Sprintf("Validator: %s\n", b.ValidatorAddr.String())
	resp += fmt.Sprintf("Shares: %s", b.Shares.String())
	resp += fmt.Sprintf("Height: %d", b.Height)

	return resp, nil
}
