package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spm/cosmoscmd"

	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/booking/client/cli"
	"github.com/sharering/shareledger/x/booking/types"
)

func ExCmdCreateBooking(clientCtx client.Context, assetUUID, duration string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetUUID, duration}
	args = append(args, network.SkipConfirmation, network.BlockBroadcast)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdBook(), args)
}

func ExCmdCGetBooking(clientCtx client.Context, bookingID string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{bookingID}
	args = append(args, network.GetFlagsQuery()...)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdBooking(), args)
}

func ExCmdCCompleteBooking(clientCtx client.Context, bookingID string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{bookingID}
	args = append(args, network.SkipConfirmation, network.BlockBroadcast)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdComplete(), args)
}

func BookingJsonUnmarshal(t *testing.T, data []byte) types.Booking {
	var a types.QueryBookingResponse
	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	err := encCfg.Marshaler.UnmarshalJSON(data, &a)
	require.NoError(t, err)
	return *a.Booking

}
