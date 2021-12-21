package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIDModuleIntegration(t *testing.T) {
	networkConf := network.ShareLedgerTestingConfig()
	networkConf.NumValidators = 2
	suite.Run(t, NewIDIntegrationTestSuite(networkConf))
}
