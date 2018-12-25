package posTypes

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-amino"
	"github.com/sharering/shareledger/types"
)

// Delegation represents the bond with tokens held by an account.  It is
// owned by one delegator, and is associated with the voting power of one
// pubKey.
type Delegation struct {
	DelegatorAddr    sdk.AccAddress `json:"delegator_addr"`
	ValidatorAddr    sdk.AccAddress `json:"validator_addr"`
	Shares           types.Dec   `json:"shares"`
	Height           int64       `json:"height"`           // Last height bond updated
	RewardAccum      types.Coin  `json:"reward_accum"`     // reward accumulation of this block til withdrawal_height
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
func MustMarshalDelegation(cdc *amino.Codec, delegation Delegation) []byte {
	val := delegationValue{
		delegation.Shares,
		delegation.Height,
		delegation.RewardAccum,
		delegation.WithdrawalHeight,
	}
	return cdc.MustMarshalBinaryLengthPrefixed(val)
}

// return the delegation without fields contained within the key for the store
func MustUnmarshalDelegation(cdc *amino.Codec, key, value []byte) Delegation {
	delegation, err := UnmarshalDelegation(cdc, key, value)
	if err != nil {
		panic(err)
	}
	return delegation
}

// return the delegation without fields contained within the key for the store
func UnmarshalDelegation(cdc *amino.Codec, key, value []byte) (delegation Delegation, err error) {
	var storeValue delegationValue
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &storeValue)
	if err != nil {
		//err = fmt.Errorf("%v: %v", ErrNoDelegation(DefaultCodespace).Data(), err)
		return
	}

	addrs := key[1:] // remove prefix bytes
	if len(addrs) != 2*types.ADDRESSLENGTH {
		//err = fmt.Errorf("%v", ErrBadDelegationAddr(DefaultCodespace).Data())
		return
	}

	delAddr := sdk.AccAddress(addrs[:types.ADDRESSLENGTH])
	valAddr := sdk.AccAddress(addrs[types.ADDRESSLENGTH:])

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
func (b Delegation) GetDelegator() sdk.AccAddress { return b.DelegatorAddr }
func (b Delegation) GetValidator() sdk.AccAddress { return b.ValidatorAddr }
func (b Delegation) GetBondShares() types.Dec  { return b.Shares }

// UpdateDelReward updating reward accumulation
func (b Delegation) UpdateDelAccum(
	currentHeight int64,
	totalRewardAccum types.Coin,
	totalShares types.Dec,
) Delegation {
	fmt.Printf("TotalRewardAccum: %v\n", totalRewardAccum)
	fmt.Printf("Shares/TotalShares: %v/%v\n", b.Shares, totalShares)
	rewardCoin := totalRewardAccum.Mul(b.Shares).Quo(totalShares)
	fmt.Printf("RewardCoin: %v\n", rewardCoin)
	b.RewardAccum = b.RewardAccum.Plus(rewardCoin)
	fmt.Printf("RewardAccum: %v\n", b.RewardAccum)

	b.WithdrawalHeight = currentHeight

	return b
}

//Human Friendly pretty printer
func (b Delegation) HumanReadableString() (string, error) {

	resp := "Delegation \n"
	resp += fmt.Sprintf("Delegator: %s\n", b.DelegatorAddr.String())
	resp += fmt.Sprintf("Validator: %s\n", b.ValidatorAddr.String())
	resp += fmt.Sprintf("Shares: %s\n", b.Shares.String())
	resp += fmt.Sprintf("Height: %d\n", b.Height)
	resp += fmt.Sprintf("RewardAccum: %s\n", b.RewardAccum)
	resp += fmt.Sprintf("WithdrawalHeight: %d\n", b.WithdrawalHeight)

	return resp, nil
}

// UnbondingDelegation reflects a delegation's passive unbonding queue.
type UnbondingDelegation struct {
	DelegatorAddr  sdk.AccAddress `json:"delegator_addr"`  // delegator
	ValidatorAddr  sdk.AccAddress `json:"validator_addr"`  // validator unbonding from operator addr
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
func MustMarshalUBD(cdc *amino.Codec, ubd UnbondingDelegation) []byte {
	val := ubdValue{
		ubd.CreationHeight,
		ubd.MinTime,
		ubd.InitialBalance,
		ubd.Balance,
	}
	return cdc.MustMarshalBinaryLengthPrefixed(val)
}

// unmarshal a unbonding delegation from a store key and value
func MustUnmarshalUBD(cdc *amino.Codec, key, value []byte) UnbondingDelegation {
	ubd, err := UnmarshalUBD(cdc, key, value)
	if err != nil {
		panic(err)
	}
	return ubd
}

// unmarshal a unbonding delegation from a store key and value
func UnmarshalUBD(cdc *amino.Codec, key, value []byte) (ubd UnbondingDelegation, err error) {
	var storeValue ubdValue
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &storeValue)
	if err != nil {
		return
	}

	addrs := key[1:] // remove prefix bytes
	if len(addrs) != 2*types.ADDRESSLENGTH {
		//err = fmt.Errorf("%v", ErrBadDelegationAddr(DefaultCodespace).Data())
		return
	}
	delAddr := sdk.AccAddress(addrs[:types.ADDRESSLENGTH])
	valAddr := sdk.AccAddress(addrs[types.ADDRESSLENGTH:])

	return UnbondingDelegation{
		DelegatorAddr:  delAddr,
		ValidatorAddr:  valAddr,
		CreationHeight: storeValue.CreationHeight,
		MinTime:        storeValue.MinTime,
		InitialBalance: storeValue.InitialBalance,
		Balance:        storeValue.Balance,
	}, nil
}

// Redelegation reflects a delegation's passive re-delegation queue.
type Redelegation struct {
	DelegatorAddr    sdk.AccAddress `json:"delegator_addr"`     // delegator
	ValidatorSrcAddr sdk.AccAddress `json:"validator_src_addr"` // validator redelegation source operator addr
	ValidatorDstAddr sdk.AccAddress `json:"validator_dst_addr"` // validator redelegation destination operator addr
	CreationHeight   int64       `json:"creation_height"`    // height which the redelegation took place
	MinTime          int64       `json:"min_time"`           // unix time for redelegation completion
	InitialBalance   types.Coin  `json:"initial_balance"`    // initial balance when redelegation started
	Balance          types.Coin  `json:"balance"`            // current balance
	SharesSrc        types.Dec   `json:"shares_src"`         // amount of source shares redelegating
	SharesDst        types.Dec   `json:"shares_dst"`         // amount of destination shares redelegating
}

type redValue struct {
	CreationHeight int64
	MinTime        int64 //time.Time
	InitialBalance types.Coin
	Balance        types.Coin
	SharesSrc      types.Dec
	SharesDst      types.Dec
}

// return the redelegation without fields contained within the key for the store
func MustMarshalRED(cdc *amino.Codec, red Redelegation) []byte {
	val := redValue{
		red.CreationHeight,
		red.MinTime,
		red.InitialBalance,
		red.Balance,
		red.SharesSrc,
		red.SharesDst,
	}
	return cdc.MustMarshalBinaryLengthPrefixed(val)
}

// unmarshal a redelegation from a store key and value
func MustUnmarshalRED(cdc *amino.Codec, key, value []byte) Redelegation {
	red, err := UnmarshalRED(cdc, key, value)
	if err != nil {
		panic(err)
	}
	return red
}

// unmarshal a redelegation from a store key and value
func UnmarshalRED(cdc *amino.Codec, key, value []byte) (red Redelegation, err error) {
	var storeValue redValue
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &storeValue)
	if err != nil {
		return
	}

	addrs := key[1:] // remove prefix bytes
	if len(addrs) != 3*types.ADDRESSLENGTH {
		//err = fmt.Errorf("%v", posTypes.ErrBadRedelegationAddr(DefaultCodespace).Data())
		return
	}
	delAddr := sdk.AccAddress(addrs[:types.ADDRESSLENGTH])
	valSrcAddr := sdk.AccAddress(addrs[types.ADDRESSLENGTH : 2*types.ADDRESSLENGTH])
	valDstAddr := sdk.AccAddress(addrs[2*types.ADDRESSLENGTH:])

	return Redelegation{
		DelegatorAddr:    delAddr,
		ValidatorSrcAddr: valSrcAddr,
		ValidatorDstAddr: valDstAddr,
		CreationHeight:   storeValue.CreationHeight,
		MinTime:          storeValue.MinTime,
		InitialBalance:   storeValue.InitialBalance,
		Balance:          storeValue.Balance,
		SharesSrc:        storeValue.SharesSrc,
		SharesDst:        storeValue.SharesDst,
	}, nil
}
