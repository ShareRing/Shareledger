package tests

import (
	"testing"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

func TestDocumentIntegrationTest(t *testing.T) {
	networkConfig := network.ShareLedgerTestingConfig()
	networkConfig.NumValidators = 1
	suite.Run(t, NewDocumentIntegrationTestSuite(networkConfig))
}
