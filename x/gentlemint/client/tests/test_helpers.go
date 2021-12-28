package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cli2 "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	testutil2 "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/sharering/shareledger/x/gentlemint/client/cli"
)

func CmdGetExchangeRate(ctx client.Context, flags ...string) (testutil.BufferWriter, error) {
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdShowExchangeRate(), flags)
}

func CmdSetExchangeRate(ctx client.Context, rate string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{rate}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdSetExchange(), args)
}

//CmdBurnSHR burn shr by treasurer account send shr to void
func CmdBurnSHR(ctx client.Context, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdBurnShr(), args)
}

//CmdBurnSHRP burn shrp by treasurer account send shr to void
func CmdBurnSHRP(ctx client.Context, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdBurnShrp(), args)
}

//CmdBuyCent buy cent by shrp
func CmdBuyCent(ctx client.Context, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdBuyCent(), args)
}

//CmdBuySHR buy shr by SHRP and cent
func CmdBuySHR(ctx client.Context, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdBuyShr(), args)
}

//CmdLoadSHR mint new shr coin out of thin air and send it to address require authority
func CmdLoadSHR(ctx client.Context, receiver, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{receiver, amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdLoadShr(), args)
}

//CmdLoadSHRP mint new shrp coin out of thin air and send it to address require SHRP loader role
func CmdLoadSHRP(ctx client.Context, receiver, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{receiver, amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdLoadShrp(), args)
}

//CmdSendSHR send shr to address from address
func CmdSendSHR(ctx client.Context, receiver, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{receiver, amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdSendShr(), args)
}

//CmdSendSHRP send shrp to address from address
func CmdSendSHRP(ctx client.Context, receiver, amount string, flags ...string) (testutil.BufferWriter, error) {
	var args = []string{receiver, amount}
	args = append(args, flags...)
	return clitestutil.ExecTestCLICmd(ctx, cli.CmdSendShrp(), args)
}

//CmdTotalSupply
func CmdTotalSupply(ctx client.Context, flags ...string) (testutil.BufferWriter, error) {

	return clitestutil.ExecTestCLICmd(ctx, cli2.GetCmdQueryTotalSupply(), flags)
}

func CmdQueryBalance(t *testing.T, ctx client.Context, address sdk.Address) types.QueryAllBalancesResponse {
	balRes := types.QueryAllBalancesResponse{}
	out, err := testutil2.QueryBalancesExec(ctx, address)
	if err != nil {
		t.Errorf("query balance fail %v", err)
	}
	err = ctx.Codec.UnmarshalJSON(out.Bytes(), &balRes)
	if err != nil {
		t.Errorf("unmarshal balance fail %v", err)
	}
	return balRes
}
