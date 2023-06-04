package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwapManager() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-manager [address]",
		Short: "Query SwapManager",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySwapManagerRequest{Address: reqAddress}

			res, err := queryClient.SwapManager(cmd.Context(), params)
			if err != nil {
				fmt.Print(err)
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
