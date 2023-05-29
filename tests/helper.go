package tests

import (
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
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
	assert := s.Assert()
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.Args)
			if tc.ExpectErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.RespType))
				assert.Equal(tc.Expected.String(), tc.RespType.String())
			}
		})
	}
}

type TestCaseGrpc struct {
	Name     string
	URL      string
	Headers  map[string]string
	ExpErr   bool
	RespType proto.Message
	Expected proto.Message
}

type TestCasesGrpc = []TestCaseGrpc

func RunTestCasesGrpc(s *suite.Suite, tcs TestCasesGrpc, val *network.Validator) {
	assert := s.Assert()
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			resp, err := testutil.GetRequestWithHeaders(tc.URL, tc.Headers)
			assert.NoError(err)
			if tc.ExpErr {
				assert.Error(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
			} else {
				assert.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
				assert.Equal(tc.Expected.String(), tc.RespType.String())
			}
		})
	}
}
