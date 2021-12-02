package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/electoral"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	VoterPrefix         = electoral.VoterPrefix
	VoterStatusEnrolled = electoral.StatusVoterEnrolled

)

func TestErrollVoter_And_GetVoter(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedForElectoralModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer func(proc *tests.Process, kill bool) {
		_ = proc.Stop(kill)
	}(proc, false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	user1 := f.KeyAddress(keyUser1)
	t.Logf("starting to enroll voter:adress=%s the authority user=%s", user1, authority)
	_, stdOut, _ := f.ExecuteElectoralEnrollVoter(keyAuthority, user1)
	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, ShareLedgerSuccessCode, txResponse.Code)
	tests.WaitForNextHeightTM(f.Port)

	voterInformation := f.ExecuteElectoralGetVoter(user1)

	require.NotNil(t, voterInformation, "a nil voter information mean can't create voter")

	require.Equal(t, user1.String() , voterInformation.Address.String(), "the address of voter must equal")
	require.Equal(t, VoterStatusEnrolled, voterInformation.Status, "the address of voter must equal")
	f.Cleanup()
}

//Can't run this test case must be edited later, cause the target function don't support --form flags of cosmos SDK
func TestRevokeVoterAndGetVoterToRecheck(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedForElectoralModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)

	// Save key addresses for later use
	authority := f.KeyAddressSeed(keyAuthority)
	user1 := f.KeyAddress(keyUser1)
	t.Logf("starting to enroll voter:adress=%s the authority user=%s", user1, authority)
	_, stdOut, _ := f.ExecuteElectoralEnrollVoter(keyAuthority, user1)
	txResponse := ParseStdOut(t, stdOut)
	require.Equal(t, ShareLedgerSuccessCode, txResponse.Code)
	tests.WaitForNextHeightTM(f.Port)

	voterInformation := f.ExecuteElectoralGetVoter(user1)

	require.NotNil(t, voterInformation, "a nil voter information mean can't create voter")

	require.Equal(t, user1.String() , voterInformation.Address.String(), "the address of voter must equal")
	require.Equal(t, VoterStatusEnrolled, voterInformation.Status, "the address of voter must equal")

	//revoke the voter that we added
	_, stdOut, _ = f.ExecuteElectoralRevokeVoter(keyAuthority, user1)
	t.Logf("the response after revoke %v",stdOut)
	txResponse = ParseStdOut(t, stdOut)
	require.Equal(t, ShareLedgerSuccessCode, txResponse.Code)
	tests.WaitForNextHeightTM(f.Port)

	voterInformation = f.ExecuteElectoralGetVoter(user1)

	require.Equal(t, voterInformation.Status,"revoked", "status must be revoked")

	f.Cleanup()
}
