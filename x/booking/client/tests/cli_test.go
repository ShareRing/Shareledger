package tests

import (
	"sync"
	"testing"

	"github.com/sharering/shareledger/cmd/Shareledgerd/cli"

	"github.com/stretchr/testify/suite"

	"github.com/sharering/shareledger/testutil/network"
)

var runOnce = sync.Once{}

func init() {
	runOnce.Do(func() {
		cli.InitMiddleWare()
	})
}

func TestIntegrationBookingTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 2
	suite.Run(t, NewBookingIntegrationTestSuite(cfg))
}
