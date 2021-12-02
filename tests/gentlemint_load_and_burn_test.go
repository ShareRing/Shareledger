package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/keeper"
	"github.com/stretchr/testify/require"
)

const (
	ShareLedgerSuccessCode = uint32(0)
	ShareLedgerErrorCodeUnauthorized = uint32(4)
	ShareLedgerErrorCodeInvalidCoin    = uint32(10)
	ShareLedgerErrorCodeInvalidRequest = uint32(18)

	ShareLedgerBookingAssetAlreadyBooked = uint32(3)
	ShareLedgerBookingBookerIsNotOwner = uint32(6)

	ShareLedgerErrorMessageInvalidCoinCoinIsMaximum = "SHR possible mint exceeded"
	ShareLedgerErrorMessageInvalidCoin              = "Amount must be positive"
	ShareLedgerErrorMessageUnauthorized              = "Amount must be positive"
)

func TestGentlemint_LoadSHR_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)
	amount := "10"

	loadCmd := fmt.Sprintf("load-shr %s %s", treasurer.String(), amount)

	// Load
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

func TestGentlemint_LoadSHR_ApproveAddressIsNotAuthority(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	notAuthorityAddress := f.KeyAddress(keyUser4)
	treasurer := f.KeyAddress(keyTreasurer)
	amount := "10"

	amountBeforeLoad := shrAmountOfAddress(treasurer, f)

	loadCmd := fmt.Sprintf("load-shr %s %s", treasurer.String(), amount)

	// Load
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", notAuthorityAddress.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, ShareLedgerErrorCodeUnauthorized, txResponse.Code)
	t.Logf("the transaction message after load coin %s ", txResponse)
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	amountAfterLoad := shrAmountOfAddress(treasurer, f)

	require.Equal(t, amountBeforeLoad.String(), amountAfterLoad.String())

	f.Cleanup()
}

//TestGentlemint_LoadSHR_AmountMustBePositiveNumber impossible to set the number of
func TestGentlemint_LoadSHR_AmountMustBNotEqualZero(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)
	amount := "0"

	loadCmd := fmt.Sprintf("load-shr %s %s", treasurer.String(), amount)

	// Load
	_, _, stdErr := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)


	require.Contains(t,stdErr, ShareLedgerErrorMessageInvalidCoin)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

//TestGentlemint_LoadSHR_AmountMustBePositiveNumber impossible to set the number of
func TestGentlemint_LoadSHR_AmountCannotExceedMaximumSHRCoin(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	//getting the total SHR coin in our chain and
	//callculating the amount can't make the total coin exceed maxximum

	totalCurrentSHR := f.QueryTotalSupplyOf("shr")

	amount := sdk.NewInt(keeper.MaxSHRSupply).Sub(totalCurrentSHR).Add(sdk.NewInt(10))

	loadCmd := fmt.Sprintf("load-shr %s %s", treasurer.String(), amount.String())

	// Load
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, ShareLedgerErrorCodeInvalidRequest, txResponse.Code)
	require.Contains(t, txResponse.RawLog, ShareLedgerErrorMessageInvalidCoinCoinIsMaximum)
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

func TestGentlemint_BurnSHR_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)
	amount := "10"
	burnCmd := fmt.Sprintf("burn-shr %s", amount)
	loadCmd := fmt.Sprintf("load-shr %s %s", treasurer.String(), amount)

	// Load
	f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.ExecuteGentlemintTxCommand(burnCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

func TestGentlemint_Enroll_ShrpLoader_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())

	// Load
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

func TestGentlemint_Revoke_ShrpLoader_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())
	revokeLoaderCmd := fmt.Sprintf("revoke-loaders %s", treasurer.String())

	// Enroll
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Revoke
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(revokeLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.Cleanup()
}

func TestGentlemint_LoadSHRP_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	emptyAccount := f.KeyAddress(keyEmtyUser)
	amount := "10"

	loadCmd := fmt.Sprintf("load-shrp %s %s", emptyAccount.String(), amount)
	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())

	// Enroll
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	balance := f.QueryAccount(emptyAccount)

	require.Equal(t, "9", balance.GetCoins().AmountOf("shrp").String())
	require.Equal(t, "9", balance.GetCoins().AmountOf("shr").String())

	f.Cleanup()
}

func TestGentlemint_BurnSHRP_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	// emptyAccount := f.KeyAddress(keyEmtyUser)
	amount := "10"
	burnAmount := "1"

	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())
	loadCmd := fmt.Sprintf("load-shrp %s %s", treasurer.String(), amount)
	burnCmd := fmt.Sprintf("burn-shrp %s", burnAmount)

	// Enroll
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// Burn
	_, stdOut3, _ := f.ExecuteGentlemintTxCommand(burnCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse3 := ParseStdOut(t, stdOut3)
	require.Equal(t, uint32(0), txRepsonse3.Code)

	f.Cleanup()
}

func TestGentlemint_SetExchangeRate_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use

	treasurer := f.KeyAddress(keyTreasurer)

	rate := "99"
	setExchangeRateCmd := fmt.Sprintf("set-exchange %s", rate)

	_, stdOut, _ := f.ExecuteGentlemintTxCommand(setExchangeRateCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txResponse.Code)

	cmd := fmt.Sprintf("%s query gentlemint get-exchange %v", f.GaiacliBinary, f.Flags())
	flags := []string{}
	exchangeRate, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	require.Equal(t, rate, strings.Replace(exchangeRate, "\"", "", -1))

	f.Cleanup()
}


func TestGentlemint_SetExchangeRate_AndBuySHR(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use


	//First buy with default exchange rate
	treasurer := f.KeyAddress(keyTreasurer)
	authority := f.KeyAddress(keyAuthority)
	userNormal := f.KeyAddress(keyUser4)
	currentRate := f.QueryExchangeRate(treasurer)
	// Exchange rate is 200 by default. It's mean to buy 1SHR we need to spend 1/200 ~ 0.005
	t.Logf("stating to buy SHR first time with exchange rate %s It's mean to buy 1SHR we need to spend 1/200 ~ 0.005",currentRate)
	account := f.QueryAccount(userNormal)
	accountCoinInfo := account.GetCoins()

	t.Logf("Default Amount SHR %v",accountCoinInfo.AmountOf("shr"))
	t.Logf("Default Amount SHRP %v",accountCoinInfo.AmountOf("shrp"))
	t.Logf("Default Amount Cent %v",accountCoinInfo.AmountOf("cent"))
	t.Logf("now we starting to test to set exchange rate to 99 It's mean to buy 1SHR we need to spend 1/99 ~ 0.01010101")

	rate := "99"
	setExchangeRateCmd := fmt.Sprintf("set-exchange %s", rate)

	_, stdOut, _ := f.ExecuteGentlemintTxCommand(setExchangeRateCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txResponse.Code)

	cmd := fmt.Sprintf("%s query gentlemint get-exchange %v", f.GaiacliBinary, f.Flags())

	exchangeRate, _ := tests.ExecuteT(f.T, addFlags(cmd, []string{}), "")

	require.Equal(t, rate, strings.Replace(exchangeRate, "\"", "", -1))

	tests.WaitForNextNBlocksTM(1,f.Port)

	//Start to buy
	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())
	loadCmd := fmt.Sprintf("load-shrp %s %s",userNormal.String(), "4500")
	buyCmd := fmt.Sprintf("buy-shr %s", "232")

	_, stdOut, _ = f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1,f.Port)
	txResponse = ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txResponse.Code)


	account = f.QueryAccount(userNormal)
	accountCoinInfo = account.GetCoins()
	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txResponse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txResponse2.Code)
	tests.WaitForNextNBlocksTM(1,f.Port)
	//Buy
	account = f.QueryAccount(userNormal)
	accountCoinInfo = account.GetCoins()

	_, stdOut3, _ := f.ExecuteGentlemintTxCommand(buyCmd, fmt.Sprintf("--from %s --yes --fees 1shr", userNormal.String()))
	txResponse3 := ParseStdOut(t, stdOut3)
	require.Equal(t, uint32(0), txResponse3.Code)

	tests.WaitForNextNBlocksTM(1,f.Port)

	//validate result
	tests.WaitForNextNBlocksTM(1,f.Port)
	account = f.QueryAccount(userNormal)
	accountCoinInfo = account.GetCoins()

	t.Logf("Amount SHR %v",accountCoinInfo.AmountOf("shr"))
	t.Logf("Amount SHRP %v",accountCoinInfo.AmountOf("shrp"))
	t.Logf("Amount Cent %v",accountCoinInfo.AmountOf("cent"))

	require.Equal(t,"15000230",accountCoinInfo.AmountOf("shr").String()) // buy 232 shr but the normal account must spend 2shr cause load coin buy coin
	require.Equal(t,"4497",accountCoinInfo.AmountOf("shrp").String()) // 232 ~ 2,343434343
	require.Equal(t,"65",accountCoinInfo.AmountOf("cent").String())

	f.Cleanup()
}

func TestGentlemint_BuySHR_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)

	// emptyAccount := f.KeyAddress(keyEmtyUser)
	amount := "10"
	buyAmount := "1"

	enrollLoaderCmd := fmt.Sprintf("enroll-loaders %s", treasurer.String())
	loadCmd := fmt.Sprintf("load-shrp %s %s", treasurer.String(), amount)
	buyCmd := fmt.Sprintf("buy-shr %s", buyAmount)

	// Enroll
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// Buy
	_, stdOut3, _ := f.ExecuteGentlemintTxCommand(buyCmd, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	txRepsonse3 := ParseStdOut(t, stdOut3)
	require.Equal(t, uint32(0), txRepsonse3.Code)

	f.Cleanup()
}

func shrAmountOfAddress(address sdk.AccAddress, f *Fixtures) sdk.Int {
	account := f.QueryAccount(address)

	return account.GetCoins().AmountOf("shr")

}
