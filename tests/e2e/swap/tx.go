package swap

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/swap/client/cli"
)

func (s *E2ETestSuite) TestCreateRequestOut() {
	testCases := tests.TestCasesTx{
		{
			Name: "given request greater than swap out fee and valid network should be success",
			Args: []string{
				"0xDest1",
				"eth",
				"100000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given request amount less than swap out fee and valid network should be success",
			Args: []string{
				"0xDest1",
				"eth",
				"30000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given swap out request, not exist network",
			Args: []string{
				"0xDest1",
				"eth222",
				"100000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "given swap out request, network exist without fee",
			Args: []string{
				"0xDest1",
				"schemaWithoutFee",
				"100000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "given invalid amount coin should be error",
			Args: []string{
				"0xDest1",
				"schemaWithoutFee",
				"1a0x0atom",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "given swap out transaction but fail validate basic without creator address",
			Args: []string{
				"0xDest1",
				"schemaWithoutFee",
				"1a0x0atom",
			},
			ExpectErr: true,
		},
		{
			Name: "given swap out transaction but fail validate basic not support coin",
			Args: []string{
				"0xDest1",
				"schemaWithoutFee",
				"100atom",
			},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdOut(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCreateRequestIn() {
	testCases := tests.TestCasesTx{
		{
			Name: "create request in valid should be success",
			Args: []string{
				"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
				"eth",
				"hash1:sender1:12",
				"1000000000nshr",
				"20000000000nshr",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create request in duplication tx hashes should be error",
			Args: []string{
				"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
				"eth",
				"hash1:sender1:12,hash1:sender1:12",
				"1000000000nshr",
				"20000000000nshr",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr: true,
		},
		{
			Name: "create request in valid should be success",
			Args: []string{
				"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
				"eth",
				"hash1:sender1:12",
				"1000000000nshr",
				"20000000000nshr",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create request in not valid log index should be success",
			Args: []string{
				"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
				"eth",
				"hash1:sender1:1xx2",
				"1000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "create request in but transaction singer isn't relayer account",
			Args: []string{
				"shareledger1tc9yej24s698vm5w0jvt7452mgxl3ck74nqxjg",
				"eth",
				"hash1:sender1:12",
				"1000000000nshr",
				"20000000000nshr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdIn(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCancelRequestOut() {
	testCases := tests.TestCasesTx{
		{
			Name: "given the request id correct cancel request successfully",
			Args: []string{
				"1",
				network.MakeByAccount(network.KeyAccount1),
				network.SHRFee2,
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given request id not number should be error",
			Args: []string{
				"two",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "given cancel request but creator address is empty",
			Args: []string{
				"2",
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCancel(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestRejectRequestOut() {
	testCases := tests.TestCasesTx{
		{
			Name: "reject the swap out request",
			Args: []string{
				"2",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given request id not number should be error",
			Args: []string{
				"two",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "given cancel request but creator address is empty",
			Args: []string{
				"2",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "given approved request can't reject",
			Args: []string{
				"3",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrInvalidRequest.ABCICode(),
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdReject(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCancelBatch() {
	testCases := tests.TestCasesTx{
		{
			Name: "cancel batch success",
			Args: []string{
				"2",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "cancel batch but not authorizer",
			Args: []string{
				"2",
				network.MakeByAccount(network.KeyAccount2),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "cancel batch id isn't number success",
			Args: []string{
				"two",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr: true,
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCancelBatches(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCompleteBatch() {
	testCases := tests.TestCasesTx{
		{
			Name: "complete batch success",
			Args: []string{
				"3",
				network.MakeByAccount(network.KeyApproverRelayer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "complete batch but not authorize",
			Args: []string{
				"3",
				network.MakeByAccount(network.KeyAccount3),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}

	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCompleteBatch(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestWithdraw() {
	testCases := tests.TestCasesTx{
		{
			Name: "given valid withdraw message should be success",
			Args: []string{
				network.Accounts[network.KeyAccount1].String(),
				"10shr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given invalid receiver withdraw message should be fail",
			Args: []string{
				"0x1222",
				"10shr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
		{
			Name: "given invalid coin amount withdraw message should be fail",
			Args: []string{
				network.Accounts[network.KeyAccount1].String(),
				"1oshr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},

		{
			Name: "given amount greeter than module balance withdraw message should be fail",
			Args: []string{
				network.Accounts[network.KeyAccount1].String(),
				"122000000shr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrInvalidRequest.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdWithdraw(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestDeposit() {
	testCases := tests.TestCasesTx{
		{
			Name: "given valid deposit message should be success",
			Args: []string{
				"10shr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "given the depositor not enough coin",
			Args: []string{
				"10shr",
				network.MakeByAccount(network.KeyEmpty3),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrInsufficientFunds.ABCICode(),
		},
		{
			Name: "given invalid coin amount withdraw message should be fail",
			Args: []string{
				"1oshr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdDeposit(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCreateSchema() {
	testCases := tests.TestCasesTx{
		{
			Name: "create the valid schema should be success",
			Args: []string{
				"newNet",
				"{}",
				"10shr",
				"10shr",
				"8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "create the valid schema but creator isn't authority should got error",
			Args: []string{
				"newNet",
				"{}",
				"10shr",
				"10shr",
				"8",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "given invalid coins should be fail",
			Args: []string{
				"newNet",
				"{}",
				"1x0shr",
				"10shr",
				"8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
		{
			Name: "given invalid contract exponent should be fail",
			Args: []string{
				"newNet",
				"{}",
				"10shr",
				"10shr",
				"8x",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
		{
			Name: "given existed contract should be fail",
			Args: []string{
				"eth",
				"{}",
				"10shr",
				"10shr",
				"8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrInvalidRequest.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateSchema(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestUpdateSchema() {
	testCases := tests.TestCasesTx{
		{
			Name: "update valid schema should be success",
			Args: []string{
				"hero",
				"--schema={}",
				"--fee-in=10shr",
				"--fee-out=10shr",
				"--exp=8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "update the valid schema but creator isn't authority should got error",
			Args: []string{
				"hero",
				"--schema={}",
				"--fee-in=10shr",
				"--fee-out=10shr",
				"--exp=8",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "given invalid coins should be fail",
			Args: []string{
				"hero",
				"--schema={}",
				"--fee-in=tenshr",
				"--fee-out=tenshr",
				"--exp=8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
		{
			Name: "given invalid contract exponent should be fail",
			Args: []string{
				"hero",
				"--schema={}",
				"--fee-in=10shr",
				"--fee-out=10shr",
				"--exp=8x",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
		{
			Name: "update not exist schema should be fail",
			Args: []string{
				"ethxxx",
				"--schema={}",
				"--fee-in=10shr",
				"--fee-out=10shr",
				"--exp=8",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdateSchema(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestDeleteSchema() {
	testCases := tests.TestCasesTx{
		{
			Name: "delete exist schema should be success",
			Args: []string{
				"hero1",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "delete schema but not authority should be fail",
			Args: []string{
				"hero1",
				network.MakeByAccount(network.KeyAccount1),
			},

			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "delete not exist schema should be fail",
			Args: []string{
				"hellosharering",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdDeleteSchema(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestUpdateSwappingFee() {
	testCases := tests.TestCasesTx{
		{
			Name: "Update swapping success",
			Args: []string{
				"hero",
				"--fee-in=10shr",
				"--fee-out=10shr",
				network.MakeByAccount(network.KeyTreasurer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "Update swapping fee but un supported coins",
			Args: []string{
				"hero",
				"--fee-in=10mhr",
				"--fee-out=10mshr",
				network.MakeByAccount(network.KeyTreasurer),
			},
			ExpectErr: true,
		},
		{
			Name: "Update swapping fee with not exist network",
			Args: []string{
				"no_exited",
				"--fee-in=10shr",
				"--fee-out=10shr",
				network.MakeByAccount(network.KeyTreasurer),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
		{
			Name: "Update swapping fee but not authorize",
			Args: []string{
				"no_exited",
				"--fee-in=10shr",
				"--fee-out=10shr",
				network.MakeByAccount(network.KeyAccount3),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdUpdateSwapFee(), s.network.Validators[0])
}
