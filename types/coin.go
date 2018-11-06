package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sharering/shareledger/constants"
)

type Coin struct {
	Denom  string `json:"denom"`
	Amount Dec    `json:"amount"`
}

func NewCoin(denom string, amount int64) Coin {
	return Coin{
		Denom:  denom,
		Amount: NewDecFromInt(NewInt(amount)),
	}
}

func NewPOSCoin(amount int64) Coin {
	return Coin{
		Denom:  constants.POS_DENOM,
		Amount: NewDecFromInt(NewInt(amount)),
	}
}

func NewPOSCoinFromDec(amount Dec) Coin {
	return Coin{
		Denom:  constants.POS_DENOM,
		Amount: amount,
	}
}

func NewZeroPOSCoin() Coin {
	return NewPOSCoin(0)
}

func NewCoinFromDec(denom string, amount Dec) Coin {
	return Coin{
		Denom:  denom,
		Amount: amount,
	}
}

func NewDefaultCoin() Coin {
	return Coin{
		Denom:  constants.DEFAULT_DENOM,
		Amount: NewDecFromInt(NewInt(constants.DEFAULT_AMOUNT)),
	}
}

func (coin Coin) String() string {
	v, _ := json.Marshal(coin)
	return fmt.Sprintf("%s", v)
	//return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

func (coin Coin) Plus(other Coin) Coin {

	// If account is 0
	if coin.Amount.IsZero() {
		return other
	}

	if !coin.IsSameDenom(other) {
		return coin
	}
	return NewCoinFromDec(coin.Denom, coin.Amount.Add(other.Amount))
}

func (coin Coin) Minus(other Coin) Coin {
	if !coin.IsSameDenom(other) {
		return coin
	}
	return NewCoinFromDec(coin.Denom, coin.Amount.Sub(other.Amount))
}

func (coin Coin) Mul(factor Dec) Coin {
	return NewCoinFromDec(coin.Denom, coin.Amount.Mul(factor))
}

func (coin Coin) Quo(factor Dec) Coin {
	return NewCoinFromDec(coin.Denom, coin.Amount.Quo(factor))
}

func (coin Coin) IsSameDenom(other Coin) bool {
	return (coin.Denom == other.Denom)
}

func (coin Coin) HasValidDenom() bool {
	return IsValidDenom(coin.Denom)
}

func (coin Coin) HasDenom(denom string) bool {
	return coin.Denom == denom
}

func (c Coin) IsNil() bool         { return c.Amount.IsNil() }
func (c Coin) IsZero() bool        { return c.Amount.IsZero() }
func (c Coin) Equal(o Coin) bool   { return c.IsSameDenom(o) && c.Amount.Equal(o.Amount) }
func (c Coin) GT(o Coin) bool      { return c.IsSameDenom(o) && c.Amount.GT(o.Amount) }
func (c Coin) GTE(o Coin) bool     { return c.IsSameDenom(o) && c.Amount.GTE(o.Amount) }
func (c Coin) LT(o Coin) bool      { return c.IsSameDenom(o) && c.Amount.LT(o.Amount) }
func (c Coin) LTE(o Coin) bool     { return c.IsSameDenom(o) && c.Amount.LTE(o.Amount) }
func (c Coin) Neg() Coin           { return NewCoinFromDec(c.Denom, c.Amount.Neg()) }
func (c Coin) Abs() Coin           { return NewCoinFromDec(c.Denom, c.Amount.Abs()) }
func (c Coin) IsPositive() bool    { return c.Amount.IsPositive() }
func (c Coin) IsNotNegative() bool { return c.Amount.IsNotNegative() }

//------------------------------------------------------
// Coins

type Coins []Coin

func NewDefaultCoins() Coins {
	var ret []Coin
	for k, _ := range constants.DENOM_LIST {
		ret = append(ret, NewCoin(k, constants.DEFAULT_AMOUNT))
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
	if !other.HasValidDenom() {
		return *coins
	}
	var ret []Coin
	for _, e := range *coins {
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
	if !other.HasValidDenom() {
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

// Plus combines two sets of coins
// CONTRACT: Plus will never return Coins where one Coin has a 0 amount.
func (coins Coins) PlusMany(coinsB Coins) Coins {
	// TODO check valid denoms
	for _, c := range coinsB {
		coins = coins.Plus(c)
	}

	return coins
	// sum := ([]Coin)(nil)
	// indexA, indexB := 0, 0
	// lenA, lenB := len(coins), len(coinsB)
	// for {
	// 	if indexA == lenA {
	// 		if indexB == lenB {
	// 			return sum
	// 		}
	// 		return append(sum, coinsB[indexB:]...)
	// 	} else if indexB == lenB {
	// 		return append(sum, coins[indexA:]...)
	// 	}
	// 	coinA, coinB := coins[indexA], coinsB[indexB]
	// 	switch strings.Compare(coinA.Denom, coinB.Denom) {
	// 	case -1:
	// 		sum = append(sum, coinA)
	// 		indexA++
	// 	case 0:
	// 		if coinA.Amount.Add(coinB.Amount).IsZero() {
	// 			// ignore 0 sum coin type
	// 		} else {
	// 			sum = append(sum, coinA.Plus(coinB))
	// 		}
	// 		indexA++
	// 		indexB++
	// 	case 1:
	// 		sum = append(sum, coinB)
	// 		indexB++
	// 	}
	// }
}

// Minus subtracts a set of coins from another (adds the inverse)
func (coins Coins) MinusMany(coinsB Coins) Coins {
	fmt.Printf("MinusMany: %v\n", coinsB.Negative())
	return coins.PlusMany(coinsB.Negative())
}

// Negative returns a set of coins with all amount negative
func (coins Coins) Negative() Coins {
	res := make([]Coin, 0, len(coins))
	for _, coin := range coins {
		res = append(res, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Neg(),
		})
	}
	return res
}

func (coins Coins) Abs() Coins {
	res := make([]Coin, 0, len(coins))
	for _, coin := range coins {
		res = append(res, coin.Abs())
	}
	return res
}

// func (coins Coins) IsPositive() bool {
// 	for _, c := range coins {
// 		if !c.IsPositive() {
// 			return false
// 		}
// 	}
// 	return true
// }

// IsNotNegative - all account are not negative
// func (coins Coins) IsNotNegative() bool {
// 	for _, c := range coins {
// 		if !c.IsNotNegative() {
// 			return false
// 		}
// 	}
// 	return true
// }

func (c Coins) IsNil() bool         { return c.AllCoins("IsNil") }         // is nil. All coins are Nil
func (c Coins) IsZero() bool        { return c.AllCoins("IsZero") }        // is equal to Zero. All coins are equal to zero
func (c Coins) Equal(o Coin) bool   { return c.HasCoin(o, "Equal") }       // equal. Coin with the same denom has to be equal
func (c Coins) GT(o Coin) bool      { return c.HasCoin(o, "GT") }          // greater. Coin with the same denom has to be greater
func (c Coins) GTE(o Coin) bool     { return c.HasCoin(o, "GTE") }         // greater or equal. Coin with the same denom has to be greater or equal
func (c Coins) LT(o Coin) bool      { return c.HasCoin(o, "LT") }          // less than. Coin with the same denom has to be less than
func (c Coins) LTE(o Coin) bool     { return c.HasCoin(o, "LTE") }         // less than or equal. Coin with same denom has to be less than or equal
func (c Coins) IsPositive() bool    { return c.AllCoins("IsPositive") }    // checking positivity. All coin have to be positive
func (c Coins) IsNotNegative() bool { return c.AllCoins("IsNotNegative") } // checking negativity. All coin have to be negative

// AllCoins - ensure all coins has a field which has a method return true
// Example: coins.IsNil() == AllCoins("Amount", "IsNil")
// meaning Coins IsNil() if all coins of Coins IsNil
func (coins Coins) AllCoins(method string) bool {
	for _, c := range coins {
		if !reflect.ValueOf(c).MethodByName(method).Call([]reflect.Value{})[0].Interface().(bool) {
			return false
		}
	}
	return true
}

// HasCoin - has a coin whose method satisfies the method
// Example: coins.GT(*other*) if coins has a single coin which is GT than *other*
func (coins Coins) HasCoin(o Coin, method string) bool {
	for _, c := range coins {
		if c.IsSameDenom(o) &&
			reflect.ValueOf(c).MethodByName(method).Call([]reflect.Value{reflect.ValueOf(o)})[0].Interface().(bool) {
			return true
		}
	}
	return false
}

//--------------------------------------------------------

func IsValidDenom(denom string) bool {
	for dn, _ := range constants.DENOM_LIST {
		if dn == denom {
			return true
		}
	}
	return false

}
