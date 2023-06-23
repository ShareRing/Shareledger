package electoral

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}
type E2ETestSuiteTx struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

func NewE2ETestSuiteTx(cfg network.Config) *E2ETestSuiteTx {
	cfg.NumValidators = 1
	return &E2ETestSuiteTx{cfg: cfg}
}

var (
	accSwapManager    types.AccState
	accIDSigner       types.AccState
	accVoter          types.AccState
	accDocIssuer      types.AccState
	accKeyShrpLoaders types.AccState
	accKeyRelayer     types.AccState
	accApprover       types.AccState
	accOp             types.AccState
	accDocIssuer1     types.AccState
	accIDSigner1      types.AccState

	address = "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4"
)

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger electoral module")
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
	addr, err := sdk.AccAddressFromBech32(address)
	s.NoError(err)

	accOp = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyAccOp)),
		Status:  "active",
	}
	accIDSigner = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)),
		Status:  "active",
	}
	accVoter = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyVoter)),
		Status:  "active",
	}
	accDocIssuer = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyDocIssuer)),
		Status:  "active",
	}
	accKeyShrpLoaders = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyShrpLoaders)),
		Status:  "active",
	}
	accKeyRelayer = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyRelayer)),
		Status:  "active",
	}
	accApprover = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyApprover)),
		Status:  "active",
	}
	accSwapManager = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeySwapManager)),
		Status:  "active",
	}

	genesis := types.GenesisState{}

	genesis.AccStateList = append(genesis.AccStateList, accApprover, accDocIssuer, accIDSigner, accKeyRelayer, accKeyShrpLoaders, accOp, accSwapManager, accVoter)
	genesisJSON, err := json.Marshal(genesis)
	s.NoError(err)

	s.cfg.GenesisState[types.ModuleName] = genesisJSON
	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr

	s.NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuiteTx) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger electoral tx module")
	// the nodeDir, and moniker hard code at here in cosmos-sdk:
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:398
	// So just replicate it
	rootDir := s.T().TempDir()
	moniker := fmt.Sprintf("node%d", s.cfg.NumValidators-1)
	// TestingGenesis should use the same KeyringDir as validator KeyringDir
	// github.com/sharering/cosmos-sdk@v0.47.2-shareledger/testutil/network/network.go:400
	nodeDir := filepath.Join(rootDir, moniker, "simcli")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg, nodeDir, moniker)
	s.Require().NotNil(kr)
	addr, err := sdk.AccAddressFromBech32(address)
	s.NoError(err)

	docIssuerAddr, err := sdk.AccAddressFromBech32(network.Accounts[network.KeyDocIssuer].String())
	s.NoError(err)

	idSignerAddr, err := sdk.AccAddressFromBech32(network.Accounts[network.KeyIDSigner].String())
	s.NoError(err)

	accDocIssuer1 = types.AccState{
		Address: network.Accounts[network.KeyDocIssuer].String(),
		Key:     string(types.GenAccStateIndexKey(docIssuerAddr, types.AccStateKeyDocIssuer)),
		Status:  "active",
	}
	accIDSigner1 = types.AccState{
		Address: network.Accounts[network.KeyIDSigner].String(),
		Key:     string(types.GenAccStateIndexKey(idSignerAddr, types.AccStateKeyIdSigner)),
		Status:  "active",
	}
	accOp = types.AccState{
		Address: network.Accounts[network.KeyOperator].String(),
		Key:     string(types.GenAccStateIndexKey(network.Accounts[network.KeyOperator].(sdk.AccAddress), types.AccStateKeyAccOp)),
		Status:  "active",
	}
	accIDSigner = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyIdSigner)),
		Status:  "active",
	}
	accVoter = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyVoter)),
		Status:  "active",
	}
	accDocIssuer = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyDocIssuer)),
		Status:  "active",
	}
	accKeyShrpLoaders = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyShrpLoaders)),
		Status:  "active",
	}
	accKeyRelayer = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyRelayer)),
		Status:  "active",
	}
	accApprover = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeyApprover)),
		Status:  "active",
	}
	accSwapManager = types.AccState{
		Address: address,
		Key:     string(types.GenAccStateIndexKey(addr, types.AccStateKeySwapManager)),
		Status:  "active",
	}

	var genesis types.GenesisState
	err = json.Unmarshal(s.cfg.GenesisState[types.ModuleName], &genesis)
	s.NoError(err)

	genesis.AccStateList = append(genesis.AccStateList, accDocIssuer1, accIDSigner1, accApprover, accDocIssuer, accIDSigner, accKeyRelayer, accKeyShrpLoaders, accOp, accSwapManager, accVoter)
	genesisJSON, err := json.Marshal(genesis)
	s.NoError(err)

	s.cfg.GenesisState[types.ModuleName] = genesisJSON
	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr

	s.NoError(s.network.WaitForNextBlock())
}
