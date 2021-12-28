package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/booking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdBooking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [bookID]",
		Short: "Query booking",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBookID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryBookingRequest{

				BookID: reqBookID,
			}

			res, err := queryClient.Booking(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
