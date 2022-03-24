package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/sharering/shareledger/x/id/types"
)

var _ = strconv.Itoa(0)

func CmdIdById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info-by-id [id]",
		Short: "Query for id information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query id information of an account by owner id
Example:
$ %s query %s info-by-id 123e4567-e89b-12d3-a456-426655440000`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIdByIdRequest{

				Id: reqId,
			}

			res, err := queryClient.IdById(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
