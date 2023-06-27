package swap

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sharering/shareledger/pkg/swap"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/swap/client/cli"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
)

type E2ETestApprove struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestApprove(cfg network.Config) *E2ETestApprove {
	return &E2ETestApprove{cfg: cfg}
}

func (s *E2ETestApprove) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger swap module")

	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))

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

	// swap genesis
	swapGenesis := swapTypes.GenesisState{
		Schemas: []swapTypes.Schema{
			{
				Network:          "eth",
				Schema:           `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}`,
				ContractExponent: 2,
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
			},
		},
		Requests: []swapTypes.Request{
			{
				Id:       3000,
				DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
				Amount: sdk.Coin{
					Denom:  denom.Base,
					Amount: sdk.NewInt(3455000000000), // 3455 shr
				},
				Status:     "pending",
				SrcNetwork: "shareledger",
				BatchId:    0,
			}, {
				Id:       3001,
				DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
				Amount: sdk.Coin{
					Denom:  denom.Base,
					Amount: sdk.NewInt(6733000000000), // 6733 shr
				},
				Status:     "pending",
				SrcNetwork: "shareledger",
				BatchId:    0,
			},
		},
		Batches:      []swapTypes.Batch{},
		PastTxEvent:  []swapTypes.PastTxEventGenesis{},
		BatchCount:   4,
		RequestCount: 8,
	}
	swapGenesisBz, err := s.cfg.Codec.MarshalJSON(&swapGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[swapTypes.ModuleName] = swapGenesisBz
	// Bank genesis
	bankGen := bankTypes.GenesisState{}
	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[bankTypes.ModuleName], &bankGen)
	bankGen.Balances = append(bankGen.Balances, bankTypes.Balance{
		Address: authtypes.NewModuleAddress(swapTypes.ModuleName).String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(denom.Base, sdk.NewInt(10000000*denom.ShrExponent))),
	})

	bankGenesisBz, err := s.cfg.Codec.MarshalJSON(&bankGen)
	s.NoError(err)
	s.cfg.GenesisState[bankTypes.ModuleName] = bankGenesisBz
	s.network = network.New(s.T(), rootDir, s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestApprove) TestApproveRequest() {
	cliCtx := s.network.Validators[0].ClientCtx

	// Approve
	approveOut, err := clitestutil.ExecTestCLICmd(cliCtx, cli.CmdApprove(), append([]string{"3000,3001", network.KeyAccountTestSign, "eth", network.MakeByAccount(network.KeyApproverRelayer)}, tests.DefaultTxFlag()...))
	s.Require().NoError(err)
	var resp types.TxResponse
	s.NoError(cliCtx.Codec.UnmarshalJSON(approveOut.Bytes(), &resp))

	txRes, err := tests.QueryTxWithRetry(cliCtx, resp.TxHash, tests.DefaultNumRetry)
	s.Equal(uint32(0), txRes.Code)

	log := network.ParseRawLogGetEvent(s.T(), txRes.RawLog)[0]
	attr := log.Events.GetEventByType(s.T(), swapTypes.EventTypeApproveRequests)
	batchID := attr.Get(s.T(), swapTypes.EventAttrBatchId).Value

	bQuery := fmt.Sprintf("--ids=%s", batchID)
	out, err := clitestutil.ExecTestCLICmd(cliCtx, cli.CmdBatches(), []string{bQuery, "--output=json"})
	s.NoError(err)

	batchRes := swapTypes.QueryBatchesResponse{}

	batchIDNum, err := strconv.ParseUint(batchID, 10, 64)
	s.NoError(err)
	err = cliCtx.Codec.UnmarshalJSON(out.Bytes(), &batchRes)
	s.NoErrorf(err, "fail to get back batch %s", out.String())
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))
	signDetail := swap.NewSignDetail([]swapTypes.Request{
		{
			Id:       3000,
			DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
			Amount: sdk.Coin{
				Denom:  denom.Base,
				Amount: sdk.NewInt(3455000000000), // 3455 shr
			},
			Status:     "pending",
			SrcNetwork: "shareledger",
			BatchId:    batchIDNum,
		}, {
			Id:       3001,
			DestAddr: "0x97b98d335c28f9ad9c123e344a78f00c84146431",
			Amount: sdk.Coin{
				Denom:  denom.Base,
				Amount: sdk.NewInt(6733000000000), // 6733 shr
			},
			Status:     "pending",
			SrcNetwork: "shareledger",
			BatchId:    batchIDNum,
		},
	}, swapTypes.Schema{
		Network:          "eth",
		Schema:           `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}`,
		ContractExponent: 2,
		Fee: &swapTypes.Fee{
			In:  &coin1,
			Out: &coin2,
		},
	})
	digest, err := signDetail.Digest()
	s.NoError(err, "get digest fail")
	ks := keyring.NewKeyRingETH(cliCtx.Keyring)
	_, npk, err := ks.Sign(network.KeyAccountTestSign, digest.Bytes())
	s.NoError(err, "sign the swap request fail")
	sig, _ := hexutil.Decode(batchRes.Batches[0].Signature)

	s.Equalf(true, npk.VerifySignature(digest.Bytes(), sig), "verify sign fail")
	s.Equalf("0xd9a5705095d8c83fc051fde2dda2e47fb81d16ee23f11f9322c0656e6020ee9001f73fd951ad0d1d7d36a59900d2bd481c47477d01fbd1a2a6da5c7f6d78129d1c", batchRes.Batches[0].GetSignature(), "eip sign not same")
	s.Equalf("0xb63b8aa6f75b29271051d9069070d0555f4e6cdaf35e72d69ffcb366a4d47a08", digest.String(), "digest not equal")
}
