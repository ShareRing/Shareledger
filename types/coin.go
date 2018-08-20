package types

import (
	"fmt"
)

type Coin struct {
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

func NewCoin(denom string, amount int64) Coin {
	return Coin{
		Denom:  denom,
		Amount: amount,
	}
}

func (coin Coin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

func (coin Coin) Plus(other Coin) Coin {

	// If account is 0
	if coin.Amount == 0 {
		return other
	}

	if !coin.IsSameDenom(other) {
		return coin
	}
	return NewCoin(coin.Denom, coin.Amount+other.Amount)
}

func (coin Coin) Minus(other Coin) Coin {
	if !coin.IsSameDenom(other) {
		return coin
	}
	return NewCoin(coin.Denom, coin.Amount-other.Amount)
}

func (coin Coin) IsSameDenom(other Coin) bool {
	return (coin.Denom == other.Denom)
}

func (coin Coin) IsPositive() bool {
	return coin.Amount > 0
}

func (coin Coin) IsNotNegative() bool {
	return coin.Amount >= 0
}

//------------------------------------------------------
// Coins

type Coins []Coin

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}

func (coins Coins) GetCoins() Coins {
	return coins
}

func (coins *Coins) SetCoins(co Coins) {
	*coins = co
}
