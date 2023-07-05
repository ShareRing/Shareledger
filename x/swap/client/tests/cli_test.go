package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/testutil/network"
)

var (
	runOnce = sync.Once{}
	NetConf = network.DefaultConfig()
)

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}
func TestSwappingModuleIntegration(t *testing.T) {
	NetConf.NumValidators = 2
	suite.Run(t, NewSwapIntegrationTestSuite(NetConf))
}
