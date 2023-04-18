package e2e

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// globalfee in genesis is set to be "0.00001nshr"
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

/*
global fee e2e tests:
initial setup: initial globalfee = 10000nshr, min_gas_price = 100000nshr
(This initial value setup is to pass other e2e tests)

test1: gov proposal globalfee = [], min_gas_price=100nshr, query globalfee still get empty
- tx with fee denom photon, fail
- tx with zero fee denom photon, fail
- tx with fee denom nshr, pass
- tx with fee empty, fail

test2: gov propose globalfee =  0.000001nshr(lower than min_gas_price)
- tx with fee higher than 0.000001nshr but lower than 0.00001nshr, fail
- tx with fee higher than/equal to 0.00001nshr, pass
- tx with fee photon fail

test3: gov propose globalfee = 0.0001nshr (higher than min_gas_price)
- tx with fee equal to 0.0001nshr, pass
- tx with fee equal to 0.00001nshr, fail

test4: gov propose globalfee =  0.000001nshr (lower than min_gas_price), 0photon
- tx with fee 0.0000001photon, fail
- tx with fee 0.000001photon, pass
- tx with empty fee, pass
- tx with fee photon pass
- tx with fee 0photon, 0.000005nshr fail
- tx with fee 0photon, 0.00001nshr pass
test5: check balance correct: all the successful bank sent tokens are received
test6: gov propose change back to initial globalfee = 0.00001photon, This is for not influence other e2e tests.
*/
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
		1000*time.Second,
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
		// {
		// 	from:      submitter,
		// 	to:        recipient,
		// 	amt:       token.String(),
		// 	fees:      "0" + denom,
		// 	log:       "Tx fee is zero coin with correct denom: nshr, fail",
		// 	expectErr: true,
		// },
		// {
		// 	from:      submitter,
		// 	to:        recipient,
		// 	amt:       token.String(),
		// 	fees:      "",
		// 	log:       "Tx fee is empty, fail",
		// 	expectErr: true,
		// },
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "4" + photonDenom,
			log:       "Tx with wrong denom: photon, fail",
			expectErr: true,
		},
		// {
		// 	from:      submitter,
		// 	to:        recipient,
		// 	amt:       token.String(),
		// 	fees:      "0" + photonDenom,
		// 	log:       "Tx fee is zero coins of wrong denom: photon, fail",
		// 	expectErr: true,
		// },
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
	paidFeeAmtLowMinGasHighGlobalFee := math.LegacyMustNewDecFromStr(lowGlobalFeesAmt).
		Mul(math.LegacyNewDec(2)).
		Mul(math.LegacyNewDec(gas)).
		String()
	paidFeeAmtLowGlobalFee := math.LegacyMustNewDecFromStr(lowGlobalFeesAmt).Quo(math.LegacyNewDec(2)).String()

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
			fees:      paidFeeAmtLowGlobalFee + denom,
			log:       "Tx fee lower than/equal to min_gas_price and global fee, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmtLowMinGasHighGlobalFee + denom,
			log:       "Tx fee lower than/equal global fee and lower than min_gas_price, fail",
			expectErr: true,
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

	// ---------------------------- test4: global fee with two denoms -----------------------------------
	// prepare gov globalfee proposal
	mixGlobalFee := sdk.DecCoins{
		sdk.NewDecCoinFromDec(photonDenom, sdk.NewDec(0)),
		sdk.NewDecCoinFromDec(denom, sdk.MustNewDecFromStr(lowGlobalFeesAmt)),
	}.Sort()
	proposalCounter++
	s.govProposeNewGlobalfee(mixGlobalFee, proposalCounter, submitter, paidFeeAmt+denom)

	// equal to min_gas_price
	paidFeeAmt = math.LegacyMustNewDecFromStr(minGasPrice).Mul(math.LegacyNewDec(gas)).String()
	paidFeeAmtLow := math.LegacyMustNewDecFromStr(lowGlobalFeesAmt).
		Quo(math.LegacyNewDec(2)).
		Mul(math.LegacyNewDec(gas)).
		String()

	s.T().Logf("test case: global fees contain multiple denoms: one zero coin, one non-zero coin, globalfee=%s, min_gas_price=%s", mixGlobalFee.String(), minGasPrice+denom)
	txBankSends = []txBankSend{
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmt + denom,
			log:       "Tx with fee higher than/equal to one of denom's amount the global fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      paidFeeAmtLow + denom,
			log:       "Tx with fee lower than one of denom's amount the global fee, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "",
			log:       "Tx with fee empty fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "0" + photonDenom,
			log:       "Tx with zero coin in the denom of zero coin of global fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "0" + photonDenom,
			log:       "Tx with zero coin in the denom of zero coin of global fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "2" + photonDenom,
			log:       "Tx with non-zero coin in the denom of zero coin of global fee, pass",
			expectErr: false,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "0" + photonDenom + "," + paidFeeAmtLow + denom,
			log:       "Tx with multiple fee coins, zero coin and low fee, fail",
			expectErr: true,
		},
		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "0" + photonDenom + "," + paidFeeAmt + denom,
			log:       "Tx with multiple fee coins, zero coin and high fee, pass",
			expectErr: false,
		},

		{
			from:      submitter,
			to:        recipient,
			amt:       token.String(),
			fees:      "2" + photonDenom + "," + paidFeeAmt + denom,
			log:       "Tx with multiple fee coins, all higher than global fee and min_gas_price",
			expectErr: false,
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
