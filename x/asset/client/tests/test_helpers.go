package tests

import (
	"github.com/ShareRing/Shareledger/app"
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/ShareRing/Shareledger/x/asset/client/cli"
	"github.com/ShareRing/Shareledger/x/asset/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	"testing"
)

func ExCmdCreateAsset(clientCtx client.Context, assetUUID, assetHash, assetStatus, assetFee string, txCreator string) (testutil.BufferWriter, error) {
	args := []string{assetHash, assetUUID, assetStatus, assetFee}
	args = append(args, network.GetDefaultFlags2SHR(txCreator)...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreate(), args)
}

func ExCmdUpdateAsset(clientCtx client.Context, assetUUID, assetHash, assetStatus, assetFee string, txCreator string) (testutil.BufferWriter, error) {
	args := []string{assetHash, assetUUID, assetStatus, assetFee}
	args = append(args, network.GetDefaultFlags2SHR(txCreator)...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdUpdate(), args)
}

func ExCmdDeleteAsset(clientCtx client.Context, assetUUID string, txCreator string) (testutil.BufferWriter, error) {
	args := []string{assetUUID}
	args = append(args, network.GetDefaultFlags2SHR(txCreator)...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDelete(), args)
}

func ExCmdGetAsset(clientCtx client.Context, assetUUID string) (testutil.BufferWriter, error) {
	args := []string{assetUUID}
	args = append(args, network.GetFlagsQuery()...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdAssetByUUID(), args)
}

func AssetJsonUnmarshal(t *testing.T, data []byte) types.Asset {
	var a types.QueryAssetByUUIDResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(data, &a)
	require.NoError(t, err)
	return *a.Asset

}
