package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestGentlemintIntegrationTestSuite(t *testing.T) {
	networkConfig := network.ShareLedgerTestingConfig()
	networkConfig.NumValidators = 1
	suite.Run(t, NewGentlemintIntegrationTestSuite(networkConfig))
}
