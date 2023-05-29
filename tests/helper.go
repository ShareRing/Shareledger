package tests

import (
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Name      string
	Args      []string
	ExpectErr bool
	RespType  proto.Message
	Expected  proto.Message
}

type TestCases = []TestCase

func RunTestCases(as *require.Assertions, tcs TestCases, cmd *cobra.Command, val *network.Validator) {
	for _, tc := range tcs {
		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.Args)
		if tc.ExpectErr {
			as.Error(err)
		} else {
			as.NoError(err)
			as.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.RespType))
			as.Equal(tc.Expected.String(), tc.RespType.String())
		}
	}
}
