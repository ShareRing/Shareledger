//go:build e2e

package document

import (
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/document/types"
	idtypes "github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
)

var (
	firstDoc = types.Document{
		Holder:  "Id1",
		Issuer:  "shareledger1adlezqeqsajf7fmdlluyhs2nv8tturwv9s3paq", // doc issuer account address
		Proof:   "testProof",
		Data:    "testData",
		Version: 0,
	}

	// test update with mismatch doc issuer acc
	secondDoc = types.Document{
		Holder:  "Id7",
		Issuer:  "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak", // doc issuer account address
		Proof:   "testProof2",
		Data:    "testData2",
		Version: 0,
	}

	firstId = idtypes.Id{
		Id: "Id1",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak",
			BackupAddress: "BackupAddress1",
			OwnerAddress:  "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
			ExtraData:     "ExtraData1",
		},
	}

	secondId = idtypes.Id{
		Id: "Id2",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress2",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData2",
		},
	}

	thirdId = idtypes.Id{
		Id: "Id3",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress3",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData3",
		},
	}

	fourthId = idtypes.Id{
		Id: "Id4",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress4",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData4",
		},
	}

	fifthId = idtypes.Id{
		Id: "Id5",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress5",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData5",
		},
	}

	sixthId = idtypes.Id{
		Id: "Id6",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress6",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData6",
		},
	}

	eightId = idtypes.Id{
		Id: "Id8",
		Data: &idtypes.BaseID{
			IssuerAddress: "shareledger1zqhw26j0el2u080ua62u8zrcassxx93h7cddlf",
			BackupAddress: "BackupAddress8",
			OwnerAddress:  "shareledger1mfru9azs5nua2wxcd4sq64g5nt7nn4n85mcr0s",
			ExtraData:     "ExtraData8",
		},
	}
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger document module")

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)
	docGenesis := types.GenesisState{
		Documents: []*types.Document{&firstDoc, &secondDoc},
	}

	idGenesis := idtypes.GenesisState{
		IDs: []*idtypes.Id{&firstId, &secondId, &thirdId, &fourthId, &fifthId, &sixthId, &eightId},
	}

	docGenesisBz, err := s.cfg.Codec.MarshalJSON(&docGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[types.ModuleName] = docGenesisBz

	idGenesisBz, err := s.cfg.Codec.MarshalJSON(&idGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[idtypes.ModuleName] = idGenesisBz

	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
