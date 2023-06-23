package booking

import (
	"fmt"
	"path/filepath"

	"github.com/sharering/shareledger/testutil/network"
	assettypes "github.com/sharering/shareledger/x/asset/types"
	"github.com/sharering/shareledger/x/booking/types"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

var (
	booking1 = &types.Booking{
		BookID:      "BookID1",
		Booker:      "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw",
		UUID:        "UUID",
		Duration:    100,
		IsCompleted: true,
	}
	asset1 = &assettypes.Asset{
		Creator: "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4",
		Hash:    []byte{123},
		UUID:    "AssetUUID1",
		Status:  true,
		Rate:    1,
	}
	asset2 = &assettypes.Asset{
		Creator: "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4",
		Hash:    []byte{123},
		UUID:    "AssetUUID2",
		Status:  true,
		Rate:    1,
	}
)

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for booking module")
	// the nodeDir, and moniker hard code at here in cosmos-sdk:
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:398
	// So just reuse it
	rootDir := s.T().TempDir()
	moniker := fmt.Sprintf("node%d", s.cfg.NumValidators-1)
	// TestingGenesis should use the same KeyringDir as validator KeyringDir
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:400
	nodeDir := filepath.Join(rootDir, moniker, "simcli")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg, nodeDir, moniker)
	s.Require().NotNil(kr)
	bookingGenesis := &types.GenesisState{
		Bookings: []*types.Booking{booking1},
	}
	assetGenesis := &assettypes.GenesisState{
		Assets: []*assettypes.Asset{asset1, asset2},
	}
	s.cfg.GenesisState[types.ModuleName] = s.cfg.Codec.MustMarshalJSON(bookingGenesis)
	s.cfg.GenesisState[assettypes.ModuleName] = s.cfg.Codec.MustMarshalJSON(assetGenesis)

	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.NoError(s.network.WaitForNextBlock())
}
