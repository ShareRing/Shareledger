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
