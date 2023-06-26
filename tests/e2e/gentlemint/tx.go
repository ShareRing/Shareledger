package gentlemint

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"

	"github.com/sharering/shareledger/x/gentlemint/client/cli"
	"github.com/sharering/shareledger/x/utils/denom"
)

func (s *E2ETestSuite) TestSetActionLevelFee() {
	testCases := tests.TestCasesTx{
		{
			Name: "set action level fee",
			Args: []string{
				defaultActionLevelFees[2].Action,
				"high",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "set invalid action level fee",
			Args: []string{
				"test",
				"high",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
		{
			Name: "set invalid action level fee",
			Args: []string{
				defaultActionLevelFees[3].Action,
				"test",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdSetActionLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestDeleteActionLevelFee() {
	testCases := tests.TestCasesTx{
		{
			Name: "delete action level fee with invalid user",
			Args: []string{
				defaultActionLevelFees[3].Action,
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "delete action level fee ok",
			Args: []string{
				defaultActionLevelFees[3].Action,
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "delete invalid action level fee",
			Args: []string{
				"test",
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdDeleteActionLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestBurn() {
	testCases := tests.TestCasesTx{
		{
			Name: "burn coin",
			Args: []string{
				"1000shr",
				network.MakeByAccount(network.KeyTreasurer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "burn coin with permissionless signer",
			Args: []string{
				"100shr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "burn coin with invalid coin",
			Args: []string{
				"100test",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdBurn(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestBuyShr() {
	testCases := tests.TestCasesTx{
		{
			Name: "buy shr",
			Args: []string{
				"100",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "buy shr with insufficient funds",
			Args: []string{
				"100",
				network.MakeByAccount(network.KeyEmpty1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrInsufficientFunds.ABCICode(),
		},

		{
			Name: "buy invalid coin type",
			Args: []string{
				"100test",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdBuyShr(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestSetLevelFee() {
	testCases := tests.TestCasesTx{
		{
			Name: "set level fee",
			Args: []string{
				"high",
				sdk.NewDecCoinFromDec(denom.ShrP, sdk.MustNewDecFromStr("0.05")).String(),
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "set level fee with permissionless signer",
			Args: []string{
				"new",
				"10shrp",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "set level fee with invalid coin",
			Args: []string{
				"new",
				"10shr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdSetLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestLoad() {
	testAcc := network.MustAddressFormKeyring(s.network.Validators[0].ClientCtx.Keyring, network.KeyAccount2)
	testCases := tests.TestCasesTx{
		{
			Name: "load coin",
			Args: []string{
				testAcc.String(),
				"100shr",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "load coin with permissionless user",
			Args: []string{
				testAcc.String(),
				"100shr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
		{
			Name: "load coin with invalid user",
			Args: []string{
				"test",
				"100shr",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdLoad(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestSend() {
	testAcc := network.MustAddressFormKeyring(s.network.Validators[0].ClientCtx.Keyring, network.KeyAccount2)
	testCases := tests.TestCasesTx{
		{
			Name: "send coin",
			Args: []string{
				testAcc.String(),
				"100shr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "send coin with invalid coin",
			Args: []string{
				testAcc.String(),
				"100test",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
		{
			Name: "send coin with invalid user",
			Args: []string{
				"invaliduser",
				"100shr",
				network.MakeByAccount(network.KeyAccount1),
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdSend(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestSetExchange() {
	testCases := tests.TestCasesTx{
		{
			Name: "set exchange",
			Args: []string{
				"400",
				network.MakeByAccount(network.KeyTreasurer),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
		{
			Name: "set exchange with permissionless signer",
			Args: []string{
				"500",
				network.MakeByAccount(network.KeyAuthority),
			},
			ExpectErr:    false,
			ExpectedCode: sdkerrors.ErrUnauthorized.ABCICode(),
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdSetExchange(), s.network.Validators[0])
}
