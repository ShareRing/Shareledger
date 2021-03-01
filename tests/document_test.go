package tests

import (
	"fmt"
	"testing"

	"github.com/ShareRing/modules/document"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDocCreate_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	f.Cleanup()
}

func TestDocCreate_Duplicate(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	ok, stdOut, _ := f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.True(t, ok)
	require.Contains(t, stdOut, document.ErrDocExisted.Error())

	f.Cleanup()
}

func TestDocCreateInBatch_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderIds := []string{"id-001", "id-002", "id-003"}
	proofs := []string{"a89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "b89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"}
	datas := []string{"extradata-001", "extradata-002", "extradata-003"}

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ok, stdOut, stdErr := f.CreateNewDocInBatch(holderIds, proofs, datas, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))
	fmt.Println(stdOut)
	fmt.Println(stdErr)

	require.True(t, ok)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	for i := 0; i < len(holderIds); i++ {
		doc := f.QueryDocByProof(proofs[i])
		require.NotNil(t, doc)
		require.Equal(t, proofs[i], doc.Proof)
		require.Equal(t, issuer.String(), doc.Issuer.String())
		require.Equal(t, datas[i], doc.Data)
		require.Equal(t, holderIds[i], doc.Holder)

	}

	f.Cleanup()
}

func TestDocCreateInBatch_Duplicate(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderIds := []string{"id-001", "id-002", "id-003"}
	proofs := []string{"a89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "b89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"}
	datas := []string{"extradata-001", "extradata-002", "extradata-003"}

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderIds[2], proofs[2], "test2222", fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ok, stdOut, _ := f.CreateNewDocInBatch(holderIds, proofs, datas, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	require.True(t, ok)
	require.Contains(t, stdOut, document.ErrDocExisted.Error())

	// for i := 0; i < len(holderIds); i++ {
	// 	doc := f.QueryDocByProof(proofs[i])
	// 	fmt.Println(doc.String())
	// 	// require.NotNil(t, doc)
	// 	// require.Equal(t, proofs[i], doc.Proof)
	// 	// require.Equal(t, issuer.String(), doc.Issuer.String())
	// 	// require.Equal(t, datas[i], doc.Data)
	// 	// require.Equal(t, holderIds[i], doc.Holder)

	// }

	f.Cleanup()
}

func TestDocUpdate_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	newData := "new-data"
	ok, _, _ := f.UdpateDoc(holderId, proof, newData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.True(t, ok)

	docUpdated := f.QueryDocByProof(proof)
	require.NotNil(t, docUpdated)
	require.Equal(t, proof, docUpdated.Proof)
	require.Equal(t, issuer.String(), docUpdated.Issuer.String())
	require.Equal(t, newData, docUpdated.Data)
	require.Equal(t, doc.Version+1, docUpdated.Version)

	f.Cleanup()
}

func TestDocUpdate_NotExist_Holder(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	newData := "new-data"
	ok, stdOut, _ := f.UdpateDoc(holderId+"1", proof, newData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))
	require.Contains(t, stdOut, document.ErrDocNotExisted.Error())

	require.True(t, ok)

	docUpdated := f.QueryDocByProof(proof)
	require.NotNil(t, docUpdated)
	require.Equal(t, proof, docUpdated.Proof)
	require.Equal(t, issuer.String(), docUpdated.Issuer.String())
	require.Equal(t, data, docUpdated.Data)
	require.Equal(t, doc.Version, docUpdated.Version)

	f.Cleanup()
}

func TestDocUpdate_NotExist_Proof(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	newData := "new-data"
	ok, stdOut, _ := f.UdpateDoc(holderId, proof+"1", newData, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))
	require.Contains(t, stdOut, document.ErrDocNotExisted.Error())

	require.True(t, ok)

	docUpdated := f.QueryDocByProof(proof)
	require.NotNil(t, docUpdated)
	require.Equal(t, proof, docUpdated.Proof)
	require.Equal(t, issuer.String(), docUpdated.Issuer.String())
	require.Equal(t, data, docUpdated.Data)
	require.Equal(t, doc.Version, docUpdated.Version)

	f.Cleanup()
}

func TestDocUpdate_NotExist_Issuer(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderId := "id-001"
	proof := "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6"
	data := "extradata-001"
	// user1 := f.KeyAddress(keyUser1)
	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	f.CreateNewDoc(holderId, proof, data, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	doc := f.QueryDocByProof(proof)
	require.NotNil(t, doc)
	require.Equal(t, proof, doc.Proof)
	require.Equal(t, issuer.String(), doc.Issuer.String())
	require.Equal(t, data, doc.Data)

	newData := "new-data"
	ok, stdOut, _ := f.UdpateDoc(holderId, proof, newData, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))
	require.Contains(t, stdOut, "is not document issuer")

	require.True(t, ok)

	docUpdated := f.QueryDocByProof(proof)
	require.NotNil(t, docUpdated)
	require.Equal(t, proof, docUpdated.Proof)
	require.Equal(t, issuer.String(), docUpdated.Issuer.String())
	require.Equal(t, data, docUpdated.Data)
	require.Equal(t, doc.Version, docUpdated.Version)

	f.Cleanup()
}

func TestDocRevoke_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderIds := []string{"id-001", "id-002", "id-003"}
	proofs := []string{"a89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "b89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"}
	datas := []string{"extradata-001", "extradata-002", "extradata-003"}

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	ok, _, _ := f.CreateNewDocInBatch(holderIds, proofs, datas, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.True(t, ok)
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	revokeRs, _, _ := f.RevokeDoc(holderIds[2], proofs[2], fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	tests.WaitForNextNBlocksTM(1, f.Port)

	require.True(t, revokeRs)
	for i := 0; i < len(holderIds)-1; i++ {
		doc := f.QueryDocByProof(proofs[i])
		require.NotNil(t, doc)
		require.Equal(t, proofs[i], doc.Proof)
		require.Equal(t, issuer.String(), doc.Issuer.String())
		require.Equal(t, datas[i], doc.Data)
		require.Equal(t, holderIds[i], doc.Holder)
		require.Equal(t, uint16(0), doc.Version)
	}

	docRevoked := f.QueryDocByProof(proofs[2])
	require.NotNil(t, docRevoked)
	require.Equal(t, proofs[2], docRevoked.Proof)
	require.Equal(t, issuer.String(), docRevoked.Issuer.String())
	require.Equal(t, datas[2], docRevoked.Data)
	require.Equal(t, holderIds[2], docRevoked.Holder)
	require.Equal(t, uint16(document.DocRevokeFlag), docRevoked.Version)

	f.Cleanup()
}

func TestDocRevoke_NotExist(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderIds := []string{"id-001", "id-002", "id-003"}
	proofs := []string{"a89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "b89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"}
	datas := []string{"extradata-001", "extradata-002", "extradata-003"}

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	ok, _, _ := f.CreateNewDocInBatch(holderIds, proofs, datas, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	require.True(t, ok)
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	revokeRs, stdOut, _ := f.RevokeDoc(holderIds[2], "not-exist", fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))

	tests.WaitForNextNBlocksTM(1, f.Port)

	require.True(t, revokeRs)
	require.Contains(t, stdOut, document.ErrDocNotExisted.Error())

	for i := 0; i < len(holderIds)-1; i++ {
		doc := f.QueryDocByProof(proofs[i])
		require.NotNil(t, doc)
		require.Equal(t, proofs[i], doc.Proof)
		require.Equal(t, issuer.String(), doc.Issuer.String())
		require.Equal(t, datas[i], doc.Data)
		require.Equal(t, holderIds[i], doc.Holder)
		require.Equal(t, uint16(0), doc.Version)
	}

	f.Cleanup()
}

func TestDocQuery_ByHolder_OK(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	holderIds := []string{"id-001", "id-001", "id-003"}
	proofs := []string{"a89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6", "b89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc7", "c89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc8"}
	datas := []string{"extradata-001", "extradata-002", "extradata-003"}

	// Enroll doc issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))

	// wait for a block confirmation
	// tests.WaitForNextNBlocksTM(1, f.Port)

	ok, stdOut, stdErr := f.CreateNewDocInBatch(holderIds, proofs, datas, fmt.Sprintf("--from %s --yes --fees 1shr", issuer.String()))
	fmt.Println(stdOut)
	fmt.Println(stdErr)

	require.True(t, ok)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	docOf1 := f.QueryDocByHolder(holderIds[0])

	fmt.Println(docOf1)
	require.Equal(t, 2, len(docOf1))

	f.Cleanup()
}
