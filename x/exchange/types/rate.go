package types

import (
	"fmt"

	"github.com/sharering/shareledger/constants"
	"github.com/sharering/shareledger/types"
)

// ExchangeRate - struct to hold exchange rate
type ExchangeRate struct {
	FromDenom string    `json:"fromDenom"`
	ToDenom   string    `json:"toDenom"`
	Rate      types.Dec `json:"rate"` // FromDenom = ToDenom * Rate
}

// NewExchangeRate - new ExchangeRate
func NewExchangeRate(
	from string,
	to string,
	rate types.Dec,
) ExchangeRate {
	return ExchangeRate{
		FromDenom: from,
		ToDenom:   to,
		Rate:      rate,
	}
}

// UpdateRate - update rate of this exchange rate
func (e ExchangeRate) UpdateRate(newRate types.Dec) ExchangeRate {
	e.Rate = newRate
	return e
}

// Convert - convert from one coin to another using this exchange rate
func (e ExchangeRate) Convert(from types.Coin) (to types.Coin) {
	if !from.HasDenom(e.FromDenom) {
		panic(fmt.Sprintf(constants.EXC_INVALID_DENOM, e.FromDenom, from.Denom))
	}

	toAmount := from.Amount.Mul(e.Rate)
	return types.NewCoinFromDec(e.ToDenom, toAmount)
}
