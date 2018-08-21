package types

import (
	"github.com/sharering/shareledger/constants"
)

type Account interface {
	GetCoins() Coins
	SetCoins(Coins)
}

//----------------------------------------------
// AppAccount

var _ Account = (*AppAccount)(nil)

// Simple account struct
type AppAccount struct {
	Coins Coin `json:"coins"`
}

func NewDefaultAccount() AppAccount {
	return AppAccount{
		Coins: NewCoin(constants.DEFAULT_DENOM, constants.DEFAULT_AMOUNT),
	}
}

func (acc AppAccount) GetCoins() Coins {
	return Coins([]Coin{acc.Coins})
}

func (acc *AppAccount) SetCoins(coins Coins) {
	acc.Coins = coins[0]
}
