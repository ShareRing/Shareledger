package tests

import (
	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"sync"
	"testing"

	"github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}
func TestIDModuleIntegration(t *testing.T) {
	networkConf := network.ShareLedgerTestingConfig()
	networkConf.NumValidators = 1
	suite.Run(t, NewIDIntegrationTestSuite(networkConf))
}
