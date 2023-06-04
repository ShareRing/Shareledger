//go:build e2e

package swap

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

func TestE2ETestSuite(t *testing.T) {
	conf := network.DefaultConfig()
	conf.NumValidators = 1
	suite.Run(t, NewE2ETestSuite(conf))
}
