package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/asset/client/cli"
	"github.com/sharering/shareledger/x/asset/types"
)

func ExCmdCreateAsset(clientCtx client.Context, assetUUID, assetHash, assetStatus, assetFee string, userFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetHash, assetUUID, assetStatus, assetFee}
	args = append(args, network.SkipConfirmation, network.BlockBroadcast)
	args = append(args, userFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreate(), args)
}

func ExCmdUpdateAsset(clientCtx client.Context, assetUUID, assetHash, assetStatus, assetFee string, userFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetHash, assetUUID, assetStatus, assetFee}
	args = append(args, network.SkipConfirmation, network.BlockBroadcast)
	args = append(args, userFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdUpdate(), args)
}

func ExCmdDeleteAsset(clientCtx client.Context, assetUUID string, userFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetUUID}
	args = append(args, network.SkipConfirmation, network.BlockBroadcast)
	args = append(args, userFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDelete(), args)
}

func ExCmdGetAsset(clientCtx client.Context, assetUUID string, userFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetUUID}
	args = append(args, userFlags...)
	args = append(args, network.JSONFlag)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdAssetByUUID(), args)
}

func AssetJsonUnmarshal(t *testing.T, data []byte) types.Asset {
	var a types.QueryAssetByUUIDResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(data, &a)
	require.NoError(t, err)
	return *a.Asset

}
