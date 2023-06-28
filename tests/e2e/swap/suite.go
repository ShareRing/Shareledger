//go:build e2e

package swap

import (
	"fmt"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sharering/shareledger/testutil/network"
	swapTypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"
	"github.com/stretchr/testify/suite"
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
	s.T().Log("setting up e2e test suite for shareledger swap module")
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
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))

	//swap genesis
	swapGenesis := swapTypes.GenesisState{
		Schemas: []swapTypes.Schema{
			{
				Network: "eth",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
			},
			{
				Network: "schemaWithoutFee",
				Fee:     nil,
			},
			{
				Network: "hero",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
				Schema: "{}",
			},
			{
				Network: "hero1",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
				Schema: "{}",
			},

			{
				Network:          "k_eth",
				Schema:           "{\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Swap\":[{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"tos\",\"type\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\"}]},\"primaryType\":\"Swap\",\"domain\":{\"name\":\"ShareRingSwap\",\"version\":\"2.0\",\"chainId\":\"0x3\",\"verifyingContract\":\"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560\",\"salt\":\"\"}}",
				ContractExponent: 2,
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
			},
		},
		Requests: []swapTypes.Request{
			{
				Id:          uint64(1),
				SrcAddr:     network.Accounts[network.KeyAccount1].String(),
				DestAddr:    "0xx1234x",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameEthereum,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				Status:      swapTypes.SwapStatusPending,
			},
			{
				Id:          uint64(2),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				Status:      swapTypes.SwapStatusPending,
			},
			{
				Id:          uint64(3),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(4),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(5),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(6),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(7),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(8),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0x5a15dfb7a7f3dccf75e62101af3a908a0078b4f5",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: "k_eth",
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusPending,
			},
		},
		Batches: []swapTypes.Batch{
			{
				Id:         1,
				Signature:  "xx",
				RequestIds: []uint64{3},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
			{
				Id:         2, // for test cancel batch
				Signature:  "xx",
				RequestIds: []uint64{4, 5},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
			{
				Id:         3, // for test complete batch
				Signature:  "xx",
				RequestIds: []uint64{6, 7},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
		},
		PastTxEvent: []swapTypes.PastTxEventGenesis{
			{
				SrcAddr:  "src111",
				DestAddr: "dest111",
				TxHash:   "hashxxxxx",
				LogIndex: 1,
			},
			{
				SrcAddr:  "src222",
				DestAddr: "dest22",
				TxHash:   "hashxxxxx",
				LogIndex: 2,
			},
		},
		BatchCount:   4,
		RequestCount: 8,
	}
	swapGenesisBz, err := s.cfg.Codec.MarshalJSON(&swapGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[swapTypes.ModuleName] = swapGenesisBz
	//Bank genesis
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

func (s *E2ETestQuerySuite) SetupSuite() {
	s.T().Log("setting up e2e test suite for shareledger swap module")
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
	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))

	//swap genesis
	swapGenesis := swapTypes.GenesisState{
		Schemas: []swapTypes.Schema{
			{
				Network: "eth",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
			},
			{
				Network: "schemaWithoutFee",
				Fee:     nil,
			},
			{
				Network: "hero",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
				Schema: "{}",
			},
			{
				Network: "hero1",
				Fee: &swapTypes.Fee{
					In:  &coin1,
					Out: &coin2,
				},
				Schema: "{}",
			},
		},
		Requests: []swapTypes.Request{
			{
				Id:          uint64(1),
				SrcAddr:     network.Accounts[network.KeyAccount1].String(),
				DestAddr:    "0xx1234x",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameEthereum,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				Status:      swapTypes.SwapStatusPending,
			},
			{
				Id:          uint64(2),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				Status:      swapTypes.SwapStatusPending,
			},
			{
				Id:          uint64(3),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(4),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(5),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(6),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
			{
				Id:          uint64(7),
				SrcAddr:     "shareledger1wh7w0p2naly7anxdsyes3u6028pd2uc4vnxt2y",
				DestAddr:    "0xx1234",
				SrcNetwork:  swapTypes.NetworkNameShareLedger,
				DestNetwork: swapTypes.NetworkNameBinanceSmartChain,
				Amount:      sdk.NewCoin(denom.Base, sdk.NewInt(1000*denom.ShrExponent)),
				Fee:         coin2,
				BatchId:     1,
				Status:      swapTypes.SwapStatusApproved,
			},
		},
		Batches: []swapTypes.Batch{
			{
				Id:         1,
				Signature:  "xx",
				RequestIds: []uint64{3},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
			{
				Id:         2, // for test cancel batch
				Signature:  "xx",
				RequestIds: []uint64{4, 5},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
			{
				Id:         3, // for test complete batch
				Signature:  "xx",
				RequestIds: []uint64{6, 7},
				Status:     swapTypes.BatchStatusPending,
				Network:    swapTypes.NetworkNameBinanceSmartChain,
			},
		},
		PastTxEvent: []swapTypes.PastTxEventGenesis{
			{
				SrcAddr:  "src111",
				DestAddr: "dest111",
				TxHash:   "hashxxxxx",
				LogIndex: 1,
			},
			{
				SrcAddr:  "src222",
				DestAddr: "dest22",
				TxHash:   "hashxxxxx",
				LogIndex: 2,
			},
		},
		BatchCount:   4,
		RequestCount: 7,
	}
	swapGenesisBz, err := s.cfg.Codec.MarshalJSON(&swapGenesis)
	s.Require().NoError(err)
	s.cfg.GenesisState[swapTypes.ModuleName] = swapGenesisBz
	//Bank genesis
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
