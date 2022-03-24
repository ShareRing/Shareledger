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

func CmdDocumentOfHolderByIssuer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "document-by-issuer [holder] [issuer]",
		Short: "Get all document of a holder issue by a specific issuer",
		Long: strings.TrimSpace(fmt.Sprintf(`
Example:
$ %s query %s document-by-issuer uuid-5312 shareledger1z8h2ymcpr7w75l4pwxkgr3fmg6wqv9fk5mrc79`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqHolder := args[0]
			reqIssuer := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDocumentOfHolderByIssuerRequest{

				Holder: reqHolder,
				Issuer: reqIssuer,
			}

			res, err := queryClient.DocumentOfHolderByIssuer(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
