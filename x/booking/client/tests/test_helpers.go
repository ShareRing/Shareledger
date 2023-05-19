package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/sharering/shareledger/app"
	"github.com/sharering/shareledger/testutil/network"
	"github.com/sharering/shareledger/x/booking/client/cli"
	"github.com/sharering/shareledger/x/booking/types"
	"github.com/stretchr/testify/require"
)

func ExCmdCreateBooking(clientCtx client.Context, assetUUID, duration string, additionalFlags ...string) (testutil.BufferWriter, error) {
	args := []string{assetUUID, duration}
	args = append(args, network.SkipConfirmation, network.SyncBroadcast)
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
	args = append(args, network.SkipConfirmation, network.SyncBroadcast)
	args = append(args, additionalFlags...)
	return clitestutil.ExecTestCLICmd(clientCtx, cli.CmdComplete(), args)
}

func BookingJsonUnmarshal(t *testing.T, data []byte) types.Booking {
	var a types.QueryBookingResponse
	encodingConfig := app.MakeTestEncodingConfig()
	err := encodingConfig.Codec.UnmarshalJSON(data, &a)
	require.NoError(t, err)
	return *a.Booking
}
