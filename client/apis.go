package client

import (
	"encoding/hex"
	"fmt"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
)

// LoadBalanceAPI - load coin
func LoadBalanceAPI(
	client string, // node URL Ex: tcp://192.168.1.234:46657
	privKey string, // private key of faucet/reserve account
	toAddress string, // receiving address
	amount string, // amount, decimal is possible
	denom string, // SHR or SHRP
) (res Response, err error) {

	// convert string to sdk.Address
	addressBytes, err := hex.DecodeString(toAddress)
	if err != nil {
		return res, err
	}

	addr := sdk.Address(addressBytes)

	dec, err := types.NewDecFromStr(amount)

	if err != nil {
		return res, err
	}

	coin := types.NewCoinFromDec(denom, dec)

	context := NewCoreContextWithClient(privKey, client)

	return context.LoadBalance(addr, coin)
}

func SendCoinAPI(
	client string, // node URL Ex: tcp://192.168.1.234:46657
	privKey string, // *private key* of the sender
	toAddress string, // *address* of the receiver
	amount string, // amount, decimal is possible Ex: 1.23
	denom string,
) (res Response, err error) {

	// convert string to sdk.Address
	addressBytes, err := hex.DecodeString(toAddress)
	if err != nil {
		return res, err
	}

	addr := sdk.Address(addressBytes)

	dec, err := types.NewDecFromStr(amount)

	if err != nil {
		return res, err
	}

	coin := types.NewCoinFromDec(denom, dec)

	context := NewCoreContextWithClient(privKey, client)

	return context.SendCoins(addr, coin)
}

// NOTE
// types.Coins = []types.Coin
// types.Coin = {
// 	 Denom: string,
//   Amount: types.Dec,   // NOTE: custom typed
// }

type Balance struct {
	SHR  string // amount of SHR coins
	SHRP string // amount of SHRP coins
}

// CheckBalanceAPI - return balance of an address
func CheckBalanceAPI(
	client string, // node URL Ex: tcp://192.168.1.234:46657
	privKey string, // *private key* of any account
	address string, // account to be checked balance
) (balance Balance, err error) {

	// convert string to sdk.Address
	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		return balance, err
	}

	addr := sdk.Address(addressBytes)

	context := NewCoreContextWithClient(privKey, client)

	coins, err := context.CheckBalance(addr)

	if err != nil {
		return balance, err
	}
	for _, coin := range coins {
		if coin.Denom == "SHR" {
			// shr, err := coin.Amount.MarshalAmino()
			// if err != nil {
			// 	return balance, err
			// }
			// balance.SHR = fmt.Sprintf("%s", shr)
			balance.SHR = coin.Amount.String()
		}
		if coin.Denom == "SHRP" {
			// shrp, err := coin.Amount.MarshalAmino()
			// if err != nil {
			// 	return balance, err
			// }
			// balance.SHRP = fmt.Sprintf("%s", shrp)
			balance.SHRP = coin.Amount.String()
		}
	}

	return balance, nil
}

// GenerateAccountAPI - randomly generate account
func GenerateAccountAPI() (privKey string, address string) {
	pubK, privK := types.GenerateKeyPair()

	address = fmt.Sprintf("%X", pubK.Address())

	return fmt.Sprintf("%x", privK[:]), address
}
