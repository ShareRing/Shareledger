package cli

import (
	"fmt"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func QueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	bookingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the electoral module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	bookingQueryCmd.AddCommand(flags.GetCommands(
		CmdGetVoter(storeKey, cdc),
	)...)

	return bookingQueryCmd
}

func CmdGetVoter(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [address]",
		Short: "return status of voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			voter := args[0]
			voterID := "voter" + voter
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/voter/%s", queryRoute, voterID), nil)
			if err != nil {
				fmt.Printf("could not get voter - %s \n", voterID)
				return nil
			}

			var out types.Voter
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
