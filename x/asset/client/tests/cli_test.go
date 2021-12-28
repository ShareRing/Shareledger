package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.ShareLedgerTestingConfig()
	cfg.NumValidators = 2
	suite.Run(t, NewAssetIntegrationTestSuite(cfg))
}
