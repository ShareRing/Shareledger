package electoral

import (
	"encoding/json"

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
	accSwapManager    types.AccState
	accIDSigner       types.AccState
	accVoter          types.AccState
	accDocIssuer      types.AccState
	accKeyShrpLoaders types.AccState
	accKeyRelayer     types.AccState
	accApprover       types.AccState
	accOp             types.AccState

	address         = "shareledger19ac3d6cwqwpzvaxr4xv9kfduwtyswad88fjgw4"
	idSignerAccKey  = "idsignershareledger1z8q5ml2nemt63zd50u3frtvcxfuuttkllvdlsy"
	idsignerAddress = "shareledger1z8q5ml2nemt63zd50u3frtvcxfuuttkllvdlsy"

	docIssuerAccKey  = "docIssuershareledger14gytjg3zdpqmakreduy26hdpmevpsd8dycvmte"
	docIssuerAddress = "shareledger14gytjg3zdpqmakreduy26hdpmevpsd8dycvmte"

	accDocIssuer1 = types.AccState{
		Address: docIssuerAddress,
		Key:     docIssuerAccKey,
		Status:  "active",
	}
	accIDSigner1 = types.AccState{
		Address: idsignerAddress,
		Key:     idSignerAccKey,
		Status:  "active",
	}
)

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger electoral module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
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

	var genesis types.GenesisState
	err = json.Unmarshal(s.cfg.GenesisState[types.ModuleName], &genesis)
	s.NoError(err)

	genesis.AccStateList = append(genesis.AccStateList, accApprover, accDocIssuer, accIDSigner, accKeyRelayer, accKeyShrpLoaders, accOp, accSwapManager, accVoter)
	genesisJSON, err := json.Marshal(genesis)
	s.NoError(err)

	s.cfg.GenesisState[types.ModuleName] = genesisJSON
	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr

	s.NoError(s.network.WaitForNextBlock())
}
