package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "2"

	user2 := f.KeyAddressSeed(keyUser2)
	ac:=f.QueryAccount(user2)
	t.Logf("the account information before crate asset %s",ac)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)

	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)
	user2 = f.KeyAddressSeed(keyUser2)
	ac=f.QueryAccount(user2)
	t.Logf("the account information after create asset %s",ac)
	require.Equal(t, ShareLedgerSuccessCode,txResponse.Code)
	assetAfterCrate := f.ExecuteAssetGet(assetID,keyUser2)

	require.Equal(t, assetID,assetAfterCrate.UUID,"the asset ID must match")
	require.Equal(t, assetHash,string(assetAfterCrate.Hash),"the asset Hash must match")
	require.Equal(t, assetStatus,fmt.Sprintf("%v",assetAfterCrate.Status),"the asset status must match")
	require.Equal(t, SHRPFee,fmt.Sprintf("%d",assetAfterCrate.Rate),"the asset rate must match")

	f.Cleanup()
}

func TestCreateDuplicateAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetHash2 := "cc6f58bd1ada876f0a4941ad579908easa726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "2"

	user2 := f.KeyAddressSeed(keyUser2)
	ac:=f.QueryAccount(user2)


	t.Logf("the account information before crate asset %s",ac)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)
	_,stdOut,_ = f.ExecuteAssetCreate(assetHash2,assetID,assetStatus,SHRPFee,keyUser2)

	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)

	require.Equal(t, ShareLedgerErrorCodeInvalidRequest,txResponse.Code)
	f.Cleanup()
}


func TestUpdateExistedAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetHash2 := "cc6f58bd1ada876f0a4941ad579908easa726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	assetStatus2 := "false"
	SHRPFee:= "2"

	user2 := f.KeyAddressSeed(keyUser2)
	ac:=f.QueryAccount(user2)


	t.Logf("the account information before crate asset %s",ac)

	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)
	_,stdOut,_ = f.ExecuteAssetUpdate(assetHash2,assetID,assetStatus2,SHRPFee,keyUser2)
	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)

	assetAfterEdit := f.ExecuteAssetGet(assetID,keyUser2)

	require.Equal(t, assetHash2,string(assetAfterEdit.Hash))
	require.Equal(t, assetID,assetAfterEdit.UUID)
	require.Equal(t, assetStatus2,fmt.Sprintf("%v",assetAfterEdit.Status))
	require.Equal(t, SHRPFee,fmt.Sprintf("%d",assetAfterEdit.Rate))

	require.Equal(t, ShareLedgerSuccessCode,txResponse.Code)
	f.Cleanup()
}


//TestUpdateNotExistedAsset
//
func TestUpdateNotExistedAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "2"

	_,stdOut,_ := f.ExecuteAssetUpdate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)
	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)

	require.Equal(t, ShareLedgerErrorCodeInvalidRequest,txResponse.Code,fmt.Sprintf("the asset:%v doestn't exist in our blockchanin must return invalid reqeust error or not found asset",assetID))
	f.Cleanup()
}


//TestUpdateNotExistedAsset
func TestUpdateAssetUpdaterNotOwnerOfAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "2"


	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)

	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)

	require.Equal(t, ShareLedgerSuccessCode,txResponse.Code)

	_,stdOut,_ = f.ExecuteAssetUpdate(assetHash,assetID,assetStatus,SHRPFee,keyUser1)

	txResponse = ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerErrorCodeUnauthorized,txResponse.Code)
	f.Cleanup()

}

//TestUpdateNotExistedAsset
func TestDeleteAsset(t *testing.T) {
	t.Parallel()
	f := InitFixturesKeySeedModule(t)

	// start gaiad server with minimum fees
	minGasPrice, _ := sdk.NewDecFromStr("0.000006")
	proc := f.GDStart(fmt.Sprintf("--minimum-gas-prices=%s", sdk.NewDecCoinFromDec(feeDenom, minGasPrice)))
	defer proc.Stop(true)


	assetHash := "cc6f58bd1ada876f0a4941ad579908eda726d6da"
	assetID := "1eb07acc-6c2d-4148-889f-61752c49a4b3"
	assetStatus := "true"
	SHRPFee:= "2"


	_,stdOut,_ := f.ExecuteAssetCreate(assetHash,assetID,assetStatus,SHRPFee,keyUser2)

	txResponse := ParseStdOut(t,stdOut)

	tests.WaitForNextHeightTM(f.Port)

	require.Equal(t, ShareLedgerSuccessCode,txResponse.Code)

	_,stdOut,_ = f.ExecuteAssetDelete(assetID,keyUser1)

	txResponse = ParseStdOut(t,stdOut)

	require.Equal(t, ShareLedgerSuccessCode,txResponse.Code)
	f.Cleanup()


	f.Cleanup()
}