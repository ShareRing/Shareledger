package gov

import (
	"fmt"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	paramchangecli "github.com/cosmos/cosmos-sdk/x/params/client/cli"
	"github.com/sharering/shareledger/tests"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/spf13/cobra"

	dtxtypes "github.com/sharering/shareledger/x/distributionx/types"
)

var distributionxParams = &dtxtypes.Params_ConfigPercent{
	WasmMasterBuilder: sdk.NewDecWithPrec(150, 3),
	WasmContractAdmin: sdk.NewDecWithPrec(250, 3),
	WasmDevelopment:   sdk.NewDecWithPrec(400, 3),
	WasmValidator:     sdk.NewDecWithPrec(200, 3),
	NativeValidator:   sdk.NewDecWithPrec(600, 3),
	NativeDevelopment: sdk.NewDecWithPrec(400, 3),
}

func (s *E2ETestSuite) QueryDistributionxParams() {
	params := dtxtypes.DefaultParams()
	params.ConfigPercent = distributionxParams
	devPoolAccount = network.MustAddressFormKeyring(s.network.Validators[0].ClientCtx.Keyring, network.KeyAccount3).String()
	params.BuilderWindows = 15
	params.TxThreshold = 3
	params.DevPoolAccount = devPoolAccount
	val := s.network.Validators[0]
	buildUrl := func(surfix string) string {
		return fmt.Sprintf("%s/sharering/shareledger/distributionx/%s", val.APIAddress, surfix)
	}
	testCases := tests.TestCasesGrpc{
		{
			Name:      "gRPC get params",
			URL:       buildUrl("params"),
			Headers:   map[string]string{},
			ExpectErr: false,
			RespType:  &dtxtypes.QueryParamsResponse{},
			Expected: &dtxtypes.QueryParamsResponse{
				Params: params,
			},
		},
	}
	tests.RunTestCasesGrpc(&s.Suite, testCases, val)
}

func (s *E2ETestSuite) TestChangeDistributionxParams() {
	s.UpdateDistributionxParams()
	s.VoteProposal()
	fmt.Printf("Waiting 60 for proposal to finish...\n")
	time.Sleep(time.Second * 70)
	s.QueryDistributionxParams()

	fmt.Printf("Start testing for distributionx module using new params ...\n")
	s.DistributionxTest()
}

func (s *E2ETestSuite) SubmitProposal(cmd *cobra.Command, args ...string) {
	args = append(args, network.MakeByAccount(network.KeyAuthority))
	testCases := tests.TestCasesTx{
		{
			Name:         "submit new proposal",
			Args:         args,
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cmd, s.network.Validators[0])
	globalProposalId = globalProposalId + 1
}

func (s *E2ETestSuite) VoteProposal() {
	valAddress := s.network.Validators[0].Address
	testCases := tests.TestCasesTx{
		{
			Name: "vote yes new proposal",
			Args: []string{
				strconv.Itoa(globalProposalId),
				"yes",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, valAddress.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			},
			ExpectErr:    false,
			ExpectedCode: errorsmod.SuccessABCICode,
		},
	}
	tests.RunTestCasesTx(&s.Suite, testCases, cli.NewCmdVote(), s.network.Validators[0])
}

func (s *E2ETestSuite) UpdateDistributionxParams() {
	args := []string{
		"./proposal.json",
	}
	cmd := paramchangecli.NewSubmitParamChangeProposalTxCmd()
	flags.AddTxFlagsToCmd(cmd)
	s.SubmitProposal(cmd, args...)
}

func (s *E2ETestSuite) DistributionxTest() {
	s.CheckRewardWasmTx(4_000_000_000)
	s.CheckRewardNormalTx(4_000_000_000)
	s.CheckWithdrawReward(10_000_000_000)
}
