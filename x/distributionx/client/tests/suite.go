package tests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil2 "github.com/cosmos/cosmos-sdk/x/bank/client/testutil"
	distributionxType "github.com/sharering/shareledger/x/distributionx/types"
	"github.com/sharering/shareledger/x/electoral/client/tests"
	"github.com/sharering/shareledger/x/utils/denom"
	"os"

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
	disXGen.BuilderListCount = 1
	bz := s.cfg.Codec.MustMarshalJSON(&disXGen)
	s.cfg.GenesisState[distributionxType.ModuleName] = bz

	s.network, _ = network.New(s.T(), s.dir, s.cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)

	//sdk.AccAddress(accs[i].PubKey.Address())

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
	//Execute the contract with fee

	accByte, err := testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].Address)
	s.NoError(err)

	validatorAccBalance := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		netutilts.Accounts[netutilts.KeyAccount2])
	s.NoError(err)
	makeTransactionAccBalance := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
	s.T().Log("validator balance", validatorAccBalance, s.network.Validators[0].Address)
	s.T().Log("contract execute balance", makeTransactionAccBalance, netutilts.Accounts[netutilts.KeyAccount2])

	out, err := ExCmdExecuteContract(
		s.network.Validators[0].ClientCtx,
		"shareledger14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9smpkm95",
		`{"increment":{}}`,
		netutilts.JSONFlag,
		netutilts.SkipConfirmation,
		netutilts.MakeByAccount(netutilts.KeyAccount2),
		netutilts.BlockBroadcast,
		netutilts.SHRFee(50),
	)
	s.Require().NoError(err, "execute the contract fail")
	res := netutilts.ParseStdOut(s.T(), out.Bytes())

	s.Equalf(netutilts.ShareLedgerSuccessCode, res.Code, "broadcast transaction fail %v", res.String())
	_ = s.network.WaitForNextBlock()

	contractOwnerReward, err := ExCmdQueryReward(s.network.Validators[0].ClientCtx, "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw")
	s.NoError(err)

	s.T().Log("contract owner reward", contractOwnerReward, "shareledger1hq7wjjgeymvs3q4vmkvac3dghfsjwvjvf8jdaw")

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		TinyAddr("shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr"))
	s.NoError(err)

	devPool := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	s.T().Log("dev pool balance", devPool, "shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr")

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		netutilts.Accounts[netutilts.KeyAccount2])
	s.NoError(err)
	makeTransactionAccBalanceAfterEx := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
	s.T().Log("contract execute balance after", makeTransactionAccBalance, netutilts.Accounts[netutilts.KeyAccount2])

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].Address)
	s.NoError(err)

	validatorAccBalanceAfter := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())
	//Assertion time
	s.Require().Equalf(sdk.NewInt(6250000000).String(), contractOwnerReward.Reward.GetAmount().AmountOf(denom.Base).String(), "the contract owner reward must be increment")
	s.Require().Equalf(sdk.NewInt(12500000000).String(), devPool.Balances.AmountOf(denom.Base).String(), "Devpool must take 25%")
	s.Require().Equalf(makeTransactionAccBalance.Balances.AmountOf(denom.Base).Sub(sdk.NewInt(50*denom.ShrExponent)), makeTransactionAccBalanceAfterEx.Balances.AmountOf(denom.Base), "the transaction execute maker must be reduce by the fee that input")
	s.Require().Equalf(validatorAccBalance.Balances.AmountOf(denom.Base).Add(sdk.NewInt(25*denom.ShrExponent)).String(), validatorAccBalanceAfter.Balances.AmountOf(denom.Base).String(), "the validator must take 50% transaction fee from 50shr fee")

}
func (s *DistributionXIntegrationTestSuite) TestDistributionXNormalTransaction() {
	accByte, err := testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		TinyAddr("shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr"))
	s.NoError(err)

	devPoolBalanceBeforeTxn := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].Address)
	s.NoError(err)

	validatorBalanceBeforeTxn := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	out, err := tests.ExCmdEnrollDocIssuer(s.network.Validators[0].ClientCtx,
		[]string{"shareledger19w2g2pdcwpj5kutn6ve3r4ay5twptlejlcvkq8"},
		netutilts.SHRFee(50),
		netutilts.MakeByAccount(netutilts.KeyOperator),
		netutilts.SkipConfirmation,
		netutilts.BlockBroadcast,
	)
	if err != nil {
		s.Require().NoError(err, "init doc issuer fail")
	}
	s.Require().NoError(s.network.WaitForNextBlock())
	txResponse := netutilts.ParseStdOut(s.T(), out.Bytes())
	s.Equal(netutilts.ShareLedgerSuccessCode, txResponse.Code, "%s", out.String())

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		TinyAddr("shareledger1t3g4570e23h96h5hm5gdtfrjprmvk9qwmrglfr"))
	s.NoError(err)

	devPoolBalanceAfterTxn := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	accByte, err = testutil2.QueryBalancesExec(
		s.network.Validators[0].ClientCtx,
		s.network.Validators[0].Address)
	s.NoError(err)

	validatorBalanceAfterTxn := netutilts.BalanceJsonUnmarshal(s.T(), accByte.Bytes())

	s.Require().Equalf(
		devPoolBalanceBeforeTxn.
			Balances.
			AmountOf(denom.Base).
			Add(sdk.NewInt(25*denom.ShrExponent)).String(),
		devPoolBalanceAfterTxn.Balances.AmountOf(denom.Base).String(),
		"dev pool account must take 50% of 50shr transaction fee",
	)

	s.Require().Equalf(
		validatorBalanceBeforeTxn.
			Balances.
			AmountOf(denom.Base).
			Add(sdk.NewInt(25*denom.ShrExponent)).String(),
		validatorBalanceAfterTxn.Balances.AmountOf(denom.Base).String(),
		"dev pool account must take 50% of 50shr transaction fee",
	)
}
