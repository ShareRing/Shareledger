package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

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
	f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.ExecuteGentlemintTxCommand(burnCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))
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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Revoke
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(revokeLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))
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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// Burn
	_, stdOut3, _ := f.ExecuteGentlemintTxCommand(burnCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

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
	setExchageRateCmd := fmt.Sprintf("set-exchange %s", rate)

	_, stdOut, _ := f.ExecuteGentlemintTxCommand(setExchageRateCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	cmd := fmt.Sprintf("%s query gentlemint get-exchange %v", f.GaiacliBinary, f.Flags())
	flags := []string{}
	exchangeRate, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	require.Equal(t, rate, strings.Replace(exchangeRate, "\"", "", -1))

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
	_, stdOut, _ := f.ExecuteGentlemintTxCommand(enrollLoaderCmd, fmt.Sprintf("--from %s --yes", authority.String()))

	txRepsonse := ParseStdOut(t, stdOut)
	require.Equal(t, uint32(0), txRepsonse.Code)

	// Load
	_, stdOut2, _ := f.ExecuteGentlemintTxCommand(loadCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

	txRepsonse2 := ParseStdOut(t, stdOut2)
	require.Equal(t, uint32(0), txRepsonse2.Code)

	// Buy
	_, stdOut3, _ := f.ExecuteGentlemintTxCommand(buyCmd, fmt.Sprintf("--from %s --yes", treasurer.String()))

	txRepsonse3 := ParseStdOut(t, stdOut3)
	require.Equal(t, uint32(0), txRepsonse3.Code)

	f.Cleanup()
}
