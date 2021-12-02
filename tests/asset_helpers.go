package tests

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/tests"
	app "github.com/sharering/shareledger"
	"github.com/sharering/shareledger/x/asset"
	"github.com/stretchr/testify/require"
)

func (f *Fixtures) ExecuteAssetCreate(assetHash,assetID ,status ,SHRPFeeFloat string,userKey string) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", userKey)}
	cmd := fmt.Sprintf("%s tx asset create %v %v %v %v %v", f.GaiacliBinary,assetHash ,assetID,status,SHRPFeeFloat, f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}


func (f *Fixtures) ExecuteAssetUpdate(assetHash,assetID ,status ,SHRPFeeFloat string,userKey string) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", userKey)}
	cmd := fmt.Sprintf("%s tx asset update %v %v %v %v %v", f.GaiacliBinary,assetHash ,assetID,status,SHRPFeeFloat, f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}

func (f *Fixtures) ExecuteAssetDelete(assetID ,userKey string) (bool, string, string) {
	flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 5shr", userKey)}
	cmd := fmt.Sprintf("%s tx asset delete %v %v", f.GaiacliBinary ,assetID, f.Flags())
	return executeWriteRetStdStreams(f.T,addFlags(cmd,flag), DefaultKeyPass)
}


func (f *Fixtures) ExecuteAssetGet(assetID string,userKey string)asset.Asset{
	//flag := []string{fmt.Sprintf("--key-seed ./%s_key_seed.json --yes --fees 1shr", userKey)}
	cmd := fmt.Sprintf("%s query asset get %v %v", f.GaiacliBinary ,assetID, f.Flags())

	out, _ := tests.ExecuteT(f.T, cmd, "")
	var assetHolding asset.Asset
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &assetHolding)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)

	return assetHolding
}