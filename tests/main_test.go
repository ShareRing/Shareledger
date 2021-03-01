package tests

import (
	"fmt"
	"testing"

	app "bitbucket.org/shareringvietnam/shareledger-fix"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/gentlemint"
	"github.com/stretchr/testify/require"

	id "bitbucket.org/shareringvietnam/shareledger-modules/id"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutil "github.com/cosmos/cosmos-sdk/x/genutil"
)

func TestMain_ExportGenesis_Gentlemint_ID(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))

	// Save key addresses for later use
	issuer := f.KeyAddress(keyIdSigner)
	issuer2 := f.KeyAddress(keyUser1)
	authority := f.KeyAddress(keyAuthority)
	treasurer := f.KeyAddress(keyTreasurer)
	operator := f.KeyAddress(keyAccOp)

	_, owners := createRandomAddr(3)
	_, backups := createRandomAddr(3)

	ids := []string{"id-b-1", "id-b-2", "id-b-3"}
	extras := []string{"extras-b-1", "extras-b-2", "extras-b-3"}
	proofs := []string{"proofs-b-1", "proofs-b-2", "proofs-b-3"}

	// Enroll sepcial account
	f.EnrollIdSigner([]sdk.AccAddress{issuer, issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", operator.String()))
	f.EnrollDocIssuer([]sdk.AccAddress{issuer, issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", operator.String()))
	f.EnrollAccountOperator([]sdk.AccAddress{issuer, issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateIdInBatch(ids, backups, owners, extras, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))
	f.CreateNewDocInBatch(ids, proofs, extras, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	f.RevokeIdSigner([]sdk.AccAddress{issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", operator.String()))
	f.RevokeDocIssuer([]sdk.AccAddress{issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", operator.String()))
	f.RevokeAccountOperator([]sdk.AccAddress{issuer2}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	proc.Stop(false)

	cdc := app.MakeCodec()

	gen, _ := f.ExportGenesis("")

	appState, _ := genutil.GenesisStateFromGenDoc(cdc, *gen)

	gentlemintGenState := gentlemint.GetGenesisStateFromAppState(cdc, appState)
	require.Equal(t, authority.String(), gentlemintGenState.Authority)
	require.Equal(t, treasurer.String(), gentlemintGenState.Treasurer)

	require.Equal(t, 2, len(gentlemintGenState.IdSigners))
	require.Equal(t, 2, len(gentlemintGenState.DocumentIssuer))
	require.Equal(t, 3, len(gentlemintGenState.AccountOperators))

	for _, signer := range gentlemintGenState.IdSigners {
		if signer.Address.Equals(issuer2) {
			require.Equal(t, gentlemint.InactiveStatus, signer.Status)
		} else {
			require.Equal(t, gentlemint.ActiveStatus, signer.Status)
		}
	}

	for _, operator := range gentlemintGenState.AccountOperators {
		if operator.Address.Equals(issuer2) {
			require.Equal(t, gentlemint.InactiveStatus, operator.Status)
		} else {
			require.Equal(t, gentlemint.ActiveStatus, operator.Status)
		}
	}

	for _, issuer := range gentlemintGenState.DocumentIssuer {
		if issuer.Address.Equals(issuer2) {
			require.Equal(t, gentlemint.InactiveStatus, issuer.Status)
		} else {
			require.Equal(t, gentlemint.ActiveStatus, issuer.Status)
		}
	}

	idGenState := id.GetGenesisStateFromAppState(cdc, appState)
	fmt.Println(idGenState.IDs)
	require.Equal(t, len(ids), len(idGenState.IDs))

	f.Cleanup()
}
