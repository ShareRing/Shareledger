package tests

import (
	"testing"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

func TestIDModuleIntegration(t *testing.T) {
	networkConf := network.ShareLedgerTestingConfig()
	networkConf.NumValidators = 2
	suite.Run(t, NewIDIntegrationTestSuite(networkConf))
}
