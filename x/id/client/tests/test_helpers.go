package tests

import (
	"github.com/ShareRing/Shareledger/app"
	"github.com/ShareRing/Shareledger/x/id/client/cli"
	"github.com/ShareRing/Shareledger/x/id/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"
	"testing"
)

func CmdExNewID(clientCtx client.Context, t *testing.T, userID, backupAddress, addressOwner, exData string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userID, backupAddress, addressOwner, exData}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateId(), args)
	if err != nil {
		t.Errorf("fail create id: %v", err)
	}

	return out
}

func CmdExNewIDInBatch(clientCtx client.Context, t *testing.T, userIDs, backupAddresses, addressOwners, exDatas string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userIDs, backupAddresses, addressOwners, exDatas}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdCreateIdBatch(), args)
	if err != nil {
		t.Errorf("fail create id: %v", err)
	}

	return out
}

func CmdExGetID(clientCtx client.Context, t *testing.T, userID string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userID}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdIdById(), args)
	if err != nil {
		t.Errorf("fail create id: %v", err)
	}

	return out
}

func CmdExUpdateID(clientCtx client.Context, t *testing.T, userID, exData string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userID, exData}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdUpdateId(), args)
	if err != nil {
		t.Errorf("fail update id: %v", err)
	}

	return out
}

func CmdExReplaceIdOwner(clientCtx client.Context, t *testing.T, userID, newAddress string, extraFlags ...string) testutil.BufferWriter {
	args := []string{userID, newAddress}
	args = append(args, extraFlags...)
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cli.CmdReplaceIdOwner(), args)
	if err != nil {
		t.Errorf("fail replace id: %v", err)
	}

	return out
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
