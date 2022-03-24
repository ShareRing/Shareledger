package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/sharering/shareledger/x/document/types"
)

var _ = strconv.Itoa(0)

func CmdDocumentByProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof [proof]",
		Short: "Query for document information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query document information by the proof.
Example:
$ %s query %s proof 5wpluxhf4qru2ewy58kc3w4tkzm3v`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqProof := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDocumentByProofRequest{

				Proof: reqProof,
			}

			res, err := queryClient.DocumentByProof(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
