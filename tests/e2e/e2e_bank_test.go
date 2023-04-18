package e2e

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *IntegrationTestSuite) testBankTokenTransfer() {
	s.Run("send_nshr_between_accounts", func() {
		senderAddress, err := s.chainA.validators[0].keyInfo.GetAddress()
		s.Require().NoError(err)
		sender := senderAddress.String()

		recipientAddress, err := s.chainA.validators[1].keyInfo.GetAddress()
		s.Require().NoError(err)
		recipient := recipientAddress.String()

		chainAAPIEndpoint := fmt.Sprintf("http://%s", s.valResources[s.chainA.id][0].GetHostPort("1317/tcp"))

		var (
			beforeSenderBalance    sdk.Coin
			beforeRecipientBalance sdk.Coin
		)

		s.Require().Eventually(
			func() bool {
				beforeSenderBalance, err = querySpecificBalance(chainAAPIEndpoint, sender, denom)
				s.Require().NoError(err)

				beforeRecipientBalance, err = querySpecificBalance(chainAAPIEndpoint, recipient, denom)
				s.Require().NoError(err)
				return beforeSenderBalance.IsValid() && beforeRecipientBalance.IsValid()
			}, 10*time.Second, 5*time.Second,
		)

		s.execBankSend(s.chainA, 0, sender, recipient, tokenAmount.String(), standardFees.String(), false)

		s.Require().Eventually(func() bool {
			afterSenderBalance, err := querySpecificBalance(chainAAPIEndpoint, sender, denom)
			s.Require().NoError(err)

			afterReceipientBalance, err := querySpecificBalance(chainAAPIEndpoint, recipient, denom)
			s.Require().NoError(err)

			decremented := beforeSenderBalance.Sub(tokenAmount).Sub(standardFees).IsEqual(afterSenderBalance)
			incremented := beforeRecipientBalance.Add(tokenAmount).IsEqual(afterReceipientBalance)

			return decremented && incremented
		}, time.Minute, 5*time.Second)
	})
}
