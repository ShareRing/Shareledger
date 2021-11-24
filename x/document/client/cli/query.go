package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/ShareRing/Shareledger/x/document/types"
)

func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetDocByProofCmd(types.QuerierRoute),
		GetDocByHolderCmd(types.QuerierRoute),
	)

	return cmd
}

func GetDocByProofCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof <proof>",
		Short: "Query for doc information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query document information by the proof.
Example:
$ %s query %s proof 5wpluxhf4qru2ewy58kc3w4tkzm3v`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			params := types.QueryDocumentByProofRequest{Proof: args[0]}
			res, err := queryClient.DocumentByProof(cmd.Context(), &params)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetDocByHolderCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "holder <holder id>",
		Short: "Get all docs of a holder.",
		Long: strings.TrimSpace(fmt.Sprintf(`
Get all docs of a holder.
Example:
$ %s query %s holder uid-11594`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(cliCtx)
			params := types.QueryDocumentByHolderIdRequest{Id: args[0]}
			res, err := queryClient.DocumentByHolderId(cmd.Context(), &params)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
