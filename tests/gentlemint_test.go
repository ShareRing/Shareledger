package tests

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/sharering/shareledger"
	"github.com/sharering/shareledger/x/gentlemint"
	"github.com/sharering/shareledger/x/gentlemint/types"
	gentTypes "github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/stretchr/testify/require"
)

func TestGentlemint_EnrollAccountOperator_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	_, operators := createRandomAddr(3)

	// Enroll
	_, stdOut, _ := f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	info := f.QueryAllAccountOperator()

	// Append the pre-operators
	operators = append(operators, f.KeyAddress(keyAccOp))

	require.Equal(t, len(operators), len(info))

	require.Contains(t, operators, info[0].Address)
	require.Contains(t, operators, info[1].Address)
	require.Contains(t, operators, info[2].Address)
	require.Contains(t, operators, info[3].Address)

	require.Equal(t, gentTypes.Active, info[0].Status)
	require.Equal(t, gentTypes.Active, info[1].Status)
	require.Equal(t, gentTypes.Active, info[2].Status)
	require.Equal(t, gentTypes.Active, info[3].Status)

	acc0 := f.QueryAccountOperator(operators[0])
	require.Equal(t, gentTypes.Active, acc0.Status)
	require.Equal(t, operators[0], acc0.Address)

	// Check events
	for _, event := range txRepsonse.Logs[0].Events {
		if event.Type == gentTypes.EventTypeEnrollAccOp {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: operators[0].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: operators[1].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: operators[2].String()})
		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeEnrollAccOp})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: authority.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_EnrollAccountOperator_NotAuthority(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyUser1)
	_, operators := createRandomAddr(1)

	// Enroll
	_, stdOut, _ := f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	require.Contains(t, stdOut, types.ErrSenderIsNotAuthority)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	acc0 := f.QueryAccountOperator(operators[0])
	require.Empty(t, acc0.Status)
	require.Empty(t, acc0.Address)

	f.Cleanup()
}

func TestGentlemint_RevokeAccountOperator_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	_, operators := createRandomAddr(3)

	// Enroll
	f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeAccountOperator([]sdk.AccAddress{operators[0], operators[1]}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	acc0 := f.QueryAccountOperator(operators[0])
	require.Equal(t, gentTypes.Inactive, acc0.Status)
	require.Equal(t, operators[0], acc0.Address)

	// acc1 := f.QueryAccountOperator(operators[1])
	// require.Equal(t, gentTypes.Inactive, acc1.Status)
	// require.Equal(t, operators[1], acc1.Address)

	// Check events
	for _, event := range txRepsonse.Logs[0].Events {
		if event.Type == gentTypes.EventTypeRevokeAccOp {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: operators[0].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: operators[1].String()})
		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeRevokeAccOp})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: authority.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_RevokeAccountOperator_DoesNotExist(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	_, operators := createRandomAddr(3)
	_, doesNotExistOperators := createRandomAddr(1)

	// Enroll
	f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeAccountOperator([]sdk.AccAddress{operators[0], doesNotExistOperators[0]}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	require.Contains(t, stdOut, types.ErrDoesNotExist.Error())

	acc0 := f.QueryAccountOperator(operators[0])
	require.Equal(t, gentTypes.Active, acc0.Status)
	require.Equal(t, operators[0], acc0.Address)

	f.Cleanup()
}

func TestGentlemint_RevokeAccountOperator_Inactive(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	_, operators := createRandomAddr(3)

	// Enroll
	f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.RevokeAccountOperator([]sdk.AccAddress{operators[0]}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeAccountOperator([]sdk.AccAddress{operators[1], operators[0]}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	require.Contains(t, stdOut, types.ErrDoesNotExist.Error())

	acc0 := f.QueryAccountOperator(operators[0])
	require.Equal(t, gentTypes.Inactive, acc0.Status)
	require.Equal(t, operators[0], acc0.Address)

	acc1 := f.QueryAccountOperator(operators[1])
	require.Equal(t, gentTypes.Active, acc1.Status)
	require.Equal(t, operators[1], acc1.Address)

	f.Cleanup()
}

func TestGentlemint_RevokeAccountOperator_NotAuthority(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	authority := f.KeyAddress(keyAuthority)
	user1 := f.KeyAddress(keyUser1)
	_, operators := createRandomAddr(1)

	// Enroll
	f.EnrollAccountOperator(operators, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeAccountOperator([]sdk.AccAddress{operators[0]}, fmt.Sprintf("--from %s --yes --fees 1shr", user1.String()))

	require.Contains(t, stdOut, types.ErrSenderIsNotAuthority)

	acc0 := f.QueryAccountOperator(operators[0])
	require.Equal(t, gentTypes.Active, acc0.Status)
	require.Equal(t, operators[0], acc0.Address)

	f.Cleanup()
}

//-______________________________________________

func TestGentlemint_EnrollIssuer_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accOperator := f.KeyAddress(keyAccOp)
	_, issuers := createRandomAddr(3)

	// Enroll issuer
	_, stdOut, _ := f.EnrollDocIssuer(issuers, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))
	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	acc := f.QueryDocumentIssuer(issuers[0])

	require.Equal(t, acc.Address, issuers[0])
	require.Equal(t, acc.Status, gentlemint.ActiveStatus)

	// Check events
	for _, event := range txRepsonse.Logs[0].Events {
		if event.Type == gentTypes.EventTypeEnrollDocIssuer {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[0].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[1].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[2].String()})
		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeEnrollDocIssuer})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: accOperator.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_RevokeIssuer_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accOperator := f.KeyAddress(keyAccOp)
	_, issuers := createRandomAddr(3)

	// Enroll issuer
	f.EnrollDocIssuer(issuers, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeDocIssuer([]sdk.AccAddress{issuers[0], issuers[1]}, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))
	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	acc := f.QueryDocumentIssuer(issuers[0])

	require.Equal(t, acc.Address, issuers[0])
	require.Equal(t, acc.Status, gentlemint.InactiveStatus)

	// Check events
	for _, event := range txRepsonse.Logs[0].Events {
		if event.Type == gentTypes.EventTypeRevokeDocIssuer {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[0].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[1].String()})
			// require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: issuers[2].String()})
		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeRevokeDocIssuer})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: accOperator.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_RevokeIssuer_NotAccountOperator(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accOperator := f.KeyAddress(keyAccOp)
	issuer := f.KeyAddress(keyIdSigner)
	authority := f.KeyAddress(keyAuthority)

	// Enroll issuer
	f.EnrollDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", accOperator.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	_, stdOut, _ := f.RevokeDocIssuer([]sdk.AccAddress{issuer}, fmt.Sprintf("--from %s --yes --fees 1shr", authority.String()))

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	acc := f.QueryDocumentIssuer(issuer)

	require.Equal(t, acc.Address, issuer)
	require.Equal(t, acc.Status, gentlemint.ActiveStatus)
	require.Contains(t, stdOut, gentTypes.ErrSenderIsNotAccountOperator)

	f.Cleanup()
}

//_________________________________________________________________________________

func TestGentlemint_EnrollIdSigner_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	// user1 := f.KeyAddress(keyUser1)
	// user2 := f.KeyAddress(keyUser2)
	// user3 := f.KeyAddress(keyUser3)

	_, signers := createRandomAddr(3)
	// Enroll id signer
	success, stdOut, _ := f.EnrollIdSigner(signers, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))

	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	require.True(t, success)
	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ids1 := f.QuerySigner(signers[0])

	require.NotNil(t, ids1)
	require.Equal(t, ids1.Address.String(), signers[0].String())
	require.True(t, ids1.IsActive())

	ids2 := f.QuerySigner(signers[1])

	require.NotNil(t, ids2)
	require.Equal(t, ids2.Address.String(), signers[1].String())
	require.True(t, ids2.IsActive())

	ids3 := f.QuerySigner(signers[2])

	require.NotNil(t, ids3)
	require.Equal(t, ids3.Address.String(), signers[2].String())
	require.True(t, ids3.IsActive())

	// Check events
	expectEvents := []string{gentTypes.EventTypeEnrollIdSigner, sdk.EventTypeMessage, "transfer"}
	for _, event := range txRepsonse.Logs[0].Events {
		require.Contains(t, expectEvents, event.Type)
		if event.Type == gentTypes.EventTypeEnrollIdSigner {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: signers[0].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: signers[1].String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: signers[2].String()})
		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeEnrollIdSigner})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: accountOperator.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_EnrollIdSigner_NotAccountOperator(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	treasurer := f.KeyAddress(keyTreasurer)
	user1 := f.KeyAddress(keyUser1)
	user2 := f.KeyAddress(keyUser2)

	// Enroll id signer
	success, stdOut, _ := f.EnrollIdSigner([]sdk.AccAddress{user1, user2}, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))
	require.Contains(t, stdOut, gentTypes.ErrSenderIsNotAccountOperator)
	require.True(t, success)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ids1 := f.QuerySigner(user1)
	require.NotNil(t, ids1)
	require.True(t, ids1.IsEmpty())
	require.False(t, ids1.IsActive())

	ids2 := f.QuerySigner(user2)
	require.NotNil(t, ids2)
	require.True(t, ids2.IsEmpty())
	require.False(t, ids2.IsActive())

	f.Cleanup()
}

func TestGentlemint_RevokeIdSigner_Success(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	user1 := f.KeyAddress(keyUser1)
	user2 := f.KeyAddress(keyUser2)
	user3 := f.KeyAddress(keyUser3)

	// Enroll id signer
	success, _, _ := f.EnrollIdSigner([]sdk.AccAddress{user1, user2, user3}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))
	require.True(t, success)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Revoke
	revokeSuccess, stdOut, _ := f.RevokeIdSigner([]sdk.AccAddress{user2, user3}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))
	require.True(t, revokeSuccess)

	txRepsonse := sdk.TxResponse{}
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdOut), &txRepsonse)
	require.Nil(t, err)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ids1 := f.QuerySigner(user1)

	require.NotNil(t, ids1)
	require.Equal(t, ids1.Address.String(), user1.String())
	require.True(t, ids1.IsActive())

	ids2 := f.QuerySigner(user2)

	require.NotNil(t, ids2)
	require.False(t, ids2.IsActive())

	ids3 := f.QuerySigner(user3)

	require.NotNil(t, ids3)
	require.False(t, ids3.IsActive())

	// Check events
	expectEvents := []string{gentTypes.EventTypeRevokeIdSigner, sdk.EventTypeMessage, "transfer"}
	for _, event := range txRepsonse.Logs[0].Events {
		require.Contains(t, expectEvents, event.Type)
		if event.Type == gentTypes.EventTypeRevokeIdSigner {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: user2.String()})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: gentTypes.EventAttrAddress, Value: user3.String()})

		} else if event.Type == sdk.EventTypeMessage {
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyModule, Value: gentTypes.ModuleName})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeyAction, Value: gentTypes.EventTypeRevokeIdSigner})
			require.Contains(t, event.Attributes, sdk.Attribute{Key: sdk.AttributeKeySender, Value: accountOperator.String()})
		}
	}

	f.Cleanup()
}

func TestGentlemint_RevokeIdSigner_NotAccountOperator(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(false)

	// Save key addresses for later use
	accountOperator := f.KeyAddress(keyAccOp)
	treasurer := f.KeyAddress(keyTreasurer)
	user1 := f.KeyAddress(keyUser1)
	user2 := f.KeyAddress(keyUser2)
	user3 := f.KeyAddress(keyUser3)

	// Enroll id signer
	success, _, _ := f.EnrollIdSigner([]sdk.AccAddress{user1, user2, user3}, fmt.Sprintf("--from %s --yes --fees 1shr", accountOperator.String()))
	require.True(t, success)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Revoke
	_, stdOut, _ := f.RevokeIdSigner([]sdk.AccAddress{user1}, fmt.Sprintf("--from %s --yes --fees 1shr", treasurer.String()))
	require.Contains(t, stdOut, gentTypes.ErrSenderIsNotAccountOperator)

	// wait for a block confirmation
	tests.WaitForNextNBlocksTM(1, f.Port)

	ids1 := f.QuerySigner(user1)

	require.NotNil(t, ids1)
	require.Equal(t, ids1.Address.String(), user1.String())
	require.True(t, ids1.IsActive())

	ids2 := f.QuerySigner(user2)

	require.NotNil(t, ids2)
	require.Equal(t, ids2.Address.String(), user2.String())
	require.True(t, ids2.IsActive())

	f.Cleanup()
}
