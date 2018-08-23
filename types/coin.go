package types

import (
	"encoding/json"
	"fmt"
	"github.com/sharering/shareledger/constants"
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
	v, _ := json.Marshal(coin)
	return fmt.Sprintf("%s", v)
	//return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
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

func (coin Coin) HasValidDenoms() bool {
	for denom, _ := range constants.DENOM_LIST {
		if coin.Denom == denom {
			return true
		}
	}
	return false
}

//------------------------------------------------------
// Coins

type Coins []Coin

func NewDefaultCoins() Coins {
	var ret []Coin
	for k, _ := range constants.DENOM_LIST {
		ret = append(ret, NewCoin(k, 0))
	}
	return ret
}

func (coins Coins) String() string {
	//if len(coins) == 0 {
	//return ""
	//}

	//out := "{"
	//for _, coin := range coins {
	//out += fmt.Sprintf("%v,", coin.String())
	//}
	//return out[:len(out)-1]
	if v, err := json.Marshal(coins); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", v)
	}
}

func (coins Coins) GetCoins() Coins {
	return coins
}

func (coins *Coins) SetCoins(co Coins) {
	*coins = co
}

func (coins Coins) HasValidDenoms() bool {
	if len(coins) != len(constants.DENOM_LIST) {
		return false
	}
	checked := make(map[string]bool)
	for _, c := range coins {
		if constants.DENOM_LIST[c.Denom] && !checked[c.Denom] {
			checked[c.Denom] = true
		} else {
			return false
		}
	}
	return true
}

func (coins *Coins) Plus(other Coin) Coins {
	if !other.HasValidDenoms() {
		return *coins
	}
	var ret []Coin
	for _, e := range *coins {
		fmt.Println(e)
		if e.IsSameDenom(other) {
			ret = append(ret, e.Plus(other))
		} else {
			ret = append(ret, e)
		}
	}
	return ret
}

//func (coins *Coins) PlusCoins(others Coins) Coins {
//ret := Coins(*coins)
//for _, e := range others {
//ret = ret.Plus(e)
//}
//return ret
//}

// Minus - ensure *coins* and *co* have valid denoms
func (coins *Coins) Minus(other Coin) Coins {
	if !other.HasValidDenoms() {
		return *coins
	}
	var ret []Coin
	for _, e := range *coins {
		if e.IsSameDenom(other) {
			ret = append(ret, e.Minus(other))
		} else {
			ret = append(ret, e)
		}
	}

	return ret
}

// IsPositive - all account has positive
func (coins Coins) IsPositive() bool {
	for _, c := range coins {
		if !c.IsPositive() {
			return false
		}
	}
	return true
}

// IsNotNegative - all account are not negative
func (coins Coins) IsNotNegative() bool {
	for _, c := range coins {
		if !c.IsNotNegative() {
			return false
		}
	}
	return true
}
