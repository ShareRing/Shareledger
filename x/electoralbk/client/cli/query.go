package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ShareRing/Shareledger/x/electoralbk/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group electoralbk queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdGetVoter(types.QuerierRoute),
	)

	return cmd
}

func CmdGetVoter(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "get [address]",
		Short: "return status of voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// cliCtx := context.NewCLIContext().WithCodec(cdc)
			// voter := args[0]
			// voterID := "voter" + voter
			// res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/voter/%s", queryRoute, voterID), nil)
			// if err != nil {
			// 	fmt.Printf("could not get voter - %s \n", voterID)
			// 	return nil
			// }

			// var out types.Voter
			// cdc.MustUnmarshalJSON(res, &out)
			// return cliCtx.PrintOutput(out)

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			voter := args[0]
			voterID := types.VoterPrefix + voter

			queryClient := types.NewQueryClient(cliCtx)
			params := types.NewQueryVoterRequest(voterID)
			res, err := queryClient.Voter(cmd.Context(), params)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
}
