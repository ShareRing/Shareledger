//go:build e2e

package id

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/id/client/cli"
	"github.com/sharering/shareledger/x/id/types"
	"github.com/stretchr/testify/suite"
)

var id1 = types.Id{
	Id: "Id1",
	Data: &types.BaseID{
		IssuerAddress: "shareledger18g8x9censnr3k2y7x6vwntlhvz254ym4qflcak",
		BackupAddress: "BackupAddress",
		OwnerAddress:  "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
		ExtraData:     "ExtraData",
	},
}

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	genesisState := s.cfg.GenesisState
	var idGenesis types.GenesisState
	s.Require().NoError(s.cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &idGenesis))
	idGenesis.IDs = []*types.Id{&id1}
	idGenesisBz, err := s.cfg.Codec.MarshalJSON(&idGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = idGenesisBz
	s.cfg.GenesisState = genesisState
	s.network = network.New(s.T(), s.cfg)
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TestGetByID() {
	testCases := tests.TestCases{
		{
			Name: "valid query id by id",
			Args: []string{
				"Id1",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			ExpectErr: false,
			RespType:  &types.QueryIdByIdResponse{},
			Expected: &types.QueryIdByIdResponse{
				Id: &id1,
			},
		}, {
			Name:      "query id by id not pass Id",
			Args:      []string{},
			ExpectErr: true,
			RespType:  nil,
			Expected:  nil,
		},
	}
	tests.RunTestCases(s.Require(), testCases, cli.CmdIdById(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestGetByAddress() {
	testCases := tests.TestCases{{
		Name: "valid query id by address",
		Args: []string{
			"shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr",
			fmt.Sprintf("--%s=json", flags.FlagOutput),
		},
		ExpectErr: false,
		RespType:  &types.QueryIdByIdResponse{},
		Expected: &types.QueryIdByIdResponse{
			Id: &id1,
		},
	}, {
		Name:      "query id by address not pass address",
		Args:      []string{},
		ExpectErr: true,
		RespType:  nil,
		Expected:  nil,
	}}
	tests.RunTestCases(s.Require(), testCases, cli.CmdIdByAddress(), s.network.Validators[0])
}
