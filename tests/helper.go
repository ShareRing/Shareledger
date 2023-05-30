package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type TestCase struct {
	Name      string
	Args      []string
	ExpectErr bool
	RespType  proto.Message
	Expected  proto.Message
}

type TestCases = []TestCase

func RunTestCases(s *suite.Suite, tcs TestCases, cmd *cobra.Command, val *network.Validator) {
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.Args)
			if tc.ExpectErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.RespType))
				s.Equal(tc.Expected.String(), tc.RespType.String())
			}
		})
	}
}

type TestCaseGrpc struct {
	Name      string
	URL       string
	Headers   map[string]string
	ExpectErr bool
	RespType  proto.Message
	Expected  proto.Message
}

type TestCasesGrpc = []TestCaseGrpc

func RunTestCasesGrpc(s *suite.Suite, tcs TestCasesGrpc, val *network.Validator) {
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			resp, err := testutil.GetRequestWithHeaders(tc.URL, tc.Headers)
			s.NoError(err)
			if tc.ExpectErr {
				s.Error(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
			} else {
				s.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
				s.Equal(tc.Expected.String(), tc.RespType.String())
			}
		})
	}
}

type TestCaseTx struct {
	Name         string
	Args         []string
	ExpectErr    bool
	ExpectedCode uint32
}

type TestCasesTx = []TestCaseTx

func RunTestCasesTx(s *suite.Suite, tcs TestCasesTx, cmd *cobra.Command, val *network.Validator) {
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.Args)
			if tc.ExpectErr {
				s.Error(err)
			} else {
				s.NoError(err)
				var resp types.TxResponse
				s.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				s.Equal(tc.ExpectedCode, resp.Code)
			}
		})
	}
}
