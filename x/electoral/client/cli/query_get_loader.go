package cli

import (
	"github.com/ShareRing/Shareledger/x/utils"
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


func CmdGetLoadersFromFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-loaders-from-file [filepath]",
		Short: "get shrp loaders from json file of addresses",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addrList, err := utils.GetAddressFromFile(args[0])

			queryClient := types.NewQueryClient(clientCtx)

			for _, addr := range addrList {
				params := &types.QueryGetLoaderRequest{
					Address: addr,
				}

				res, err := queryClient.GetLoader(cmd.Context(), params)
				if err != nil {
					return err
				}

				if err := clientCtx.PrintProto(res); err != nil {
					return err
				}
			}
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
