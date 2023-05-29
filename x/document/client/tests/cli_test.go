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

func TestDocumentIntegrationTest(t *testing.T) {
	networkConfig := network.DefaultConfig()
	networkConfig.NumValidators = 1
	suite.Run(t, NewDocumentIntegrationTestSuite(networkConfig))
}
