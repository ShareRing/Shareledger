//go:build e2e

package id

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/gogoproto/proto"
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

func (s *E2ETestSuite) TestGetIDs() {
	testCases := []struct {
		name      string
		args      []string
		expectErr bool
		respType  proto.Message
		expected  proto.Message
	}{{
		name: "valid query id by id",
		args: []string{
			"Id1",
			fmt.Sprintf("--%s=json", flags.FlagOutput),
		},
		expectErr: false,
		respType:  &types.QueryIdByIdResponse{},
		expected: &types.QueryIdByIdResponse{
			Id: &id1,
		},
	}, {
		name: "query id by id not pass Id",
		args: []string{
			fmt.Sprintf("--%s=json", flags.FlagOutput),
		},
		expectErr: true,
		respType:  nil,
		expected:  nil,
	}}
	val := s.network.Validators[0]
	for _, tc := range testCases {
		cmd := cli.CmdIdById()
		out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.args)
		if tc.expectErr {
			s.Require().Error(err)
		} else {
			s.Require().NoError(err)
			s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), tc.respType))
			s.Require().Equal(tc.expected.String(), tc.respType.String())
		}
	}
}
