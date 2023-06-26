package gentlemint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/gentlemint/client/cli"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/sharering/shareledger/x/utils/denom"
)

func (s *E2ETestSuite) TestQueryActionLevelFee() {
	testCases := tests.TestCases{
		{
			Name:      "query list action level fee",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryActionLevelFeesResponse{},
			Expected: &types.QueryActionLevelFeesResponse{
				// ignore the last item that used to test delete
				ActionLevelFee: defaultActionLevelFees[:2],
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdListActionLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestShowActionLevelFee() {
	testCases := tests.TestCases{
		{
			Name: "query action level fee ok",
			Args: []string{
				defaultActionLevelFees[2].Action,
			},
			ExpectErr: false,
			RespType:  &types.QueryActionLevelFeeResponse{},
			Expected: &types.QueryActionLevelFeeResponse{
				Action: defaultActionLevelFees[2].Action,
				Level:  defaultActionLevelFees[2].Level,
				Fee:    defaultLevelFees[2].Fee.String(),
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowActionLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryBalances() {
	oneThousandSHRDecCoin := sdk.NewDecCoin(denom.Shr, sdk.NewInt(10000))
	oneHundredSHRPDecCoin := sdk.NewDecCoin(denom.ShrP, sdk.NewInt(100))

	// should not check balances of any account that used to sign transaction, because its balances will be changed
	// at runtime when the test run.
	authorityAcc := network.MustAddressFormKeyring(s.network.Validators[0].ClientCtx.Keyring, network.KeyAccount8)
	testCases := tests.TestCases{
		{
			Name: "query balances ok",
			Args: []string{
				authorityAcc.String(),
			},
			ExpectErr: false,
			RespType:  &types.QueryBalancesResponse{},
			Expected: &types.QueryBalancesResponse{
				Coins: []*sdk.DecCoin{&oneThousandSHRDecCoin, &oneHundredSHRPDecCoin},
			},
		},
		{
			Name: "query balances invlid address",
			Args: []string{
				"testAddress",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdBalances(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestCheckFees() {
	docIssuerAcc := network.MustAddressFormKeyring(s.network.Validators[0].ClientCtx.Keyring, network.KeyDocIssuer)
	totalFee := lowConvertedFee.Add(minConvertedFee)
	testCases := tests.TestCases{
		{
			Name: "query check fee",
			Args: []string{
				docIssuerAcc.String(),
				"gentlemint_load",
			},
			ExpectErr: false,
			RespType:  &types.QueryCheckFeesResponse{},
			Expected: &types.QueryCheckFeesResponse{
				ConvertedFee:         &lowConvertedFee,
				SufficientFee:        true,
				SufficientFundForFee: true,
				CostLoadingFee:       nil,
			},
		},
		{
			Name: "query check fee invalid request",
			Args: []string{
				"",
			},
			ExpectErr: true,
		},
		{
			Name: "query check fee multi action",
			Args: []string{
				docIssuerAcc.String(),
				"gentlemint_load",
				"gentlemint_send",
			},
			ExpectErr: false,
			RespType:  &types.QueryCheckFeesResponse{},
			Expected: &types.QueryCheckFeesResponse{
				ConvertedFee:         &totalFee,
				SufficientFee:        true,
				SufficientFundForFee: true,
				CostLoadingFee:       nil,
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdCheckFees(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestQueryExchangeRate() {
	testCases := tests.TestCases{
		{
			Name:      "query exchange rate",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryExchangeRateResponse{},
			Expected: &types.QueryExchangeRateResponse{
				Rate: defaultRate,
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowExchangeRate(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestShowMiniMumgasPrices() {
	testCases := tests.TestCases{
		{
			Name:      "query minimum gas price",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryMinimumGasPricesResponse{},
			Expected: &types.QueryMinimumGasPricesResponse{
				MinimumGasPrices: params.MinimumGasPrices,
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowMinimumGasPrices(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestShowLevelFee() {
	testCases := tests.TestCases{
		{
			Name: "query level fee",
			Args: []string{
				"low",
			},
			ExpectErr: false,
			RespType:  &types.QueryLevelFeeResponse{},
			Expected: &types.QueryLevelFeeResponse{
				LevelFee: types.LevelFeeDetail{
					Level:        defaultLevelFees[0].Level,
					Creator:      defaultLevelFees[0].Creator,
					OriginalFee:  defaultLevelFees[0].Fee,
					ConvertedFee: &lowConvertedFee,
				},
			},
		},
		{
			Name: "query invalid level fee",
			Args: []string{
				"test",
			},
			ExpectErr: true,
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdShowLevelFee(), s.network.Validators[0])
}

func (s *E2ETestSuite) TestLevelFees() {
	testCases := tests.TestCases{
		{
			Name:      "query list level fees",
			Args:      []string{},
			ExpectErr: false,
			RespType:  &types.QueryLevelFeesResponse{},
			Expected: &types.QueryLevelFeesResponse{
				LevelFees: []types.LevelFeeDetail{
					{
						Level:        defaultLevelFees[0].Level,
						Creator:      defaultLevelFees[0].Creator,
						OriginalFee:  defaultLevelFees[0].Fee,
						ConvertedFee: &lowConvertedFee,
					},
					{
						Level:        defaultLevelFees[1].Level,
						Creator:      defaultLevelFees[1].Creator,
						OriginalFee:  defaultLevelFees[1].Fee,
						ConvertedFee: &mediumConvertedFee,
					},
					{
						Level:        defaultLevelFees[2].Level,
						Creator:      defaultLevelFees[2].Creator,
						OriginalFee:  defaultLevelFees[2].Fee,
						ConvertedFee: &highConvertedFee,
					},
					{
						Level:        levelZero.Level,
						OriginalFee:  levelZero.Fee,
						ConvertedFee: &zeroConvertedFee,
					},
					{
						Level:        levelMin.Level,
						OriginalFee:  levelMin.Fee,
						ConvertedFee: &minConvertedFee,
					},
				},
			},
		},
	}
	tests.RunTestCases(&s.Suite, testCases, cli.CmdListLevelFee(), s.network.Validators[0])
}
