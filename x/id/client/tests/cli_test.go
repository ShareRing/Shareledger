package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/testutil/network"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}
func TestIDModuleIntegration(t *testing.T) {
	networkConf := network.DefaultConfig()
	networkConf.NumValidators = 1
	suite.Run(t, NewIDIntegrationTestSuite(networkConf))
}
