package tests

import (
	"os"
	"strings"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distributionxType "github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/swap/client/tests"
	swapmoduletypes "github.com/sharering/shareledger/x/swap/types"
	"github.com/sharering/shareledger/x/utils/denom"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	netutilts "github.com/sharering/shareledger/testutil/network"
	"github.com/stretchr/testify/suite"
)

type (
	TinyAddr string
)

func (t TinyAddr) String() string {
	return string(t)
}

type DistributionXIntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	dir     string
	network *network.Network
}

func NewDistributionXIntegrationTestSuite(cf *network.Config) *DistributionXIntegrationTestSuite {
	return &DistributionXIntegrationTestSuite{
		cfg: *cf,
	}
}

func (s *DistributionXIntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite for distributionx module")

	kb, dir := netutilts.GetTestingGenesis(s.T(), &s.cfg)
	s.dir = dir
	delete(s.cfg.GenesisState, "capability")

	disXGenBz := s.cfg.GenesisState[distributionxType.ModuleName]

	disXGen := distributionxType.GenesisState{}
	err := s.cfg.Codec.UnmarshalJSON(disXGenBz, &disXGen)
	s.Require().NoError(err, "setup fail")

	disXGen.BuilderListList = []distributionxType.BuilderList{
		{Id: 1, ContractAddress: "shareledger1suhgf5svhu4usrurvxzlgn54ksxmn8gljarjtxqnapv8kjnp4nrsy56wda"},
	}

	swapGen := swapmoduletypes.GenesisState{}

	s.cfg.Codec.MustUnmarshalJSON(s.cfg.GenesisState[swapmoduletypes.ModuleName], &swapGen)
	feeInCoin := sdk.NewCoin(denom.Base, sdk.NewInt(10000000000))
	feeOutCoin := sdk.NewCoin(denom.Base, sdk.NewInt(20000000000))
	swapGen.Schemas = []swapmoduletypes.Schema{
		{
			Network:          "eth",
			ContractExponent: 2,
			Fee: &swapmoduletypes.Fee{
				In:  &feeInCoin,
				Out: &feeOutCoin,
			},
			Schema: `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"},{"name":"verifyingContract","type":"address"}],"Swap":[{"name":"ids","type":"uint256[]"},{"name":"tos","type":"address[]"},{"name":"amounts","type":"uint256[]"}]},"primaryType":"Swap","domain":{"name":"ShareRingSwap","version":"2.0","chainId":"0x3","verifyingContract":"0x3AE875a6e8E8EB6fa4a0748156CE6b9030E4a560","salt":""}}`,
		},
	}

	disXGen.BuilderListCount = 1
	bz := s.cfg.Codec.MustMarshalJSON(&disXGen)
	s.cfg.GenesisState[distributionxType.ModuleName] = bz

	bz2 := s.cfg.Codec.MustMarshalJSON(&swapGen)
	s.cfg.GenesisState[swapmoduletypes.ModuleName] = bz2

	s.network, _ = network.New(s.T(), s.dir, s.cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	// override the keyring
	s.network.Validators[0].ClientCtx.Keyring = kb

	s.T().Log("setting up integration test suite successfully")
}

func (s *DistributionXIntegrationTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(s.dir), "cleanup test case fails")
	s.network.Cleanup()
	s.T().Log("tearing down integration test suite")
}

func (s *DistributionXIntegrationTestSuite) TestDistributionXWasmTransaction() {
	// Execute the contract with fee
	devPoolRewardBefore, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr")
	if err != nil && !strings.Contains(err.Error(), "not found") {
		s.T().Fail()
	}
	outStandingRewardBefore, err := ExCmdQueryOutStandingReward(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].ValAddress.String(),
	)
	s.NoError(err, "Get validator reward error")

	accByte, err := clitestutil.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		netutilts.Accounts[netutilts.KeyAccount2])
	s.NoError(err)
	makeTransactionAccBalance := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	out, err := ExCmdExecuteContract(
		s.network.Validators[0].ClientCtx,
		"shareledger14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9smpkm95",
		`{"increment":{}}`,
		netutilts.JSONFlag,
		netutilts.SkipConfirmation,
		netutilts.MakeByAccount(netutilts.KeyAccount2),
		netutilts.SyncBroadcast,
		netutilts.SHRFee(50),
	)
	s.Require().NoError(err, "execute the contract fail")
	res := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "broadcast transaction fail %v", res.String())
	_ = s.network.WaitForNextBlock()
	_ = s.network.WaitForNextBlock()
	_ = s.network.WaitForNextBlock()

	_, err = ExCmdListReward(s.network.Validators[0].ClientCtx)

	contractOwnerReward, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw")
	s.NoError(err)

	devPoolReward, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr")
	s.NoError(err)
	// list, err := ExCmdListReward(s.network.Validators[0].ClientCtx)
	// s.T().Log(list)
	accByte, err = clitestutil.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		netutilts.Accounts[netutilts.KeyAccount2])
	s.NoError(err)
	makeTransactionAccBalanceAfterEx := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	// s.T().Log("delegator address", s.network.Validators[0].Address.String())
	s.T().Log("validator address", s.network.Validators[0].ValAddress.String())

	accByte, err = clitestutil.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].Address)
	s.NoError(err)

	outStandingRewardAfter, err := ExCmdQueryOutStandingReward(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].ValAddress.String(),
	)

	// Assertion time
	s.Require().Equalf(sdk.NewInt(6250000000).String(), contractOwnerReward.Reward.GetAmount().AmountOf(denom.Base).String(), "the contract owner reward must be increment")

	s.Require().Equalf(devPoolRewardBefore.Reward.Amount.AmountOf(denom.Base).Add(sdk.NewInt(12500000000)).String(), devPoolReward.Reward.GetAmount().AmountOf(denom.Base).String(), "Devpool must take 25%")
	s.Require().Equalf(makeTransactionAccBalance.Balances.AmountOf(denom.Base).Sub(sdk.NewInt(50*denom.ShrExponent)), makeTransactionAccBalanceAfterEx.Balances.AmountOf(denom.Base), "the transaction execute maker must be reduce by the fee that input")
	s.Require().Equalf(outStandingRewardBefore.Rewards.AmountOf(denom.Base).Add(sdk.NewDec(25*denom.ShrExponent)).String(), outStandingRewardAfter.Rewards.AmountOf(denom.Base).String(), "the validator must take 50% transaction fee from 50shr fee")
}

func (s *DistributionXIntegrationTestSuite) TestDistributionXNormalTransaction() {
	outStandingRewardBefore, err := ExCmdQueryOutStandingReward(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].ValAddress.String(),
	)
	s.NoError(err, "Get validator reward error")
	devPoolRewardBefore, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr")
	if err != nil && !strings.Contains(err.Error(), "not found") {
		s.T().Fail()
	}

	out, err := tests.CmdOut(s.network.Validators[0].ClientCtx,
		"0x7b9039bd633411b48a5a5c4262b5b1a16546d209",
		"eth",
		"100shr",
		netutilts.SHRFee(50),
		netutilts.SkipConfirmation,
		netutilts.SyncBroadcast,
		netutilts.MakeByAccount(netutilts.KeyAccount3))
	if err != nil {
		s.Fail("fail when init the swap out request", err)
	}
	txRes := netutilts.ParseStdOut(s.T(), out.Bytes())
	if txRes.Code != netutilts.ShareLedgerSuccessCode {
		s.Failf("fail when init the swap out request %s", txRes.String())
	}
	_ = s.network.WaitForNextBlock()
	_ = s.network.WaitForNextBlock()
	_ = s.network.WaitForNextBlock()
	devPoolRewardAfter, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr")
	outStandingRewardAfter, err := ExCmdQueryOutStandingReward(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].ValAddress.String(),
	)
	s.NoError(err, "Get validator reward error")

	s.Require().Equalf(
		devPoolRewardBefore.Reward.Amount.
			AmountOf(denom.Base).
			Add(sdk.NewInt(25*denom.ShrExponent)).String(),
		devPoolRewardAfter.Reward.Amount.AmountOf(denom.Base).String(),
		"dev pool account must take 50% of 50shr transaction fee",
	)
	s.Require().Equalf(outStandingRewardBefore.Rewards.AmountOf(denom.Base).Add(sdk.NewDec(25*denom.ShrExponent)).String(), outStandingRewardAfter.Rewards.AmountOf(denom.Base).String(), "the validator must take 50% transaction fee from 50shr fee")
}
