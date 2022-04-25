package tests

import (
	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}
func TestSwappingModuleIntegration(t *testing.T) {
	networkConf := network.DefaultConfig()
	networkConf.NumValidators = 1
	suite.Run(t, NewSwapIntegrationTestSuite(networkConf))
}
