package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestElectoralModule(t *testing.T) {
	cf := network.ShareLedgerTestingConfig()
	cf.NumValidators = 2
	suite.Run(t, NewElectoralIntegrationTestSuite(cf))
}
