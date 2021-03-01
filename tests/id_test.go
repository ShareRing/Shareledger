package tests

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gatecheck"
	"github.com/stretchr/testify/require"
)

func TestIdCreate_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	id := "id-001"
	extraData := "extradata-001"

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Equal(t, backup.String(), rsId.BackupAddr.String())
	require.Equal(t, owner.String(), rsId.OwnerAddr.String())
	require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
	require.Equal(t, extraData, rsId.ExtraData)

	rsId2 := f.QueryIdByOwner(owner)

	require.NotNil(t, rsId2)
	require.Equal(t, id, rsId2.Id)
	require.Equal(t, backup, rsId2.BackupAddr)
	require.Equal(t, owner.String(), rsId2.OwnerAddr.String())
	require.Equal(t, issuer, rsId2.IssuerAddr)
	require.Equal(t, extraData, rsId2.ExtraData)

	f.Cleanup()
}

func TestIdCreateInBatch(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	issuer := f.KeyAddress(keyIdSigner)
	accountOperator := f.KeyAddress(keyAccOp)

	_, owners := createRandomAddr(3)
	_, backups := createRandomAddr(3)

	ids := []string{"id-b-1", "id-b-2", "id-b-3"}
	extras := []string{"extras-b-1", "extras-b-2", "extras-b-3"}

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateIdInBatch(ids, backups, owners, extras, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	for i, id := range ids {
		rsId := f.QueryIdById(id)

		require.NotNil(t, rsId)
		require.Equal(t, ids[i], rsId.Id)
		require.Equal(t, backups[i].String(), rsId.BackupAddr.String())
		require.Equal(t, owners[i].String(), rsId.OwnerAddr.String())
		require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
		require.Equal(t, extras[i], rsId.ExtraData)
	}

	f.Cleanup()
}

func TestIdCreateInBatch_CallerIsNotIdSigner(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	treasurer := f.KeyAddress(keyTreasurer)

	_, owners := createRandomAddr(3)
	_, backups := createRandomAddr(3)

	ids := []string{"id-b-1", "id-b-2", "id-b-3"}
	extras := []string{"extras-b-1", "extras-b-2", "extras-b-3"}

	_, stdOut, _ := f.CreateIdInBatch(ids, backups, owners, extras, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))

	require.Contains(t, stdOut, gatecheck.NotIdSignerErr)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	for _, id := range ids {
		rsId := f.QueryIdById(id)

		require.NotNil(t, rsId)
		require.True(t, rsId.IsEmpty())
	}

	f.Cleanup()
}

func TestIdCreate_CallerIsNotIssuer(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	issuer := f.KeyAddress(keyValidator)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	id := "id-not-issuer"
	extraData := "extradata-002"

	_, stdOut, _ := f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.Contains(t, stdOut, gatecheck.NotIdSignerErr)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Empty(t, rsId.BackupAddr)
	require.Empty(t, rsId.OwnerAddr.String())
	require.Empty(t, rsId.IssuerAddr)
	require.Empty(t, rsId.ExtraData)

	rsId2 := f.QueryIdByOwner(owner)

	require.NotNil(t, rsId2)
	require.Empty(t, rsId2.Id)
	require.Empty(t, rsId2.BackupAddr)
	require.Empty(t, rsId2.OwnerAddr.String())
	require.Empty(t, rsId2.IssuerAddr)
	require.Empty(t, rsId2.ExtraData)

	f.Cleanup()
}

func TestIdUpdate_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	id := "id-001"
	extraData := "extradata-001"

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Equal(t, backup.String(), rsId.BackupAddr.String())
	require.Equal(t, owner.String(), rsId.OwnerAddr.String())
	require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
	require.Equal(t, extraData, rsId.ExtraData)

	// Update id
	newExtradata := "newExtradata"
	f.UpdateId(id, newExtradata, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	updatedId := f.QueryIdById(id)

	require.NotNil(t, updatedId)
	require.Equal(t, id, updatedId.Id)
	require.Equal(t, backup.String(), updatedId.BackupAddr.String())
	require.Equal(t, owner.String(), updatedId.OwnerAddr.String())
	require.Equal(t, issuer.String(), updatedId.IssuerAddr.String())
	require.Equal(t, newExtradata, updatedId.ExtraData)

	f.Cleanup()
}

func TestIdUpdate_CallerIsNotIdSigner(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	user3 := f.KeyAddress(keyUser3)
	id := "id-001"
	extraData := "extradata-001"

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Equal(t, backup.String(), rsId.BackupAddr.String())
	require.Equal(t, owner.String(), rsId.OwnerAddr.String())
	require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
	require.Equal(t, extraData, rsId.ExtraData)

	// Update id
	newExtradata := "newExtradata"
	_, stdOut, _ := f.UpdateId(id, newExtradata, fmt.Sprintf("--from %s --yes --fees 1shr", user3.String()))

	require.Contains(t, stdOut, gatecheck.NotIdSignerErr)

	updatedId := f.QueryIdById(id)

	require.NotNil(t, updatedId)
	require.Equal(t, id, updatedId.Id)
	require.Equal(t, backup.String(), updatedId.BackupAddr.String())
	require.Equal(t, owner.String(), updatedId.OwnerAddr.String())
	require.Equal(t, issuer.String(), updatedId.IssuerAddr.String())
	require.Equal(t, extraData, updatedId.ExtraData)

	f.Cleanup()
}

func TestIdReplaceOwner_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	id := "id-001"
	extraData := "extradata-001"

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Equal(t, backup.String(), rsId.BackupAddr.String())
	require.Equal(t, owner.String(), rsId.OwnerAddr.String())
	require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
	require.Equal(t, extraData, rsId.ExtraData)

	// Update id
	f.ReplaceIdOwner(id, backup, fmt.Sprintf("--from %s --yes --fees 1shr", backup.String()))

	updatedId := f.QueryIdById(id)

	require.NotNil(t, updatedId)
	require.Equal(t, id, updatedId.Id)
	require.Equal(t, backup.String(), updatedId.BackupAddr.String())
	require.Equal(t, backup.String(), updatedId.OwnerAddr.String())
	require.Equal(t, issuer.String(), updatedId.IssuerAddr.String())
	require.Equal(t, extraData, updatedId.ExtraData)

	f.Cleanup()
}

func TestIdReplaceOwner_CallerIsNotBackup(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	owner := f.KeyAddress(keyUser1)
	backup := f.KeyAddress(keyUser2)
	id := "id-001"
	extraData := "extradata-001"

	// Enroll id signer
	f.EnrollIdSigner([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateId(id, backup, owner, extraData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	rsId := f.QueryIdById(id)

	require.NotNil(t, rsId)
	require.Equal(t, id, rsId.Id)
	require.Equal(t, backup.String(), rsId.BackupAddr.String())
	require.Equal(t, owner.String(), rsId.OwnerAddr.String())
	require.Equal(t, issuer.String(), rsId.IssuerAddr.String())
	require.Equal(t, extraData, rsId.ExtraData)

	// Update id
	_, stdOut, _ := f.ReplaceIdOwner(id, backup, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.Contains(t, stdOut, gatecheck.NotBackupAccountErr)

	updatedId := f.QueryIdById(id)

	require.NotNil(t, updatedId)
	require.Equal(t, id, updatedId.Id)
	require.Equal(t, backup.String(), updatedId.BackupAddr.String())
	require.Equal(t, owner.String(), updatedId.OwnerAddr.String())
	require.Equal(t, issuer.String(), updatedId.IssuerAddr.String())
	require.Equal(t, extraData, updatedId.ExtraData)

	f.Cleanup()
}
