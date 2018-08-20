package auth

import (
	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	constants "github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
)

// BaseAccount is an interface providing sequence number to avoid replay attack
// and public key for authentication
type BaseAccount interface {
	types.Account

	GetAddress() sdk.Address
	SetAddress(sdk.Address) error // errors if already set

	GetPubKey() types.PubKey // can return nil
	SetPubKey(types.PubKey) error

	GetNonce() int64
	SetNonce(int64)
	IncreaseNonce()
}

//-------------------------------------------------------
// SHRAccount

var _ BaseAccount = (*SHRAccount)(nil)

// SHRAccount - a ShareLedger account
type SHRAccount struct {
	Address sdk.Address  `json:"address"`
	Coins   types.Coins  `json:"coins"`
	PubKey  types.PubKey `json:"pub_key"`
	Nonce   int64        `json:"nonce"`
}

// NewSHRAccountWithAddress create  a SHRAccount with address
func NewSHRAccountWithAddress(addr sdk.Address) SHRAccount {
	return SHRAccount{
		Address: addr,
	}
}

// Implement BaseAccount interface

func (acc SHRAccount) GetAddress() sdk.Address {
	return acc.Address
}

func (acc *SHRAccount) SetAddress(addr sdk.Address) {
	if len(acc.Address) != 0 {
		return errors.New(constants.SHRACCOUNT_EXISTING_ADDRESS)
	}
	acc.Address = addr
	return nil
}

func (acc SHRAccount) GetPubKey() types.PubKey {
	return acc.PubKey
}

func (acc *SHRAccount) SetPubKey(pk types.PubKey) {
	acc.PubKey = pk
	return nil
}

func (acc SHRAccount) GetNonce() int64 {
	return acc.Nonce
}

func (acc *SHRAccount) SetNonce(no int64) {
	acc.Nonce = no
	return nil
}

func (acc *SHRAccount) IncreaseNonce() {
	acc.Nonce += 1
	return nil
}
