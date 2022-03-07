package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/x/id/client/cli"
	"github.com/sharering/shareledger/x/id/types"
)

func CmdExNewID(clientCtx client.Context, userID, backupAddress, addressOwner, exData string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{userID, backupAddress, addressOwner, exData}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateId(), args)
}

func CmdExNewIDInBatch(clientCtx client.Context, userIDs, backupAddresses, addressOwners, exDatas string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{userIDs, backupAddresses, addressOwners, exDatas}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateIds(), args)
}

func CmdExGetID(clientCtx client.Context, t *testing.T, userID string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userID}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdIdById(), args)
	if err != nil {
		t.Errorf("fail get id: %v", err)
	}

	return out
}

func CmdExUpdateID(clientCtx client.Context, userID, exData string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{userID, exData}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdUpdateId(), args)
}

func CmdExReplaceIdOwner(clientCtx client.Context, userID, newAddress string, extraFlags ...string) (testutil.BufferWriter, error) {
	args := []string{userID, newAddress}
	args = append(args, extraFlags...)

	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdReplaceIdOwner(), args)
}

func CmdExGetIDByAddress(clientCtx client.Context, t *testing.T, addr string, extraFlags ...string) testutil.BufferWriter {
	args := []string{addr}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdIdByAddress(), args)
	if err != nil {
		t.Errorf("fail get id: %v", err)
	}

	return out
}

func IDJsonUnmarshal(t *testing.T, data []byte) types.Id {
	var a types.QueryIdByIdResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(data, &a)
	require.NoError(t, err)
	return *a.Id

}
