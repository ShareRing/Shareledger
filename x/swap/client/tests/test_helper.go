package tests

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/sharering/shareledger/x/swap/client/cli"
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

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdFund(), extraFlags)
}
