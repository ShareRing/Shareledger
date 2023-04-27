package e2e

var (
	runBankTest                   = true
	runBypassMinFeeTest           = true
	runEncodeTest                 = true
	runEvidenceTest               = true
	runFeeGrantTest               = true
	runGlobalFeesTest             = true
	runGovTest                    = true
	runIBCTest                    = true
	runSlashingTest               = false
	runStakingAndDistributionTest = false
	runVestingTest                = false
)

func (s *IntegrationTestSuite) TestBank() {
	if !runBankTest {
		s.T().Skip()
	}
	s.testBankTokenTransfer()
}

func (s *IntegrationTestSuite) TestByPassMinFee() {
	if !runBypassMinFeeTest {
		s.T().Skip()
	}
	s.testByPassMinFeeWithdrawReward()
}

func (s *IntegrationTestSuite) TestEncode() {
	if !runEncodeTest {
		s.T().Skip()
	}
	s.testEncode()
	s.testDecode()
}

func (s *IntegrationTestSuite) TestEvidence() {
	if !runEvidenceTest {
		s.T().Skip()
	}
	s.testEvidence()
}

func (s *IntegrationTestSuite) TestFeeGrant() {
	if !runFeeGrantTest {
		s.T().Skip()
	}
	s.testFeeGrant()
}

func (s *IntegrationTestSuite) TestGlobalFees() {
	if !runGlobalFeesTest {
		s.T().Skip()
	}
	s.testGlobalFees()
	s.testQueryGlobalFeesInGenesis()
}

func (s *IntegrationTestSuite) TestGov() {
	if !runGovTest {
		s.T().Skip()
	}
	s.GovSoftwareUpgrade()
	s.GovCancelSoftwareUpgrade()
}

func (s *IntegrationTestSuite) TestIBC() {
	if !runIBCTest {
		s.T().Skip()
	}
	s.testIBCTokenTransfer()
}
