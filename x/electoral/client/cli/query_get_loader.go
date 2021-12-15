package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdGetLoader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-loader [address]",
		Short: "get shrp loader from address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLoaderRequest{

				Address: reqAddress,
			}

			res, err := queryClient.GetLoader(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
