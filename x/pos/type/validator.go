package posTypes

import (
	"bytes"
	"fmt"
	"time"

	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/types"
)

// Validator defines the total amount of bond shares and their exchange rate to
// coins. Accumulation of interest is modelled as an in increase in the
// exchange rate, and slashing as a decrease.  When coins are delegated to this
// validator, the validator is credited with a Delegation whose number of
// bond shares is based on the amount of coins delegated divided by the current
// exchange rate. Voting power can be calculated as total bonds multiplied by
// exchange rate.
type Validator struct {
	Owner   sdk.Address      `json:"owner"`   // sender of BondTx - UnbondTx returns here
	PubKey  types.PubKey     `json:"pub_key"` // pubkey of validator
	Revoked bool             `json:"revoked"` // has the validator  been revoked from bonded status?
	Status  types.BondStatus `json:"status"`  // validator status (bonded/unbonding/unbonded)

	Tokens          types.Dec `json:"tokens"`           // delegated tokens (incl. self-delegation)
	DelegatorShares types.Dec `json:"delegator_shares"` // total shares issued to a validator's delegators

	Description        Description `json:"description"`           // description terms for the validator
	BondHeight         int64       `json:"bond_height"`           // earliest height as a bonded validator
	BondIntraTxCounter int16       `json:"bond_intra_tx_counter"` // block-local tx index of validator change
	UnbondingHeight    int64       `json:"unbonding_height"`      // if unbonding, height at which this validator has begun unbonding
	UnbondingMinTime   time.Time   `json:"unbonding_time"`        // if unbonding, min time for the validator to complete unbonding
	//	ProposerRewardPool sdk.Coins   `json:"proposer_reward_pool"`  // XXX reward pool collected from being the proposer

}

// enforce the Validator type at compile time
var _ types.Validator = Validator{}

// Validators - list of Validators
type Validators []Validator

// to encode/decode of Validator
type validatorValue struct {
	PubKey             types.PubKey
	Revoked            bool
	Status             types.BondStatus
	Tokens             types.Dec
	DelegatorShares    types.Dec
	Description        Description
	BondHeight         int64
	BondIntraTxCounter int16
	UnbondingHeight    int64
	UnbondingMinTime   time.Time
}

// NewValidator - initialize a new validator
func NewValidator(owner sdk.Address, pubKey types.PubKey, description Description) Validator {
	return Validator{
		Owner:   owner,
		PubKey:  pubKey,
		Revoked: false,
		Status:  types.Unbonded,

		Tokens:             types.ZeroDec(),
		DelegatorShares:    types.ZeroDec(),
		Description:        description,
		BondHeight:         int64(0),
		BondIntraTxCounter: int16(0),
		UnbondingHeight:    int64(0),
		UnbondingMinTime:   time.Unix(0, 0).UTC(),
		//ProposerRewardPool: sdk.Coins{},
	}
}

// only the vitals - does not check bond height of IntraTxCounter
func (v Validator) Equal(c2 Validator) bool {
	return v.PubKey.Equals(c2.PubKey) &&
		bytes.Equal(v.Owner, c2.Owner) &&
		// v.PoolShares.Equal(c2.PoolShares) &&
		v.Tokens.Equal(c2.Tokens) &&
		v.DelegatorShares.Equal(c2.DelegatorShares) &&
		v.Description == c2.Description //&&
	//v.BondHeight == c2.BondHeight &&
	//v.BondIntraTxCounter == c2.BondIntraTxCounter && // counter is always changing
	// v.ProposerRewardPool.IsEqual(c2.ProposerRewardPool)
}

const DoNotModifyDesc = "[do-not-modify]"

// Description - description fields for a validator
type Description struct {
	Moniker  string `json:"moniker"`  // name
	Identity string `json:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website"`  // optional website link
	Details  string `json:"details"`  // optional details
}

func NewDescription(moniker, identity, website, details string) Description {
	return Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

func (v Validator) GetABCIPubKey() crypto.PubKeySecp256k1 {
	if pk, ok := v.GetPubKey().(types.PubKeySecp256k1); ok {
		return pk.ToABCIPubKey()
	} else {
		panic("PubKey is not of PubKeySecp256k1")
	}

}

// validator which fulfills abci validator interface for use in Tendermint
// ABCIValidator returns an abci.Validator from a staked validator type.
func (v Validator) ABCIValidator() abci.Validator {
	return abci.Validator{
		PubKey:  tmtypes.TM2PB.PubKey(v.GetABCIPubKey()),
		Address: v.GetPubKey().Address(),
		Power:   v.BondedTokens().RoundInt64(),
	}
}

// ABCIValidator returns an abci.Validator from a staked validator type.
func (v Validator) ABCIValidatorZero() abci.Validator {
	return abci.Validator{
		PubKey:  tmtypes.TM2PB.PubKey(v.GetABCIPubKey()),
		Address: v.GetPubKey().Address(),
		Power:   0,
	}
}

// UpdateStatus updates the location of the shares within a validator
// to reflect the new status
func (v Validator) UpdateStatus(pool Pool, NewStatus types.BondStatus) (Validator, Pool) {

	switch v.Status {

	case types.Unbonded:

		switch NewStatus {
		case types.Unbonded:
			return v, pool
		case types.Bonded:
			pool = pool.looseTokensToBonded(v.Tokens)
		}
	case types.Unbonding:

		switch NewStatus {
		case types.Unbonding:
			return v, pool
		case types.Bonded:
			pool = pool.looseTokensToBonded(v.Tokens)
		}
	case types.Bonded:

		switch NewStatus {
		case types.Bonded:
			return v, pool
		default:
			pool = pool.bondedTokensToLoose(v.Tokens)
		}
	}

	v.Status = NewStatus
	return v, pool
}

// Returns if the validator should be considered unbonded
func (v Validator) IsUnbonded(ctx sdk.Context) bool {
	switch v.Status {
	case types.Unbonded:
		return true
	case types.Unbonding:
		//todo: check the time if it surpass the unboundtingTime
		//ctxTime := ctx.BlockHeader().Time

		//if ctxTime.After(v.UnbondingMinTime) {
		//		return true
		//	}
		return false
	}
	return false
}

// removes tokens from a validator
func (v Validator) RemoveTokens(pool Pool, tokens types.Dec) (Validator, Pool) {
	if v.Status == types.Bonded {
		pool = pool.bondedTokensToLoose(tokens)
	}

	v.Tokens = v.Tokens.Sub(tokens)
	return v, pool
}

// SetInitialCommission attempts to set a validator's initial commission. An
// error is returned if the commission is invalid.
// func (v Validator) SetInitialCommission(commission Commission) (Validator, sdk.Error) {
// 	if err := commission.Validate(); err != nil {
// 		return v, err
// 	}

// 	v.Commission = commission
// 	return v, nil
// }

//_________________________________________________________________________________________________________

// XXX Audit this function further to make sure it's correct
// add tokens to a validator
func (v Validator) AddTokensFromDel(pool Pool, amount types.Int) (Validator, Pool, types.Dec) {

	// bondedShare/delegatedShare
	exRate := v.DelegatorShareExRate()
	amountDec := types.NewDecFromInt(amount)

	if v.Status == types.Bonded {
		pool = pool.looseTokensToBonded(amountDec)
	}

	v.Tokens = v.Tokens.Add(amountDec)
	issuedShares := amountDec.Quo(exRate)
	v.DelegatorShares = v.DelegatorShares.Add(issuedShares)

	return v, pool, issuedShares
}

// RemoveDelShares removes delegator shares from a validator.
func (v Validator) RemoveDelShares(pool Pool, delShares types.Dec) (Validator, Pool, types.Dec) {
	issuedTokens := v.DelegatorShareExRate().Mul(delShares)
	v.Tokens = v.Tokens.Sub(issuedTokens)
	v.DelegatorShares = v.DelegatorShares.Sub(delShares)

	if v.Status == types.Bonded {
		pool = pool.bondedTokensToLoose(issuedTokens)
	}

	return v, pool, issuedTokens
}

// DelegatorShareExRate gets the exchange rate of tokens over delegator shares.
// UNITS: tokens/delegator-shares
func (v Validator) DelegatorShareExRate() types.Dec {
	if v.DelegatorShares.IsZero() {
		return types.OneDec()
	}
	return v.Tokens.Quo(v.DelegatorShares)
}

// Get the bonded tokens which the validator holds
func (v Validator) BondedTokens() types.Dec {
	if v.Status == types.Bonded {
		return v.Tokens
	}
	return types.ZeroDec()
}

// unmarshal a redelegation from a store key and value
func UnmarshalValidator(cdc *wire.Codec, owner sdk.Address, value []byte) (validator Validator, err error) {
	//TODO: Checking owner address

	var storeValue validatorValue
	err = cdc.UnmarshalBinary(value, &storeValue)
	if err != nil {
		return
	}

	return Validator{
		Owner:              owner,
		PubKey:             storeValue.PubKey,
		Revoked:            storeValue.Revoked,
		Tokens:             storeValue.Tokens,
		Status:             storeValue.Status,
		DelegatorShares:    storeValue.DelegatorShares,
		Description:        storeValue.Description,
		BondHeight:         storeValue.BondHeight,
		BondIntraTxCounter: storeValue.BondIntraTxCounter,
		UnbondingHeight:    storeValue.UnbondingHeight,
		UnbondingMinTime:   storeValue.UnbondingMinTime,
	}, nil
}

// return the redelegation without fields contained within the key for the store
func MustMarshalValidator(cdc *wire.Codec, validator Validator) []byte {
	val := validatorValue{
		PubKey:             validator.PubKey,
		Revoked:            validator.Revoked,
		Status:             validator.Status,
		Tokens:             validator.Tokens,
		DelegatorShares:    validator.DelegatorShares,
		Description:        validator.Description,
		BondHeight:         validator.BondHeight,
		BondIntraTxCounter: validator.BondIntraTxCounter,
		UnbondingHeight:    validator.UnbondingHeight,
		UnbondingMinTime:   validator.UnbondingMinTime,
	}
	return cdc.MustMarshalBinary(val)
}

//______________________________________________________________________

var _ types.Validator = Validator{}

// nolint - for sdk.Validator
func (v Validator) GetMoniker() string          { return v.Description.Moniker }
func (v Validator) GetStatus() types.BondStatus { return v.Status }

func (v Validator) GetOwner() sdk.Address         { return v.Owner }
func (v Validator) GetPubKey() types.PubKey       { return v.PubKey }
func (v Validator) GetPower() types.Dec           { return v.BondedTokens() }
func (v Validator) GetDelegatorShares() types.Dec { return v.DelegatorShares }
func (v Validator) GetBondHeight() int64          { return v.BondHeight }

//Human Friendly pretty printer
func (v Validator) HumanReadableString() (string, error) {

	resp := "Validator \n"
	resp += fmt.Sprintf("Owner: %s\n", v.Owner.String())
	resp += fmt.Sprintf("Validator: %s\n", v.PubKey.String())
	//resp += fmt.Sprintf("Shares: Status %s,  Amount: %s\n", sdk.BondStatusToString(v.PoolShares.Status), v.PoolShares.Amount.String())
	resp += fmt.Sprintf("Delegator Shares: %s\n", v.DelegatorShares.String())
	resp += fmt.Sprintf("Description: %s\n", v.Description)
	resp += fmt.Sprintf("Bond Height: %d\n", v.BondHeight)
	//	resp += fmt.Sprintf("Proposer Reward Pool: %s\n", v.ProposerRewardPool.String())

	return resp, nil
}
