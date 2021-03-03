/*
 * Based on https://github.com/cosmos/gaia/blob/v2.0.12/cli_test/cli_test.go
 */
package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/sharering/shareledger"
	"github.com/sharering/shareledger/x/gentlemint"
	"github.com/stretchr/testify/require"
)

const ()

var ()

//___________________________________________________________________________________
func (f *Fixtures) EnrollDocIssuer(issuers []sdk.AccAddress, flags ...string) (bool, string, string) {
	listIssuer := make([]string, 0, len(issuers))

	for _, addr := range issuers {
		listIssuer = append(listIssuer, addr.String())
	}

	signers := strings.Join(listIssuer, " ")
	cmd := fmt.Sprintf("%s tx gentlemint enroll-doc-issuer %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) RevokeDocIssuer(issuers []sdk.AccAddress, flags ...string) (bool, string, string) {
	listIssuer := make([]string, 0, len(issuers))

	for _, addr := range issuers {
		listIssuer = append(listIssuer, addr.String())
	}

	signers := strings.Join(listIssuer, " ")
	cmd := fmt.Sprintf("%s tx gentlemint revoke-doc-issuer %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) QueryDocumentIssuer(acc sdk.AccAddress, flags ...string) gentlemint.AccState {
	cmd := fmt.Sprintf("%s query gentlemint document-issuer %s %v", f.GaiacliBinary, acc, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	// fmt.Println("out " + string(out))
	var ids gentlemint.AccState
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &ids)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return ids
}

//___________________________________________________________________________________
func (f *Fixtures) EnrollIdSigner(idSignerAddrs []sdk.AccAddress, flags ...string) (bool, string, string) {
	lstSigner := make([]string, 0, len(idSignerAddrs))

	for _, addr := range idSignerAddrs {
		lstSigner = append(lstSigner, addr.String())
	}

	signers := strings.Join(lstSigner, " ")
	cmd := fmt.Sprintf("%s tx gentlemint enroll-id-signer %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) RevokeIdSigner(idSignerAddrs []sdk.AccAddress, flags ...string) (bool, string, string) {
	lstSigner := make([]string, 0, len(idSignerAddrs))

	for _, addr := range idSignerAddrs {
		lstSigner = append(lstSigner, addr.String())
	}

	signers := strings.Join(lstSigner, " ")
	cmd := fmt.Sprintf("%s tx gentlemint revoke-id-signer %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) QuerySigner(idSignerAddr sdk.AccAddress, flags ...string) gentlemint.AccState {
	cmd := fmt.Sprintf("%s query gentlemint id-signer %s %v", f.GaiacliBinary, idSignerAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	// fmt.Println("out " + string(out))
	var ids gentlemint.AccState
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &ids)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return ids
}

//___________________________________________________________________________________
func (f *Fixtures) EnrollAccountOperator(addrs []sdk.AccAddress, flags ...string) (bool, string, string) {
	lstSigner := make([]string, 0, len(addrs))

	for _, addr := range addrs {
		lstSigner = append(lstSigner, addr.String())
	}

	signers := strings.Join(lstSigner, " ")
	cmd := fmt.Sprintf("%s tx gentlemint enroll-account-operator %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) RevokeAccountOperator(accs []sdk.AccAddress, flags ...string) (bool, string, string) {
	lstSigner := make([]string, 0, len(accs))

	for _, addr := range accs {
		lstSigner = append(lstSigner, addr.String())
	}

	signers := strings.Join(lstSigner, " ")
	cmd := fmt.Sprintf("%s tx gentlemint revoke-account-operator %s %v", f.GaiacliBinary, signers, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func (f *Fixtures) QueryAccountOperator(addr sdk.AccAddress, flags ...string) gentlemint.AccState {
	cmd := fmt.Sprintf("%s query gentlemint account-operator %s %v", f.GaiacliBinary, addr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	var acc gentlemint.AccState
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &acc)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return acc
}

func (f *Fixtures) QueryAllAccountOperator(flags ...string) []gentlemint.AccState {
	cmd := fmt.Sprintf("%s query gentlemint account-operators %v", f.GaiacliBinary, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")

	var accs []gentlemint.AccState
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &accs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return accs
}

func (f *Fixtures) ExecuteGentlemintTxCommand(command string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx gentlemint %s %v", f.GaiacliBinary, command, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), DefaultKeyPass)
}

func ParseStdOut(t *testing.T, stdOut string) sdk.TxResponse {
	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)
	return txRepsonse
}
