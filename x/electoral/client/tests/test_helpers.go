package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/x/electoral/client/cli"
	"github.com/sharering/shareledger/x/electoral/types"
)

func ExCmdEnrollDocIssuer(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := append(address, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdEnrollDocIssuers(), args)
}
func ExCmdRevokeDocIssuer(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := append(address, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeDocIssuers(), args)

}

func ExCmdGetDocIssuers(clientCtx client.Context, t *testing.T, additionalFlags ...string) (testutil.BufferWriter, error) {
	var args []string
	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDocumentIssuers(), args)
}

func ExCmdGetDocIssuer(clientCtx client.Context, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDocumentIssuer(), args)
}

func ExCmdEnrollAccountOperator(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {

	args := append(address, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdEnrollAccountOperators(), args)
}
func ExCmdQueryAccountOperator(clientCtx client.Context, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdAccountOperator(), args)
}
func ExCmdRevokeAccountOperator(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {

	args := append(address, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeAccountOperators(), args)
}

func ExCmdEnrollIdSigner(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := append(address, additionalFlags...)

	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdEnrollIdSigners(), args)
}
func ExCmdRevokeIdSigner(clientCtx client.Context, address []string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := append(address, additionalFlags...)

	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeIdSigners(), args)
}
func ExCmdGetIdSigner(clientCtx client.Context, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdIdSigner(), args)
}

func ExCmdEnrollLoader(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdEnrollLoaders(), args)
}
func ExCmdRevokeLoader(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeLoaders(), args)
}
func ExCmdGetLoader(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdGetLoader(), args)
}

func ExCmdEnrollVoters(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{address}
	args = append(args, additionalFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdEnrollVoter(), args)
}

//func ExCmdRevokeVoter(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
//	args := []string{address}
//	args = append(args, additionalFlags...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRevokeVoter(), args)
//}
//func ExCmdGetVoter(clientCtx client.Context, t *testing.T, address string, additionalFlags ...string) (testutil.BufferWriter, error) {
//	args := []string{address}
//	args = append(args, additionalFlags...)
//	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdGetVoter(), args)
//}

//func JsonVoterUnmarshal(t *testing.T, bz []byte) types.QueryGetLoaderResponse {
//	var a types.QueryGetLoaderResponse
//	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
//	err := encCfg.Marshaler.UnmarshalJSON(bz, &a)
//	require.NoError(t, err)
//	return a
//}

func JsonAccountOperatorUnmarshal(t *testing.T, bz []byte) types.QueryAccountOperatorResponse {
	var a types.QueryAccountOperatorResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(bz, &a)
	require.NoError(t, err)
	return a
}
func JsonDocIssuerUnmarshal(t *testing.T, bz []byte) types.QueryDocumentIssuerResponse {
	var a types.QueryDocumentIssuerResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(bz, &a)
	require.NoError(t, err)
	return a
}
func JsonIDSignerUnmarshal(t *testing.T, bz []byte) types.QueryIdSignerResponse {
	var a types.QueryIdSignerResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(bz, &a)
	require.NoError(t, err)
	return a
}

//func JsonLoaderUnmarshal(t *testing.T, bz []byte) types.QueryGetLoaderResponse {
//	var a types.QueryGetLoaderResponse
//	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
//	err := encCfg.Marshaler.UnmarshalJSON(bz, &a)
//	require.NoError(t, err)
//	return a
//}
