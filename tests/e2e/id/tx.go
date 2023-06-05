package id

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/id/client/cli"
)

func (s *E2ETestSuite) TestClientId() {
	testCases := tests.TestCasesTx{{
		Name: "create new id from id signer account",
		Args: []string{
			"IDID",
			"shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v",
			"shareledger18pf3zdwqjntd9wkvfcjvmdc7hua6c0q2eck5h5",
			"ExtraData",
			network.MakeByAccount(network.KeyIDSigner),
		},
		ExpectErr:    false,
		ExpectedCode: errorsmod.SuccessABCICode,
	}}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.CmdCreateId(), s.network.Validators[0])
}
