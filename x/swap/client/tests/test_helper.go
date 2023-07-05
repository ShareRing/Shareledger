package tests

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/sharering/shareledger/x/swap/client/cli"
	"github.com/sharering/shareledger/x/swap/types"
)

func CmdDeposit(clientCtx client.Context, amount string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{amount}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdDeposit(), args)
}

func CmdWithdraw(clientCtx client.Context, receiver, amount string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{receiver, amount}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdWithdraw(), args)
}

func CmdFund(clientCtx client.Context, extraFlags ...string) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdBalance(), extraFlags)
}

func CmdCancel(clientCtx client.Context, ids string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{ids}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCancel(), args)
}
func CmdOut(clientCtx client.Context, dest, network, amount string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{dest, network, amount}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdOut(), args)
}

func CmdApprove(clientCtx client.Context, signName, reqIds, netName string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{reqIds, signName, netName}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdApprove(), args)
}

func CmdSearch(clientCtx client.Context, status string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{status}
	args = append(args, extraFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdRequests(), args)
}

func ParseSearchResponse(cdc codec.Codec, sRes []byte) (types.QuerySwapResponse, error) {
	res := types.QuerySwapResponse{}
	err := cdc.Unmarshal(sRes, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
func CmdReject(clientCtx client.Context, sign, reqIds string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{sign, reqIds}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdReject(), args)
}

func CmdCreateFeeSchema(clientCtx client.Context, network, data, in, out, ce string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{network, data, in, out, ce}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateSchema(), args)
}

func CmdGetSchema(clientCtx client.Context, net string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{net}
	args = append(args, extraFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdShowSchema(), args)
}

func CmdGetBatches(clientCtx client.Context, extraFlags ...string) (testutil.BufferWriter, error) {

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdBatches(), extraFlags)
}
