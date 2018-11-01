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
	DelegatorAddr    sdk.Address `json:"delegator_addr"`
	ValidatorAddr    sdk.Address `json:"validator_addr"`
	Shares           types.Dec   `json:"shares"`
	Height           int64       `json:"height"`           // Last height bond updated
	RewardAccum      types.Coin `json:"reward_accum"`     // reward accumulation of this block til withdrawal_height
	WithdrawalHeight int64       `json:"withdrawal_height` // latest withdrawal height
}

type delegationValue struct {
	Shares           types.Dec
	Height           int64
	RewardAccum      types.Coin
	WithdrawalHeight int64
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
		delegation.RewardAccum,
		delegation.WithdrawalHeight,
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
		DelegatorAddr:    delAddr,
		ValidatorAddr:    valAddr,
		Shares:           storeValue.Shares,
		Height:           storeValue.Height,
		RewardAccum:      storeValue.RewardAccum,
		WithdrawalHeight: storeValue.WithdrawalHeight,
	}, nil
}

// ensure fulfills the sdk validator types
// var _ sdk.Delegation = Delegation{}

// nolint - for sdk.Delegation
func (b Delegation) GetDelegator() sdk.Address { return b.DelegatorAddr }
func (b Delegation) GetValidator() sdk.Address { return b.ValidatorAddr }
func (b Delegation) GetBondShares() types.Dec  { return b.Shares }

// UpdateDelReward updating reward accumulation
func (b Delegation) UpdateDelAccum(
	currentHeight int64,
	totalRewardAccum types.Coin,
) Delegation {

	rewardCoin := totalRewardAccum.Mul(b.Shares)

	b.RewardAccum = b.RewardAccum.Plus(rewardCoin)

	b.WithdrawalHeight = currentHeight

	return b
}

//Human Friendly pretty printer
func (b Delegation) HumanReadableString() (string, error) {

	resp := "Delegation \n"
	resp += fmt.Sprintf("Delegator: %s\n", b.DelegatorAddr.String())
	resp += fmt.Sprintf("Validator: %s\n", b.ValidatorAddr.String())
	resp += fmt.Sprintf("Shares: %s", b.Shares.String())
	resp += fmt.Sprintf("Height: %d", b.Height)
	resp += fmt.Sprintf("RewardAccum: %s", b.RewardAccum)
	resp += fmt.Sprintf("WithdrawalHeigh: %s", b.WithdrawalHeight)

	return resp, nil
}

// UnbondingDelegation reflects a delegation's passive unbonding queue.
type UnbondingDelegation struct {
	DelegatorAddr  sdk.Address `json:"delegator_addr"`  // delegator
	ValidatorAddr  sdk.Address `json:"validator_addr"`  // validator unbonding from operator addr
	CreationHeight int64       `json:"creation_height"` // height which the unbonding took place
	MinTime        int64       `json:"min_time"`        // unix time for unbonding completion  /*time.Time*/
	InitialBalance types.Coin  `json:"initial_balance"` // atoms initially scheduled to receive at completion
	Balance        types.Coin  `json:"balance"`         // atoms to receive at completion
}

type ubdValue struct {
	CreationHeight int64
	MinTime        int64 //time.Time
	InitialBalance types.Coin
	Balance        types.Coin
}

// return the unbonding delegation without fields contained within the key for the store
func MustMarshalUBD(cdc *wire.Codec, ubd UnbondingDelegation) []byte {
	val := ubdValue{
		ubd.CreationHeight,
		ubd.MinTime,
		ubd.InitialBalance,
		ubd.Balance,
	}
	return cdc.MustMarshalBinary(val)
}

// unmarshal a unbonding delegation from a store key and value
func MustUnmarshalUBD(cdc *wire.Codec, key, value []byte) UnbondingDelegation {
	ubd, err := UnmarshalUBD(cdc, key, value)
	if err != nil {
		panic(err)
	}
	return ubd
}

// unmarshal a unbonding delegation from a store key and value
func UnmarshalUBD(cdc *wire.Codec, key, value []byte) (ubd UnbondingDelegation, err error) {
	var storeValue ubdValue
	err = cdc.UnmarshalBinary(value, &storeValue)
	if err != nil {
		return
	}

	addrs := key[1:] // remove prefix bytes
	if len(addrs) != 2*types.ADDRESSLENGTH {
		//err = fmt.Errorf("%v", ErrBadDelegationAddr(DefaultCodespace).Data())
		return
	}
	delAddr := sdk.Address(addrs[:types.ADDRESSLENGTH])
	valAddr := sdk.Address(addrs[types.ADDRESSLENGTH:])

	return UnbondingDelegation{
		DelegatorAddr:  delAddr,
		ValidatorAddr:  valAddr,
		CreationHeight: storeValue.CreationHeight,
		MinTime:        storeValue.MinTime,
		InitialBalance: storeValue.InitialBalance,
		Balance:        storeValue.Balance,
	}, nil
}
