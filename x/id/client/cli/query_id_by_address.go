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

func CmdIdByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info-by-address [address]",
		Short: "Query for id information",
		Long: strings.TrimSpace(fmt.Sprintf(`
Query id information of an account by owner address
Example:
$ %s query %s info-by-address shareledger1s432u6zv95wpluxhf4qru2ewy58kc3w4tkzm3v`, version.Name, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIdByAddressRequest{

				Address: reqAddress,
			}

			res, err := queryClient.IdByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
