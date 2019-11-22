package client

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sharering/shareledger/types"
)

// LoadBalanceAPI - load coin
func LoadBalanceAPI(
	client string, // node URL Ex: tcp://192.168.1.234:46657
	privKey string, // private key of faucet/reserve account
	toAddress string, // receiving address
	amount string, // amount, decimal is possible
	denom string, // SHR or SHRP
) (err error) {

	// convert string to sdk.AccAddress
	addressBytes, err := hex.DecodeString(toAddress)
	if err != nil {
		return err
	}

	addr := sdk.AccAddress(addressBytes)

	dec, err := types.NewDecFromStr(amount)

	if err != nil {
		return err
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
) (err error) {

	// convert string to sdk.AccAddress
	addressBytes, err := hex.DecodeString(toAddress)
	if err != nil {
		return err
	}

	addr := sdk.AccAddress(addressBytes)

	dec, err := types.NewDecFromStr(amount)

	if err != nil {
		return err
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

// GenerateAccountAPI - randomly generate account
func GenerateAccountAPI() (privKey string, address string) {
	pubK, privK := types.GenerateKeyPair()

	address = fmt.Sprintf("%X", pubK.Address())

	return fmt.Sprintf("%x", privK[:]), address
}
