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
func TestElectoralModule(t *testing.T) {
	cf := network.ShareLedgerTestingConfig()
	cf.NumValidators = 2
	suite.Run(t, NewElectoralIntegrationTestSuite(cf))
}
