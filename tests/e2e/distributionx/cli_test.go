//go:build e2e

package distributionx

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

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, NewE2ETestSuite(network.DefaultConfig()))
}
