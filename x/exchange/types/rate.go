package types

import (
	"fmt"
	"encoding/json"

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
func (e ExchangeRate) Convert(sellingCoin types.Coin) (buyingCoin types.Coin) {
	if !sellingCoin.HasDenom(e.FromDenom) {
		panic(fmt.Sprintf(constants.EXC_INVALID_DENOM, e.FromDenom, sellingCoin.Denom))
	}

	// if convert FromDenom -> ToDenom
	buyingAmount := sellingCoin.Amount.Mul(e.Rate)

	return types.NewCoinFromDec(e.ToDenom, buyingAmount)
}

func (e ExchangeRate) Obtain(buyingCoin types.Coin) (sellingCoin types.Coin) {
	if !buyingCoin.HasDenom(e.ToDenom) {
		panic(fmt.Sprintf(constants.EXC_INVALID_DENOM, e.ToDenom, buyingCoin.Denom))
	}

	// if convert ToDemom -> FromDenom
	sellingAmount := buyingCoin.Amount.Quo(e.Rate)

	return types.NewCoinFromDec(e.FromDenom, sellingAmount)
}

func (e ExchangeRate) String() string{
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", b) 
}
