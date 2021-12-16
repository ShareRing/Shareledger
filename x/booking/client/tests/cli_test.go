package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ShareRing/Shareledger/testutil/network"
)

func TestIntegrationBookingTestSuite(t *testing.T) {
	cfg := network.ShareLedgerTestingConfig()
	cfg.NumValidators = 2
	suite.Run(t, NewBookingIntegrationTestSuite(cfg))
}
