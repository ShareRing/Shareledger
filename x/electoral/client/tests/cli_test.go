package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

func TestElectoralModule(t *testing.T) {
	cf := network.ShareLedgerTestingConfig()
	cf.NumValidators = 2
	suite.Run(t, NewElectoralIntegrationTestSuite(cf))
}
