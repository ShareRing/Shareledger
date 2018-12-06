package client

import (
	"fmt"
	"testing"
)

const (
	clientT  string = "tcp://192.168.1.234:46657"
	reserveT string = "ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5"
	privKeyT string = "d69bd8078db0b6de71ae4b41953d24b913e4a90f3f58892db4d5d3c447c86ed2"
	addressT string = "9FFF9DF9813FA36F99C03FDE288BE285046402A7"
	amountT  string = "15.23"
	denomT   string = "SHRP"
)

func TestFullFlow(*testing.T) {

	// load some coins
	loadR, err := LoadBalanceAPI(
		clientT,
		reserveT, // NOTE: loading use reserve private key, not account's priv key
		addressT,
		amountT,
		denomT,
	)

	fmt.Printf("Load balance: %v\nError: %v\n\n", loadR, err)

	// check whether load is successful
	balanceR, err := CheckBalanceAPI(
		clientT,
		privKeyT,
		addressT,
	)

	fmt.Printf("Current balance: %v\nError: %v\n\n", balanceR, err)

	// generate account randomly
	privKey, address := GenerateAccountAPI()

	fmt.Printf("Generate account: privKey=%s address=%s\n\n", privKey,
		address)

	// Send to newly generated account
	sendR, err := SendCoinAPI(
		clientT,
		privKeyT,
		address,
		amountT,
		denomT,
	)

	fmt.Printf("Send coin: %v\nError: %v\n\n", sendR, err)

	// check balance after sending
	balanceR, err = CheckBalanceAPI(
		clientT,
		privKeyT,
		addressT,
	)
	fmt.Printf("Current balance: %v\nError: %v\n", balanceR, err)
}
