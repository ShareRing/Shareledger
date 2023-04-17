package tests

import (
	"github.com/CosmWasm/wasmd/x/wasm/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	cli3 "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	types2 "github.com/cosmos/cosmos-sdk/x/distribution/types"
	cli4 "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	types3 "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	"github.com/sharering/shareledger/testutil/network"
	cli2 "github.com/sharering/shareledger/x/distributionx/client/cli"
	"github.com/sharering/shareledger/x/distributionx/types"
)

func ExCmdStoreCode(clientCtx client.Context, codePath string, additionalFlags ...string) (testutil.BufferWriter, error) {
	var args []string
	args = append(args, codePath)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.StoreCodeCmd(), args)
}

func ExCmdInstantContract(clientCtx client.Context, codeID string, jsonEncoded string, additionalFlags ...string) (testutil.BufferWriter, error) {
	var args []string
	args = append(args, codeID)
	args = append(args, jsonEncoded)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.InstantiateContractCmd(), args)
}

func ExCmdExecuteContract(clientCtx client.Context, contractAddr string, exec string, additionalFlags ...string) (testutil.BufferWriter, error) {
	var args []string
	args = append(args, contractAddr)
	args = append(args, exec)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.ExecuteContractCmd(), args)
}

func ExCmdQueryReward(clientCtx client.Context, rewardAddr string) (types.QueryGetRewardResponse, error) {
	var args []string
	args = append(args, rewardAddr)
	args = append(args, network.JSONFlag)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli2.CmdShowReward(), args)
	if err != nil {
		return types.QueryGetRewardResponse{}, errors.Wrapf(err, "addr %s", rewardAddr)
	}

	var res = types.QueryGetRewardResponse{}
	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	if err != nil {
		return types.QueryGetRewardResponse{}, err
	}
	return res, nil
}
func ExCmdListReward(clientCtx client.Context) (types.QueryAllRewardResponse, error) {
	var args []string
	//args = append(args, masterBuilder)
	args = append(args, network.JSONFlag)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli2.CmdListReward(), args)
	if err != nil {
		return types.QueryAllRewardResponse{}, err
	}

	var res = types.QueryAllRewardResponse{}
	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	if err != nil {
		return types.QueryAllRewardResponse{}, err
	}
	return res, nil
}
func ExCmdQueryParam(clientCtx client.Context) (types.QueryParamsResponse, error) {
	var args []string
	//args = append(args, masterBuilder)
	args = append(args, network.JSONFlag)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli2.CmdQueryParams(), args)
	if err != nil {
		return types.QueryParamsResponse{}, err
	}

	var res = types.QueryParamsResponse{}
	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	if err != nil {
		return types.QueryParamsResponse{}, err
	}
	return res, nil
}

func ExCmdQueryOutStandingReward(clientCtx client.Context, validator string) (types2.ValidatorOutstandingRewards, error) {
	var args []string

	args = append(args, validator)
	args = append(args, network.JSONFlag)

	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli3.GetCmdQueryValidatorOutstandingRewards(), args)
	if err != nil {
		return types2.ValidatorOutstandingRewards{}, err
	}
	var res = types2.ValidatorOutstandingRewards{}

	err = clientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	if err != nil {
		return types2.ValidatorOutstandingRewards{}, err
	}
	return res, nil
}

func ExCmdListDelegator(clientCtx client.Context, validator string) (types3.QueryValidatorDelegationsResponse, error) {
	var args []string
	args = append(args, validator)

	args = append(args, network.JSONFlag)
	//GetCmdQueryValidatorDelegations
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli4.GetCmdQueryValidatorDelegations(), args)
	if err != nil {
		return types3.QueryValidatorDelegationsResponse{}, err
	}
	//QueryValidatorDelegationsResponse
	var res = types3.QueryValidatorDelegationsResponse{}
	err = clientCtx.Codec.Unmarshal(out.Bytes(), &res)
	if err != nil {
		return types3.QueryValidatorDelegationsResponse{}, err
	}
	return res, nil
}
