package swap

import (
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

	coin1 := sdk.NewCoin(denom.Base, sdk.NewInt(20*denom.ShrExponent))
	coin2 := sdk.NewCoin(denom.Base, sdk.NewInt(30*denom.ShrExponent))

	kr, _ := network.SetTestingGenesis(s.T(), &s.cfg)

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
				DestAddr:    "0xx1234",
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
	s.network = network.New(s.T(), s.cfg)
	s.network.Validators[0].ClientCtx.Keyring = kr
	s.Require().NoError(s.network.WaitForNextBlock())
}
