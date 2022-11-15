package tests

import (
	"sync"
	"testing"

	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}
func TestElectoralModule(t *testing.T) {
	cf := network.DefaultConfig()
	cf.NumValidators = 2
	suite.Run(t, NewElectoralIntegrationTestSuite(cf))
}
