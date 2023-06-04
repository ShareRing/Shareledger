package electoral

// func (s *E2ETestSuite) TestCmdEnrollVoter() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll voter",
// 			Args:      []string{accVoter.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryVoterResponse{},
// 			Expected:  &types.QueryVoterResponse{},
// 		},
// 		{
// 			Name:      "enroll voter empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryVoterResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll voter not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryVoterResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollVoter(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeVoter() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke voter",
// 			Args:      []string{accVoter.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryVoterResponse{},
// 			Expected:  &types.QueryVoterResponse{},
// 		},
// 		{
// 			Name:      "revoke voter empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryVoterResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke voter not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryVoterResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeVoter(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollLoaders() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll loader",
// 			Args:      []string{accKeyShrpLoaders.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryLoaderResponse{},
// 			Expected:  &types.QueryLoaderResponse{},
// 		},
// 		{
// 			Name:      "enroll loader empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll loader not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollLoaders(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollLoadersFromFile() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll loader",
// 			Args:      []string{"testdata/loader.txt"},
// 			ExpectErr: false,
// 			RespType:  &types.QueryLoaderResponse{},
// 			Expected:  &types.QueryLoaderResponse{},
// 		},
// 		{
// 			Name:      "enroll loader empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll loader not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollLoadersFromFile(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeLoaders() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke loader",
// 			Args:      []string{accKeyShrpLoaders.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryLoaderResponse{},
// 			Expected:  &types.QueryLoaderResponse{},
// 		},
// 		{
// 			Name:      "revoke loader empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke loader not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeLoaders(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeLoadersFromFile() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke loader",
// 			Args:      []string{"testdata/loader.txt"},
// 			ExpectErr: false,
// 			RespType:  &types.QueryLoaderResponse{},
// 			Expected:  &types.QueryLoaderResponse{},
// 		},
// 		{
// 			Name:      "revoke loader empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke loader not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryLoaderResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeLoadersFromFile(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollIdSigners() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll idsigner",
// 			Args:      []string{accIDSigner.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryIdSignerResponse{},
// 			Expected:  &types.QueryIdSignerResponse{},
// 		},
// 		{
// 			Name:      "enroll idsigner empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll idsigner not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollIdSigners(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollIdSignersFromFile() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll idsigner",
// 			Args:      []string{"testdata/idsigner.txt"},
// 			ExpectErr: false,
// 			RespType:  &types.QueryIdSignerResponse{},
// 			Expected:  &types.QueryIdSignerResponse{},
// 		},
// 		{
// 			Name:      "enroll idsigner empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll idsigner not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollIdSignerFromFile(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeIdSigners() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke idsigner",
// 			Args:      []string{accIDSigner.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryIdSignerResponse{},
// 			Expected:  &types.QueryIdSignerResponse{},
// 		},
// 		{
// 			Name:      "revoke idsigner empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke idsigner not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryIdSignerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeIdSigners(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollDocIssuers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll docissuer",
// 			Args:      []string{accDocIssuer.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			Expected: &types.QueryDocumentIssuerResponse{
// 				AccState: &types.AccState{
// 					Address: accDocIssuer.Address,
// 				},
// 			},
// 		},
// 		{
// 			Name:      "enroll docissuer empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll docissuer not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollDocIssuers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeDocIssuers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke docissuer",
// 			Args:      []string{accDocIssuer.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			Expected:  &types.QueryDocumentIssuerResponse{},
// 		},
// 		{
// 			Name:      "revoke docissuer empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke docissuer not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryDocumentIssuerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeDocIssuers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollAccountOperators() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll account operator",
// 			Args:      []string{accOp.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			Expected: &types.QueryAccountOperatorResponse{
// 				AccState: &types.AccState{
// 					Address: accOp.Address,
// 				},
// 			},
// 		},
// 		{
// 			Name:      "enroll account operator empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll account operator not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollAccountOperators(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeAccountOperators() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke account operator",
// 			Args:      []string{accOp.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			Expected:  &types.QueryAccountOperatorResponse{},
// 		},
// 		{
// 			Name:      "revoke account operator empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke account operator not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryAccountOperatorResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeAccountOperators(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollRelayers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll relayer",
// 			Args:      []string{accKeyRelayer.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryRelayerResponse{},
// 			Expected: &types.QueryRelayerResponse{
// 				AccState: types.AccState{
// 					Address: accKeyRelayer.Address,
// 				},
// 			},
// 		},
// 		{
// 			Name:      "enroll relayer empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryRelayerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll relayer not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryRelayerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollRelayers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeRelayers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke relayer",
// 			Args:      []string{accKeyRelayer.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryRelayerResponse{},
// 			Expected:  &types.QueryRelayerResponse{},
// 		},
// 		{
// 			Name:      "revoke relayer empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryRelayerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke relayer not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryRelayerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeRelayers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollApprovers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll approver",
// 			Args:      []string{accApprover.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryApproverResponse{},
// 			Expected: &types.QueryApproverResponse{
// 				AccState: types.AccState{
// 					Address: accApprover.Address,
// 				},
// 			},
// 		},
// 		{
// 			Name:      "enroll approver empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryApproverResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll approver not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryApproverResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollApprovers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeApprover() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke approver",
// 			Args:      []string{accApprover.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QueryApproverResponse{},
// 			Expected:  &types.QueryApproverResponse{},
// 		},
// 		{
// 			Name:      "revoke approver empty",
// 			Args:      []string{},
// 			RespType:  &types.QueryApproverResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke approver not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QueryApproverResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeApprover(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdEnrollSwapManagers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "enroll swap manager",
// 			Args:      []string{accSwapManager.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			Expected: &types.QuerySwapManagerResponse{
// 				AccState: types.AccState{
// 					Address: accSwapManager.Address,
// 				},
// 			},
// 		},
// 		{
// 			Name:      "enroll swap manager empty",
// 			Args:      []string{},
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "enroll swap manager not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdEnrollSwapManagers(), s.network.Validators[0])
// }

// func (s *E2ETestSuite) TestCmdRevokeSwapManagers() {
// 	testCases := tests.TestCases{
// 		{
// 			Name:      "revoke swap manager",
// 			Args:      []string{accSwapManager.Address},
// 			ExpectErr: false,
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			Expected:  &types.QuerySwapManagerResponse{},
// 		},
// 		{
// 			Name:      "revoke swap manager empty",
// 			Args:      []string{},
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			ExpectErr: true,
// 		},
// 		{
// 			Name: "revoke swap manager not exists",
// 			Args: []string{
// 				"notExistsAddress",
// 			},
// 			RespType:  &types.QuerySwapManagerResponse{},
// 			ExpectErr: true,
// 		},
// 	}
// 	tests.RunTestCases(&s.Suite, testCases, cli.CmdRevokeSwapManagers(), s.network.Validators[0])
// }
