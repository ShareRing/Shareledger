package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"bitbucket.org/shareringvn/cosmos-sdk/wire"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/auth"
	"github.com/sharering/shareledger/x/bank"
	bmsg "github.com/sharering/shareledger/x/bank/messages"

	"github.com/sharering/shareledger/cmd/stress-test/accounts"
	"github.com/sharering/shareledger/cmd/stress-test/constants"
	"github.com/sharering/shareledger/cmd/stress-test/requests"
	"github.com/sharering/shareledger/cmd/stress-test/utils"
)

const (
	// InitialBalance initial load for each account
	InitialBalance int64 = 10000
	// TransferValue each send
	TransferValue int64 = 1
	// Denom denominator
	Denom string = "SHRP"
	// temp file
	TempFile string = "accounts.tmp"

	RateLimit int = 2000
)

func main() {
	// cdc := app.MakeCodec()

	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "stresstest",
		Short: "Perform stress test on ShareLedger",
	}

	// ------- account command
	// generate accounts to prepare for stress test

	accountCmd := &cobra.Command{
		Use:   "account",
		Short: "Generate accounts to prepare for stress test",
		RunE:  execute,
	}

	accountCmd.Flags().Int(constants.FlagN, 10, "Number of account to create")

	rootCmd.AddCommand(accountCmd)

	rootCmd.Execute()
}

func check(err error) {
	utils.Check(err)
}

func execute(cmd *cobra.Command, args []string) error {
	// testNet()
	createAccount(cmd)
	spamSendTx(cmd)
	return nil
}

func createAccount(cmd *cobra.Command) {
	number, err := cmd.Flags().GetInt(constants.FlagN)

	fmt.Printf("Number: %d\n", number)

	if err != nil {
		panic(err)
	}

	cdc := getCodec()
	nonces := map[string]int64{} // mapping of all current nonce

	privKeys := make([]types.PrivKeySecp256k1, 0)
	pubKeys := make([]types.PubKeySecp256k1, 0)

	for i := 0; i < number; i++ {

		pubKey, privKey := accounts.GetKeyPair()

		nonce := nonces[pubKey.Address().String()] + 1

		nonces[pubKey.Address().String()] = nonce

		// fmt.Println(privKey, pubKey)

		// Get a message
		loadMsg := bmsg.NewMsgLoad(pubKey.Address(),
			types.NewCoin(Denom, InitialBalance))

		authTx := encodeMsg(cdc,
			privKey,
			pubKey,
			loadMsg,
			nonce)

		requests.SendTx(authTx)

		privKeys = append(privKeys, privKey)
		pubKeys = append(pubKeys, pubKey)
	}

	// Write to File
	utils.WriteFile(TempFile, pubKeys, privKeys)
}

func spamSendTx(cmd *cobra.Command) {

	cdc := getCodec()

	fmt.Printf("ReadFile... \n")
	// Read File
	pubKeys, privKeys := utils.ReadFile(TempFile)

	// Wait between send RateLimit milisecondj
	limiter := time.Tick(time.Duration(RateLimit) * time.Millisecond)

	for {
		<-limiter
		fmt.Println("Executing goroutine")
		go func() {

			fromIdx := rand.Intn(len(pubKeys))
			toIdx := rand.Intn(len(pubKeys))

			// Get Nonce
			nonce, err := requests.QueryNonceTx(nonceQuery(cdc, pubKeys[fromIdx].Address()))
			check(err)

			// Send Tx

			sendMsg := getSendTransaction(pubKeys[toIdx], TransferValue)
			sendTx := encodeMsg(cdc, privKeys[fromIdx], pubKeys[fromIdx], sendMsg, nonce + 1)

			start := time.Now()

			res, err := requests.SendTx(sendTx)
			fmt.Printf("Result: %s\n", res)

			elapsed := time.Now().Sub(start)

			check(err)

			if err != nil {
				fmt.Println("From: %s, To: %s, Amount: %d", pubKeys[fromIdx].Address(),
					pubKeys[toIdx].Address(),
					TransferValue)
			}

		}()
	}

}

// ---------------- UTILITIES -----------------------

func encodeMsg(cdc *wire.Codec, privKey types.PrivKey, pubKey types.PubKey, msg sdk.Msg, nonce int64) string {

	// Sign Transaction
	authTx := auth.GetAuthTx(pubKey,
		privKey,
		msg,
		nonce)

	// Amino encoding
	tx, err := cdc.MarshalBinary(authTx)
	if err != nil {
		panic(err)
	}

	// Encode as a string
	return "0x" + hex.EncodeToString(tx)
}

func encodeQuery(cdc *wire.Codec, queryTx types.SHRTx) string {
	// Amino encoding
	tx, err := cdc.MarshalBinary(queryTx)

	if err != nil {
		panic(err)
	}

	// Encode as a string
	return "0x" + hex.EncodeToString(tx)
}

func getCodec() *wire.Codec {
	cdc := app.MakeCodec()

	cdc = bank.RegisterCodec(cdc)
	cdc = auth.RegisterCodec(cdc)

	return cdc
}

func testNet() {
	cdc := getCodec()
	pubKey, _ := accounts.GetKeyPair()
	// requests.SendTx()

	balances, _ := requests.QueryBalanceTx(balanceQuery(cdc, pubKey.Address()))

	fmt.Println("Balance of SHRP:", balances["SHRP"])

}

func balanceQuery(cdc *wire.Codec, address sdk.Address) string {
	checkMsg := bmsg.NewMsgCheck(address)

	queryTx := types.NewQueryTx(checkMsg)

	return encodeQuery(cdc, queryTx)

}

func nonceQuery(cdc *wire.Codec, address sdk.Address) string {
	nonceMsg := auth.NewMsgNonce(address)

	queryTx := types.NewQueryTx(nonceMsg)

	return encodeQuery(cdc, queryTx)
}

func getSendTransaction(pubKey types.PubKey, amount int64) sdk.Msg {
	coin := types.NewCoin(Denom, amount)
	return bmsg.NewMsgSend(pubKey.Address(), coin)
}
