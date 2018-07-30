package types

import (
	"github.com/sharering/shareledger/constants"
)

// Simple account struct
type AppAccount struct {
	Coins Coin `json:"coins"`
}


func NewDefaultAccount() AppAccount {
	return AppAccount{
		Coins: NewCoin(constants.DEFAULT_DENOM, constants.DEFAULT_AMOUNT),
	}
}
