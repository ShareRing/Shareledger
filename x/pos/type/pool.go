package posTypes

import (
	"fmt"

	"bitbucket.org/shareringvn/cosmos-sdk/wire"
	"github.com/sharering/shareledger/types"
)

// Pool - dynamic parameters of the current state
type Pool struct {
	LooseTokens             types.Dec `json:"loose_tokens"`               // tokens which are not bonded in a validator
	BondedTokens            types.Dec `json:"bonded_tokens"`              // reserve of bonded tokens
	DateLastCommissionReset int64     `json:"date_last_commission_reset"` // unix timestamp for last commission accounting reset (daily)

	// Fee Related
	PrevBondedShares types.Dec `json:"prev_bonded_shares"` // last recorded bonded shares - for fee calculations
}

/*
func (p Pool) Equal(p2 Pool) bool {
	bz1 := MsgCdc.MustMarshalBinary(&p)
	bz2 := MsgCdc.MustMarshalBinary(&p2)
	return bytes.Equal(bz1, bz2)
}*/

// initial pool for testing
func InitialPool() Pool {
	return Pool{
		LooseTokens:             types.ZeroDec(),
		BondedTokens:            types.ZeroDec(),
		DateLastCommissionReset: 0,
		PrevBondedShares:        types.ZeroDec(),
	}
}

//____________________________________________________________________

// Sum total of all staking tokens in the pool
func (p Pool) TokenSupply() types.Dec {
	return p.LooseTokens.Add(p.BondedTokens)
}

//____________________________________________________________________

// get the bond ratio of the global state
func (p Pool) BondedRatio() types.Dec {
	supply := p.TokenSupply()
	if supply.GT(types.ZeroDec()) {
		return p.BondedTokens.Quo(supply)
	}
	return types.ZeroDec()
}

//_______________________________________________________________________

func (p Pool) looseTokensToBonded(bondedTokens types.Dec) Pool {
	p.BondedTokens = p.BondedTokens.Add(bondedTokens)
	p.LooseTokens = p.LooseTokens.Sub(bondedTokens)
	if p.LooseTokens.LT(types.ZeroDec()) {
		panic(fmt.Sprintf("sanity check: loose tokens negative, pool: %v", p))
	}
	return p
}

func (p Pool) bondedTokensToLoose(bondedTokens types.Dec) Pool {
	p.BondedTokens = p.BondedTokens.Sub(bondedTokens)
	p.LooseTokens = p.LooseTokens.Add(bondedTokens)
	if p.BondedTokens.LT(types.ZeroDec()) {
		panic(fmt.Sprintf("sanity check: bonded tokens negative, pool: %v", p))
	}
	return p
}

// HumanReadableString returns a human readable string representation of a
// pool.
func (p Pool) HumanReadableString() string {

	resp := "Pool \n"
	resp += fmt.Sprintf("Loose Tokens: %s\n", p.LooseTokens)
	resp += fmt.Sprintf("Bonded Tokens: %s\n", p.BondedTokens)
	resp += fmt.Sprintf("Token Supply: %s\n", p.TokenSupply())
	resp += fmt.Sprintf("Bonded Ratio: %v\n", p.BondedRatio())
	resp += fmt.Sprintf("Date of Last Commission Reset: %d\n", p.DateLastCommissionReset)
	resp += fmt.Sprintf("Previous Bonded Shares: %v\n", p.PrevBondedShares)
	return resp
}

// unmarshal the current pool value from store key or panics
func MustUnmarshalPool(cdc *wire.Codec, value []byte) Pool {
	pool, err := UnmarshalPool(cdc, value)
	if err != nil {
		panic(err)
	}
	return pool
}

// unmarshal the current pool value from store key
func UnmarshalPool(cdc *wire.Codec, value []byte) (pool Pool, err error) {
	err = cdc.UnmarshalBinary(value, &pool)
	if err != nil {
		return
	}
	return
}
