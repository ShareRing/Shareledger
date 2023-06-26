package tests

import (
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"
	shareledgernetwork "github.com/sharering/shareledger/testutil/network"
	"github.com/spf13/cobra"

	"reflect"

	"github.com/stretchr/testify/suite"
)

type TestCase struct {
	Name      string
	Args      []string
	ExpectErr bool
	RespType  proto.Message
	Expected  proto.Message
}

const (
	DefaultNumRetry = 100
	RetryTimeGap    = 100 * time.Millisecond
)

func retryGap() {
	time.Sleep(RetryTimeGap)
}

type TestCases = []TestCase

func RunTestCases(s *suite.Suite, tcs TestCases, cmd *cobra.Command, val *network.Validator) {
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			tc.Args = append(tc.Args, shareledgernetwork.JSONFlag)
			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.Args)
			if tc.ExpectErr {
				s.Error(err)
			} else {
				s.NoError(err)
				s.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.RespType))
				checkRespType(s, tc.RespType, tc.Expected)
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

func checkRespType(s *suite.Suite, resp, expected interface{}) {
	typeOfResp := reflect.ValueOf(resp).Elem()
	typeOfExpected := reflect.ValueOf(expected).Elem()
	for i := 0; i < typeOfResp.NumField(); i++ {
		exp := typeOfExpected.Field(i)
		res := typeOfResp.Field(i)
		if res.Kind() == reflect.Slice || res.Kind() == reflect.Array {
			// currently only check for array/slice of first level of the struct field
			for j := 0; j < exp.Len(); j++ {
				s.Contains(res.Interface(), exp.Index(j).Interface())
			}
		} else {
			s.Equal(exp.Interface(), res.Interface())

		}
	}
}

func RunTestCasesGrpc(s *suite.Suite, tcs TestCasesGrpc, val *network.Validator) {
	for _, tc := range tcs {
		s.Run(tc.Name, func() {
			resp, err := testutil.GetRequestWithHeaders(tc.URL, tc.Headers)
			s.NoError(err)
			if tc.ExpectErr {
				s.Error(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
			} else {
				s.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.RespType))
				checkRespType(s, tc.RespType, tc.Expected)
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
			resp, err := RunCmdWithRetry(s, cmd, val, tc.Args, DefaultNumRetry)
			if tc.ExpectErr {
				s.Error(err)
			} else {
				s.NoError(err)
				// code !=0 mean that this tx failed on CheckTx call (ante step) => the tx is not committed
				if resp.Code != 0 {
					s.Equal(tc.ExpectedCode, resp.Code)
					return
				}
				resp, err = QueryTxWithRetry(val.ClientCtx, resp.TxHash, DefaultNumRetry)
				s.NoError(err)
				s.Equal(tc.ExpectedCode, resp.Code)
			}
		})
	}
}

// QueryTxWithRetry wait for tx `hashHexStr` to success
func QueryTxWithRetry(clientCtx client.Context, hashHexStr string, retry int) (*sdk.TxResponse, error) {
	resp, err := authtx.QueryTx(clientCtx, hashHexStr)
	if retry == 0 || err == nil {
		return resp, err
	}
	// take a nap before each retries
	retryGap()
	return QueryTxWithRetry(clientCtx, hashHexStr, retry-1)
}

// RunCmdWithRetry send tx and auto retry if it return ErrWrongSequence
func RunCmdWithRetry(s *suite.Suite, cmd *cobra.Command, val *network.Validator, args []string, retry int) (*types.TxResponse, error) {
	args = append(args, DefaultTxFlag()...)
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, args)
	if err != nil {
		return nil, err
	}
	var resp types.TxResponse
	s.NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &resp))
	if resp.Code == sdkerrors.ErrWrongSequence.ABCICode() {
		if retry == 0 {
			return nil, errors.New("Exceeded max retried time")
		}
		// take a nap before each retries
		retryGap()
		return RunCmdWithRetry(s, cmd, val, args, retry-1)
	}
	return &resp, nil
}

// RunCmdBlock auto retry on error and wait for tx committed
func RunCmdBlock(s *suite.Suite, cmd *cobra.Command, val *network.Validator, args []string) (*types.TxResponse, error) {
	resp, err := RunCmdWithRetry(s, cmd, val, args, 100)
	if err != nil {
		return resp, err
	}
	return QueryTxWithRetry(val.ClientCtx, resp.TxHash, 100)
}

func DefaultTxFlag() []string {
	return []string{
		shareledgernetwork.JSONFlag,
		shareledgernetwork.SkipConfirmation,
		shareledgernetwork.SyncBroadcast,
		"--gas=2000000",          // make sure that we always have enough gas
		"--fees=10000000000nshr", // fees always is 10SHR
	}
}

func RunQueryCmd(val *network.Validator, cmd *cobra.Command, args []string, resp codec.ProtoMarshaler) error {
	args = append(args, shareledgernetwork.JSONFlag)
	out, err := clitestutil.ExecTestCLICmd(
		val.ClientCtx,
		cmd,
		args)
	if err != nil {
		return err
	}
	return val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), resp)
}
