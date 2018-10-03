package pos

import (
	"fmt"
	"time"

	"github.com/sharering/shareledger/types"
	abci "github.com/tendermint/abci/types"
	crypto "github.com/tendermint/go-crypto"
	tmtypes "github.com/tendermint/tendermint/types"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
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

	Tokens          sdk.Rat `json:"tokens"`           // delegated tokens (incl. self-delegation)
	DelegatorShares sdk.Rat `json:"delegator_shares"` // total shares issued to a validator's delegators

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

// NewValidator - initialize a new validator
func NewValidator(owner sdk.Address, pubKey types.PubKeySecp256k1, description Description) Validator {
	return Validator{
		Owner:   owner,
		PubKey:  pubKey,
		Revoked: false,

		DelegatorShares:    sdk.ZeroRat(),
		Description:        description,
		BondHeight:         int64(0),
		BondIntraTxCounter: int16(0),
		UnbondingHeight:    int64(0),
		UnbondingMinTime:   time.Unix(0, 0).UTC(),
		//ProposerRewardPool: sdk.Coins{},
	}
}

// only the vitals - does not check bond height of IntraTxCounter
/*
func (v Validator) equal(c2 Validator) bool {
	return v.PubKey.Equals(c2.PubKey) &&
		bytes.Equal(v.Owner, c2.Owner) &&
		v.PoolShares.Equal(c2.PoolShares) &&
		v.DelegatorShares.Equal(c2.DelegatorShares) &&
		v.Description == c2.Description &&
		//v.BondHeight == c2.BondHeight &&
		//v.BondIntraTxCounter == c2.BondIntraTxCounter && // counter is always changing
		v.ProposerRewardPool.IsEqual(c2.ProposerRewardPool)
}
*/
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

// validator which fulfills abci validator interface for use in Tendermint
func (v Validator) ABCIValidator() abci.Validator {
	var pubKey crypto.PubKeySecp256k1
	if pk, ok := v.GetPubKey().(types.PubKeySecp256k1); ok {

		copy(pubKey[:], pk[:65])

		return abci.Validator{
			PubKey: tmtypes.TM2PB.PubKey(pubKey),
			Power:  v.GetPower().Evaluate(),
		}
	} else {
		panic("PubKey is not of PubKeySecp256k1")
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

//_________________________________________________________________________________________________________

// XXX Audit this function further to make sure it's correct
// add tokens to a validator
func (v Validator) AddTokensFromDel(pool Pool, amount sdk.Int) (Validator, Pool, sdk.Rat) {

	// bondedShare/delegatedShare
	exRate := v.DelegatorShareExRate()
	amountDec := sdk.NewRatFromInt(amount)

	if v.Status == types.Bonded {
		pool = pool.looseTokensToBonded(amountDec)
	}

	v.Tokens = v.Tokens.Add(amountDec)
	issuedShares := amountDec.Quo(exRate)
	v.DelegatorShares = v.DelegatorShares.Add(issuedShares)

	return v, pool, issuedShares
}

// RemoveDelShares removes delegator shares from a validator.
func (v Validator) RemoveDelShares(pool Pool, delShares sdk.Rat) (Validator, Pool, sdk.Rat) {
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
func (v Validator) DelegatorShareExRate() sdk.Rat {
	if v.DelegatorShares.IsZero() {
		return sdk.OneRat()
	}
	return v.Tokens.Quo(v.DelegatorShares)
}

// Get the bonded tokens which the validator holds
func (v Validator) BondedTokens() sdk.Rat {
	if v.Status == types.Bonded {
		return v.Tokens
	}
	return sdk.ZeroRat()
}

//______________________________________________________________________

var _ types.Validator = Validator{}

// nolint - for sdk.Validator
func (v Validator) GetMoniker() string          { return v.Description.Moniker }
func (v Validator) GetStatus() types.BondStatus { return v.Status }

func (v Validator) GetOwner() sdk.Address       { return v.Owner }
func (v Validator) GetPubKey() types.PubKey     { return v.PubKey }
func (v Validator) GetPower() sdk.Rat           { return v.BondedTokens() }
func (v Validator) GetDelegatorShares() sdk.Rat { return v.DelegatorShares }
func (v Validator) GetBondHeight() int64        { return v.BondHeight }

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
