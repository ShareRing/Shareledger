package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/document/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
)

var _ = strconv.Itoa(0)

func CmdDocumentByHolderId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "holder [holder id]",
		Short: "Get all documents of a holder.",
		Long: strings.TrimSpace(fmt.Sprintf(`
Get all docs of a holder.
Example:
$ %s query %s holder uid-11594`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDocumentByHolderIdRequest{

				Id: reqId,
			}

			res, err := queryClient.DocumentByHolderId(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
