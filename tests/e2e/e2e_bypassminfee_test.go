package e2e

import (
	"cosmossdk.io/math"
)

func (s *IntegrationTestSuite) testByPassMinFeeWithdrawReward() {
	paidFeeAmt := math.LegacyMustNewDecFromStr(minGasPrice).Mul(math.LegacyNewDec(gas)).String()
	payee, _ := s.chainA.validators[0].keyInfo.GetAddress()
	// pass
	s.T().Logf("bypass-msg with fee in the denom of global fee, pass")
	s.execWithdrawAllRewards(s.chainA, 0, payee.String(), paidFeeAmt+denom, false)
	// pass
	s.T().Logf("bypass-msg with low fee coin in the denom of global fee, pass")
	s.execWithdrawAllRewards(s.chainA, 0, payee.String(), "1"+denom, false)
	// fail
	s.T().Logf("bypass-msg with non-zero coin not in the denom of global fee, fail")
	s.execWithdrawAllRewards(s.chainA, 0, payee.String(), paidFeeAmt+photonDenom, true)
}
