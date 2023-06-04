package electoral

import (
	"fmt"

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

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	cfg.NumValidators = 1
	return &E2ETestSuite{cfg: cfg}
}

var (
	address           = "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4"
	accSwapManager    types.AccState
	accIDSigner       types.AccState
	accVoter          types.AccState
	accDocIssuer      types.AccState
	accKeyShrpLoaders types.AccState
	accKeyRelayer     types.AccState
	accApprover       types.AccState
)

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger electoral module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	addr, err := sdk.AccAddressFromBech32(address)
	s.NoError(err)
	accOp := types.AccState{
		Address: address,
		Key:     fmt.Sprintf("%s/%s", types.AccStateKeyAccOp, addr.String()),
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
	electoralGenesis := &types.GenesisState{
		AccStateList: []types.AccState{
			accOp,
			accIDSigner,
			accVoter,
			accDocIssuer,
			accKeyShrpLoaders,
			accKeyRelayer,
			accApprover,
			accSwapManager,
		},
		Authority: &types.Authority{
			Address: address,
		},
		Treasurer: &types.Treasurer{
			Address: address,
		},
	}

	// Marshal the electoralGenesis to JSON
	genesisJSON := s.cfg.Codec.MustMarshalJSON(electoralGenesis)

	fmt.Printf("Genesis JSON - 143: %s\n", string(genesisJSON)) // Log the genesis JSON data

	s.cfg.GenesisState[types.ModuleName] = genesisJSON

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr

	s.NoError(s.network.WaitForNextBlock())
}
