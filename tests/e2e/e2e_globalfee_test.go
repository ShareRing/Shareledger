package e2e

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// globalfee in genesis is set to be initialGlobalFeeAmt
func (s *IntegrationTestSuite) testQueryGlobalFeesInGenesis() {
	chainAAPIEndpoint := fmt.Sprintf("http://%s", s.valResources[s.chainA.id][0].GetHostPort("1317/tcp"))
	feeInGenesis, err := sdk.ParseDecCoins(initialGlobalFeeAmt + denom)
	s.Require().NoError(err)
	s.Require().Eventually(
		func() bool {
			fees, err := queryGlobalFees(chainAAPIEndpoint)
			s.T().Logf("Global Fees in Genesis: %s", fees.String())
			s.Require().NoError(err)

			return fees.IsEqual(feeInGenesis)
		},
		15*time.Second,
		5*time.Second,
	)
}

func (s *IntegrationTestSuite) testGlobalFees() {
	chainAAPIEndpoint := fmt.Sprintf("http://%s", s.valResources[s.chainA.id][0].GetHostPort("1317/tcp"))

	submitterAddr, _ := s.chainA.validators[0].keyInfo.GetAddress()
	submitter := submitterAddr.String()
	recipientAddress, _ := s.chainA.validators[1].keyInfo.GetAddress()
	recipient := recipientAddress.String()

	var beforeRecipientPhotonBalance sdk.Coin
	s.Require().Eventually(
		func() bool {
			var err error
			beforeRecipientPhotonBalance, err = querySpecificBalance(chainAAPIEndpoint, recipient, photonDenom)
			s.Require().NoError(err)

			return beforeRecipientPhotonBalance.IsValid()
		},
		10*time.Second,
		5*time.Second,
	)
	if beforeRecipientPhotonBalance.Equal(sdk.Coin{}) {
		beforeRecipientPhotonBalance = sdk.NewCoin(photonDenom, sdk.ZeroInt())
	}

	sendAmt := int64(1000)
	token := sdk.NewInt64Coin(photonDenom, sendAmt) // send 1000photon each time
	sucessBankSendCount := 0

	// ---------------------------- test1: globalfee empty --------------------------------------------
	// prepare gov globalfee proposal
	emptyGlobalFee := sdk.DecCoins{}
	proposalCounter++
	s.govProposeNewGlobalfee(emptyGlobalFee, proposalCounter, submitter, standardFees.String())
	paidFeeAmt := math.LegacyMustNewDecFromStr(minGasPrice).Mul(math.LegacyNewDec(gas)).String()

	s.T().Logf("test case: empty global fee, globalfee=%s, min_gas_price=%s", emptyGlobalFee.String(), minGasPrice+denom)
	txBankSends := []txBankSend{
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "1" + denom,
			log:       "Tx fee is zero coin with correct denom: nshr, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "",
			log:       "Tx fee is empty, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "4" + photonDenom,
			log:       "Tx with wrong denom: photon, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmt + denom,
			log:       "Tx fee is higher than min_gas_price, pass",
			expectErr: false,
		},
	}
	sucessBankSendCount += s.execBankSendBatch(s.chainA, 0, txBankSends...)

	// ------------------ test2: globalfee lower than min_gas_price -----------------------------------
	// prepare gov globalfee proposal
	lowGlobalFee := sdk.DecCoins{sdk.NewDecCoinFromDec(denom, sdk.MustNewDecFromStr(lowGlobalFeesAmt))}
	proposalCounter++
	s.govProposeNewGlobalfee(lowGlobalFee, proposalCounter, submitter, standardFees.String())
	paidFeeAmt = math.LegacyMustNewDecFromStr(minGasPrice).Mul(math.LegacyNewDec(gas)).String()

	s.T().Logf("test case: global fee is lower than min_gas_price, globalfee=%s, min_gas_price=%s", lowGlobalFee.String(), minGasPrice+denom)
	txBankSends = []txBankSend{
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmt + denom,
			log:       "Tx fee higher than/equal to min_gas_price and global fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmt + photonDenom,
			log:       "Tx fee has wrong denom, fail",
			expectErr: true,
		},
	}
	sucessBankSendCount += s.execBankSendBatch(s.chainA, 0, txBankSends...)

	// ------------------ test3: globalfee higher than min_gas_price ----------------------------------
	// prepare gov globalfee proposal
	highGlobalFee := sdk.DecCoins{sdk.NewDecCoinFromDec(denom, sdk.MustNewDecFromStr(highGlobalFeeAmt))}
	proposalCounter++
	s.govProposeNewGlobalfee(highGlobalFee, proposalCounter, submitter, paidFeeAmt+denom)

	paidFeeAmt = math.LegacyMustNewDecFromStr(highGlobalFeeAmt).Mul(math.LegacyNewDec(gas)).String()
	paidFeeAmtHigherMinGasLowerGalobalFee := math.LegacyMustNewDecFromStr(minGasPrice).
		Quo(math.LegacyNewDec(2)).String()

	s.T().Logf("test case: global fee is higher than min_gas_price, globalfee=%s, min_gas_price=%s", highGlobalFee.String(), minGasPrice+denom)
	txBankSends = []txBankSend{
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmt + denom,
			log:       "Tx fee is higher than/equal to global fee and min_gas_price, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmtHigherMinGasLowerGalobalFee + denom,
			log:       "Tx fee is higher than/equal to min_gas_price but lower than global fee, fail",
			expectErr: true,
		},
	}
	sucessBankSendCount += s.execBankSendBatch(s.chainA, 0, txBankSends...)

	// ---------------------------------------------------------------------------
	// check the balance is correct after previous txs
	s.Require().Eventually(
		func() bool {
			afterRecipientPhotonBalance, err := querySpecificBalance(chainAAPIEndpoint, recipient, photonDenom)
			s.Require().NoError(err)
			IncrementedPhoton := afterRecipientPhotonBalance.Sub(beforeRecipientPhotonBalance)
			photonSent := sdk.NewInt64Coin(photonDenom, sendAmt*int64(sucessBankSendCount))
			return IncrementedPhoton.IsEqual(photonSent)
		},
		time.Minute,
		5*time.Second,
	)

	// gov proposing to change back to original global fee
	s.T().Logf("Propose to change back to original global fees: %s", initialGlobalFeeAmt+denom)
	oldfees, err := sdk.ParseDecCoins(initialGlobalFeeAmt + denom)
	s.Require().NoError(err)
	proposalCounter++
	s.govProposeNewGlobalfee(oldfees, proposalCounter, submitter, paidFeeAmt+photonDenom)
}
