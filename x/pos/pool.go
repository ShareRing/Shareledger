package pos

import (
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"
)

// Pool - dynamic parameters of the current state
type Pool struct {
	LooseTokens             sdk.Rat `json:"loose_tokens"`               // tokens which are not bonded in a validator
	BondedTokens            sdk.Rat `json:"bonded_tokens"`              // reserve of bonded tokens
	DateLastCommissionReset int64   `json:"date_last_commission_reset"` // unix timestamp for last commission accounting reset (daily)

	// Fee Related
	PrevBondedShares sdk.Rat `json:"prev_bonded_shares"` // last recorded bonded shares - for fee calculations
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
		LooseTokens:             sdk.ZeroRat(),
		BondedTokens:            sdk.ZeroRat(),
		DateLastCommissionReset: 0,
		PrevBondedShares:        sdk.ZeroRat(),
	}
}

//____________________________________________________________________

// Sum total of all staking tokens in the pool
func (p Pool) TokenSupply() sdk.Rat {
	return p.LooseTokens.Add(p.BondedTokens)
}

//____________________________________________________________________

// get the bond ratio of the global state
func (p Pool) BondedRatio() sdk.Rat {
	supply := p.TokenSupply()
	if supply.GT(sdk.ZeroRat()) {
		return p.BondedTokens.Quo(supply)
	}
	return sdk.ZeroRat()
}

//_______________________________________________________________________

func (p Pool) looseTokensToBonded(bondedTokens sdk.Rat) Pool {
	p.BondedTokens = p.BondedTokens.Add(bondedTokens)
	p.LooseTokens = p.LooseTokens.Sub(bondedTokens)
	if p.LooseTokens.LT(sdk.ZeroRat()) {
		panic(fmt.Sprintf("sanity check: loose tokens negative, pool: %v", p))
	}
	return p
}

func (p Pool) bondedTokensToLoose(bondedTokens sdk.Rat) Pool {
	p.BondedTokens = p.BondedTokens.Sub(bondedTokens)
	p.LooseTokens = p.LooseTokens.Add(bondedTokens)
	if p.BondedTokens.LT(sdk.ZeroRat()) {
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
