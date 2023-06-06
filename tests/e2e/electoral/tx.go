package electoral

import (
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/electoral/client/cli"
)

func (s *E2ETestSuiteTx) TestCmdEnrollVoter() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll voter",
			Args:      []string{accVoter.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "enroll voter empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll voter not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollVoter(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeVoter() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke voter",
			Args:      []string{accVoter.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "revoke voter empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "revoke voter not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeVoter(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdEnrollLoaders() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll loader",
			Args:      []string{accKeyShrpLoaders.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "enroll loader empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll loader not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollLoaders(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeLoaders() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke loader",
			Args:      []string{accKeyShrpLoaders.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "revoke loader empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "revoke loader not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeLoaders(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdEnrollDocIssuers() {
	s.T().Log(network.Accounts[network.KeyOperator].String())
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll docissuer",
			Args:      []string{accDocIssuer.Address, network.MakeByAccount(network.KeyOperator)},
			ExpectErr: false,
		},
		{
			Name:      "enroll docissuer empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll docissuer not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollDocIssuers(), s.network.Validators[0])
}

// func (s *E2ETestSuiteTx) TestCmdRevokeDocIssuers() {
// 	testCasesTx := tests.TestCasesTx{
// 		{
// 			Name:      "revoke docissuer",
// 			Args:      []string{accDocIssuer.Address, network.MakeByAccount(network.KeyOperator)},
// 			ExpectErr: false,
// 		},
// 		{
// 			Name:      "revoke docissuer empty",
// 			Args:      []string{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke docissuer not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeDocIssuers(), s.network.Validators[0])
// }

func (s *E2ETestSuiteTx) TestCmdEnrollAccountOperators() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll account operator",
			Args:      []string{network.MakeByAccount(network.KeyAuthority), accOp.Address},
			ExpectErr: false,
		},
		{
			Name:      "enroll account operator empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll account operator not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollAccountOperators(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeAccountOperators() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke account operator",
			Args:      []string{accOp.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "revoke account operator empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "revoke account operator not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeAccountOperators(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdEnrollRelayers() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll relayer",
			Args:      []string{accKeyRelayer.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "enroll relayer empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll relayer not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollRelayers(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeRelayers() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke relayer",
			Args:      []string{accKeyRelayer.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "revoke relayer empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "revoke relayer not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeRelayers(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdEnrollApprovers() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll approver",
			Args:      []string{accApprover.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "enroll approver empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "enroll approver not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollApprovers(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeApprover() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke approver",
			Args:      []string{accApprover.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "revoke approver empty",
			Args:      []string{},
			ExpectErr: true,
		},
		{
			Name: "revoke approver not exists",
			Args: []string{
				"notExistsAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeApprover(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdEnrollSwapManagers() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "enroll swap manager",
			Args:      []string{accSwapManager.Address, network.MakeByAccount(network.KeyAuthority)},
			ExpectErr: false,
		},
		{
			Name:      "enroll swap manager empty",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdEnrollSwapManagers(), s.network.Validators[0])
}

func (s *E2ETestSuiteTx) TestCmdRevokeSwapManagers() {
	testCasesTx := tests.TestCasesTx{
		{
			Name:      "revoke swap manager",
			Args:      []string{network.MakeByAccount(network.KeyAuthority), accSwapManager.Address},
			ExpectErr: false,
		},
		{
			Name:      "revoke swap manager empty",
			Args:      []string{},
			ExpectErr: true,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCasesTx, cli.CmdRevokeSwapManagers(), s.network.Validators[0])
}
