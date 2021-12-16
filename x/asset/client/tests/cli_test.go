package tests

import (
	"github.com/ShareRing/Shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.ShareLedgerTestingConfig()
	cfg.NumValidators = 2
	suite.Run(t, NewAssetIntegrationTestSuite(cfg))
}
