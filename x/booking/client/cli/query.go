package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/sharering/shareledger/x/booking/types"
	"github.com/spf13/cobra"
)

func QueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	bookingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the booking module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	bookingQueryCmd.AddCommand(flags.GetCommands(
		CmdGetBooking(storeKey, cdc),
	)...)

	return bookingQueryCmd
}

func CmdGetBooking(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [bookID]",
		Short: "get booking info from bookID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bookID := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/booking/%s", queryRoute, bookID), nil)
			if err != nil {
				fmt.Printf("could not get asset - %s \n", bookID)
				return nil
			}

			var out types.Booking
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
