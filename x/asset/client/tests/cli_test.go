package tests

import (
	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.ShareLedgerTestingConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewAssetIntegrationTestSuite(&cfg))
}
