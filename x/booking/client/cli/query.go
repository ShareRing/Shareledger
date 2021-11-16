package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/booking/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group booking queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdGetBooking(types.QuerierRoute),
	)

	return cmd
}

func CmdGetBooking(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [bookID]",
		Short: "get booking info by bookID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			bookID := args[0]

			// res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/booking/%s", queryRoute, bookID), nil)
			// if err != nil {
			// 	fmt.Printf("could not get asset - %s \n", bookID)
			// 	return nil
			// }

			queryClient := types.NewQueryClient(cliCtx)
			params := types.QueryBookingRequest{bookID}
			res, err := queryClient.Booking(cmd.Context(), &params)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
