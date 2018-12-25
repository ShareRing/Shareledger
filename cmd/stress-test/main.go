package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

var (
	// InitialBalance initial load for each account
	InitialBalance int64 = 10000
	// TransferValue each send
	TransferValue int64 = 1
	// Denom denominator
	Denom string = "SHRP"
	// temp file
	TempFile string = "accounts.tmp"

	RateLimit int = 2000

	MaximumRequest int = 2

	BurstRequest = 3
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

	throttleCmd := &cobra.Command{
		Use:   "throttle",
		Short: "Throttle Tendermint with maximum number of request in burst after each interval",
		RunE:  spamSendTx,
	}

	throttleCmd.Flags().Int(constants.IntervalFlag, RateLimit, "Interval between bust")
	throttleCmd.Flags().Int(constants.BurstRequestFlag, BurstRequest, "Number of requests per burst")
	throttleCmd.Flags().Int(constants.MaximumRequestFlag, MaximumRequest, "Number of bursts")

	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(throttleCmd)

	rootCmd.Execute()
}

func check(err error) {
	utils.Check(err)
}

func execute(cmd *cobra.Command, args []string) error {
	// testNet()
	createAccount(cmd)
	// spamSendTx(cmd)
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

func spamSendTx(cmd *cobra.Command, in []string) error {

	cdc := getCodec()

	maximumRequest, _ := cmd.Flags().GetInt(constants.MaximumRequestFlag)
	burstRequest, _ := cmd.Flags().GetInt(constants.BurstRequestFlag)
	interval, _ := cmd.Flags().GetInt(constants.IntervalFlag)

	fmt.Printf("ReadFile... \n")
	// Read File
	pubKeys, privKeys := utils.ReadFile(TempFile)

	// Wait between send RateLimit milisecondj
	limiter := time.NewTicker(time.Duration(interval) * time.Millisecond)

	statistics := make(chan time.Duration)

	go collectStatistics(statistics)

	var wg sync.WaitGroup

	count := 0

	for {
		if count < maximumRequest {
			<-limiter.C
			fmt.Println("Executing goroutine")

			wg.Add(burstRequest)

			for i := 0; i < burstRequest; i++ {
				go Spam(&wg, pubKeys, privKeys, cdc, statistics)
			}

			count++

			if count == maximumRequest {
				limiter.Stop()
			}

		} else {

			break

		}
	}
	wg.Wait()
	close(statistics)
	time.Sleep(2 * time.Second)

	return nil
}

// --------------------- GOROUTINE -------
func Spam(wg *sync.WaitGroup,
	pubKeys []types.PubKeySecp256k1,
	privKeys []types.PrivKeySecp256k1,
	cdc *amino.Codec,
	stat chan time.Duration) {

	fromIdx := rand.Intn(len(pubKeys))
	toIdx := rand.Intn(len(pubKeys))

	// Get Nonce
	nonce, err := requests.QueryNonceTx(nonceQuery(cdc, pubKeys[fromIdx].Address()))
	check(err)

	// Send Tx

	sendMsg := getSendTransaction(pubKeys[toIdx], TransferValue)
	sendTx := encodeMsg(cdc, privKeys[fromIdx], pubKeys[fromIdx], sendMsg, nonce+1)

	start := time.Now()

	_, err = requests.SendTx(sendTx)
	// fmt.Printf("Result: %s\n", res)

	elapsed := time.Now().Sub(start)

	check(err)

	fmt.Printf("From: %s, To: %s, Amount: %d\n", pubKeys[fromIdx].Address(),
		pubKeys[toIdx].Address(),
		TransferValue)

	fmt.Printf("ElapsedTime: %d\n", elapsed)

	stat <- elapsed
	wg.Done()
}

// ---------------- STATISTICS ---------------------

type Statistics struct {
	Count int64 `json:"count"`
	Total int64 `json:"total"`
	Max   int64 `json:"max"`
}

func collectStatistics(received chan time.Duration) {
	stat := Statistics{0, 0, 0}

	for dur := range received {

		fmt.Println("Stat: received", dur)

		nano := dur.Nanoseconds()

		stat.Total += nano

		if stat.Max < nano {
			stat.Max = nano
		}

		stat.Count++
	}

	fmt.Printf("Statistics: %v\n", stat)
	fmt.Printf("%20s %20s %20s %20s\n", "Count", "Total(ms)", "Max(ms)", "Average(ms)")
	fmt.Printf("%20d %20f %20f %20f\n", stat.Count,
		float64(stat.Total)/math.Pow(10, 6),
		float64(stat.Max)/math.Pow(10, 6),
		(float64(stat.Total)/math.Pow(10, 6))/float64(stat.Count))
}

// ---------------- UTILITIES -----------------------

func encodeMsg(cdc *amino.Codec, privKey types.PrivKey, pubKey types.PubKey, msg sdk.Msg, nonce int64) string {

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

func encodeQuery(cdc *amino.Codec, queryTx types.SHRTx) string {
	// Amino encoding
	tx, err := cdc.MarshalBinary(queryTx)

	if err != nil {
		panic(err)
	}

	// Encode as a string
	return "0x" + hex.EncodeToString(tx)
}

func getCodec() *amino.Codec {
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

func balanceQuery(cdc *amino.Codec, address sdk.AccAddress) string {
	checkMsg := bmsg.NewMsgCheck(address)

	queryTx := types.NewQueryTx(checkMsg)

	return encodeQuery(cdc, queryTx)

}

func nonceQuery(cdc *amino.Codec, address sdk.AccAddress) string {
	nonceMsg := auth.NewMsgNonce(address)

	queryTx := types.NewQueryTx(nonceMsg)

	return encodeQuery(cdc, queryTx)
}

func getSendTransaction(pubKey types.PubKey, amount int64) sdk.Msg {
	coin := types.NewCoin(Denom, amount)
	return bmsg.NewMsgSend(pubKey.Address(), coin)
}
