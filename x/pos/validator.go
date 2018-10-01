package pos

import (
	"fmt"

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
	Owner   sdk.Address  `json:"owner"`   // sender of BondTx - UnbondTx returns here
	PubKey  types.PubKey `json:"pub_key"` // pubkey of validator
	Revoked bool         `json:"revoked"` // has the validator been revoked from bonded status?

	//PoolShares      PoolShares `json:"pool_shares"`      // total shares for tokens held in the pool
	DelegatorShares sdk.Rat `json:"delegator_shares"` // total shares issued to a validator's delegators

	Description        Description `json:"description"`           // description terms for the validator
	BondHeight         int64       `json:"bond_height"`           // earliest height as a bonded validator
	BondIntraTxCounter int16       `json:"bond_intra_tx_counter"` // block-local tx index of validator change
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
		//PoolShares:         NewUnbondedShares(sdk.ZeroRat()),
		DelegatorShares:    sdk.ZeroRat(),
		Description:        description,
		BondHeight:         int64(0),
		BondIntraTxCounter: int16(0),
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

// abci validator from stake validator type
func (v Validator) Status() types.BondStatus {
	return types.Bonded //v.PoolShares.Status
}

// XXX TEST
// get the power or potential power for a validator
// if bonded, the power is the BondedShares
// if not bonded, the power is the amount of bonded shares which the
//    the validator would have it was bonded
/*
func (v Validator) EquivalentBondedShares(pool Pool) (eqBondedShares sdk.Rat) {
	return v.PoolShares.ToBonded(pool).Amount
}
*/

//_________________________________________________________________________________________________________

// XXX Audit this function further to make sure it's correct
// add tokens to a validator
/*
func (v Validator) addTokensFromDel(pool Pool,
	amount sdk.Int) (validator2 Validator, p2 Pool, issuedDelegatorShares sdk.Rat) {

	exRate := v.DelegatorShareExRate(pool) // bshr/delshr

	var poolShares PoolShares
	var equivalentBondedShares sdk.Rat
	switch v.Status() {
	case types.Unbonded:
		pool, poolShares = pool.addTokensUnbonded(amount)
	case types.Unbonding:
		pool, poolShares = pool.addTokensUnbonding(amount)
	case types.Bonded:
		pool, poolShares = pool.addTokensBonded(amount)
	}
	v.PoolShares.Amount = v.PoolShares.Amount.Add(poolShares.Amount)
	equivalentBondedShares = poolShares.ToBonded(pool).Amount

	issuedDelegatorShares = equivalentBondedShares.Quo(exRate) // bshr/(bshr/delshr) = delshr
	v.DelegatorShares = v.DelegatorShares.Add(issuedDelegatorShares)

	return v, pool, issuedDelegatorShares
}

// remove delegator shares from a validator
// NOTE this function assumes the shares have already been updated for the validator status
func (v Validator) removeDelShares(pool Pool,
	delShares sdk.Rat) (validator2 Validator, p2 Pool, createdCoins sdk.Int) {

	amount := v.DelegatorShareExRate(pool).Mul(delShares)
	eqBondedSharesToRemove := NewBondedShares(amount)
	v.DelegatorShares = v.DelegatorShares.Sub(delShares)

	switch v.Status() {
	case types.Unbonded:
		unbondedShares := eqBondedSharesToRemove.ToUnbonded(pool).Amount
		pool, createdCoins = pool.removeSharesUnbonded(unbondedShares)
		v.PoolShares.Amount = v.PoolShares.Amount.Sub(unbondedShares)
	case types.Unbonding:
		unbondingShares := eqBondedSharesToRemove.ToUnbonding(pool).Amount
		pool, createdCoins = pool.removeSharesUnbonding(unbondingShares)
		v.PoolShares.Amount = v.PoolShares.Amount.Sub(unbondingShares)
	case types.Bonded:
		pool, createdCoins = pool.removeSharesBonded(eqBondedSharesToRemove.Amount)
		v.PoolShares.Amount = v.PoolShares.Amount.Sub(eqBondedSharesToRemove.Amount)
	}
	return v, pool, createdCoins
}

// get the exchange rate of tokens over delegator shares
// UNITS: eq-val-bonded-shares/delegator-shares
func (v Validator) DelegatorShareExRate(pool Pool) sdk.Rat {
	if v.DelegatorShares.IsZero() {
		return sdk.OneRat()
	}
	eqBondedShares := v.PoolShares.ToBonded(pool).Amount
	return eqBondedShares.Quo(v.DelegatorShares)
}
*/
//______________________________________________________________________

// ensure fulfills the sdk validator sdk
var _ types.Validator = Validator{}

// nolint - for sdk.Validator
func (v Validator) GetMoniker() string          { return v.Description.Moniker }
func (v Validator) GetStatus() types.BondStatus { return v.Status() }

func (v Validator) GetOwner() sdk.Address       { return v.Owner }
func (v Validator) GetPubKey() types.PubKey     { return v.PubKey }
func (v Validator) GetPower() sdk.Rat           { return sdk.NewRat(100) } //hard-code/*v.PoolShares.Bonded()*/ }
func (v Validator) GetDelegatorShares() sdk.Rat { return v.DelegatorShares }
func (v Validator) GetBondHeight() int64        { return v.BondHeight }

//Human Friendly pretty printer
func (v Validator) HumanReadableString() (string, error) {
	// bechOwner, err := sdk.Bech32ifyAcc(v.Owner)
	// if err != nil {
	// 	return "", err
	// }
	// bechVal, err := sdk.Bech32ifyValPub(v.PubKey)
	// if err != nil {
	// 	return "", err
	// }
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
