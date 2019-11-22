package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/sharering/shareledger/types"
)

const (
	// clientT  string = "tcp://127.0.0.1:46657"
	clientT  string = "tcp://192.168.1.234:46657"
	reserveT string = "ab83994cf95abe45b9d8610524b3f8f8fd023d69f79449011cb5320d2ca180c5"
	privKeyT string = "d69bd8078db0b6de71ae4b41953d24b913e4a90f3f58892db4d5d3c447c86ed2"
	addressT string = "9FFF9DF9813FA36F99C03FDE288BE285046402A7"
	amountT  string = "12.34"
	denomT   string = "SHRP"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		fmt.Printf(err.Error())
		t.Errorf(err.Error())
		os.Exit(1)
	}
}

// TestFullFlow - run the following scenario
// 1. Check balance
// 2. Load  balance x coins
// 3. Check balance then check whether new balance is added x coins
// 4. Generate randomly a new account
// 5. Send x coins to newly created account
// 6. Check balance after sending, whether balance returns to the original amount
func TestFullFlow(t *testing.T) {

	//-----------------------------------------
	// Check initial balance

	balanceR, err := CheckBalanceAPI(
		clientT,
		privKeyT,
		addressT,
	)
	checkError(err, t)
	fmt.Printf("Original balance: %v\n\n", balanceR)

	// Two ways to get types.Dec from string
	// First

	// var originalBalance types.Dec
	// originalBalance, err := originalBalance.UnmarshalAmino(balanceR.SHRP)

	// Second
	originalBalance, err := types.NewDecFromStr(balanceR.SHRP)

	checkError(err, t)

	fmt.Printf("Original balance in SHR: %v\n", originalBalance)

	//-----------------------------------------
	// load some coins

	loadR, err := LoadBalanceAPI(
		clientT,
		reserveT, // NOTE: loading use reserve private key, not account's priv key
		addressT,
		amountT,
		denomT,
	)

	checkError(err, t)
	fmt.Printf("Load balance result: \n %v\n\n", loadR)

	//-----------------------------------------
	// check whether load is successful

	balanceR, err = CheckBalanceAPI(
		clientT,
		privKeyT,
		addressT,
	)

	checkError(err, t)
	fmt.Printf("Current balance: %v\n\n", balanceR)

	// var newBalance types.Dec
	// err = newBalance.UnmarshalAmino(balanceR.SHRP)
	// checkError(err, t)
	newBalance, err := types.NewDecFromStr(balanceR.SHRP)
	checkError(err, t)

	var amountDec types.Dec
	err = amountDec.UnmarshalAmino(amountT)
	checkError(err, t)

	delta := newBalance.Sub(originalBalance)
	if !delta.Equal(amountDec) {
		t.Errorf("Balance after loaded is incorrect. Got: %v Expected: %v", delta, amountDec)
	}

	//-----------------------------------------
	// generate account randomly

	privKey, address := GenerateAccountAPI()

	fmt.Printf("Generate account: privKey=%s address=%s\n\n", privKey,
		address)

	//-----------------------------------------
	// Send to newly generated account
	// exchange should have
	sendR, err := SendCoinAPI(
		clientT,
		privKeyT,
		address,
		amountT,
		denomT,
	)

	checkError(err, t)
	fmt.Printf("Send coin: %v\n\n", sendR)

	//-----------------------------------------
	// check balance after sending
	balanceR, err = CheckBalanceAPI(
		clientT,
		privKeyT,
		addressT,
	)

	checkError(err, t)
	fmt.Printf("Current balance: %v\n\n", balanceR)

	var finalBalance types.Dec
	err = finalBalance.UnmarshalAmino(balanceR.SHRP)
	checkError(err, t)
	if !originalBalance.Equal(finalBalance) {
		t.Errorf("After sending, balance should revert to original one. Original: %v Final: %v", originalBalance, finalBalance)
	}

}
