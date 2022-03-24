package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/asset/types"
)

var _ = strconv.Itoa(0)

func CmdAssetByUUID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [uuid]",
		Short: "Query AssetByUUID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUUID := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAssetByUUIDRequest{

				Uuid: reqUUID,
			}

			res, err := queryClient.AssetByUUID(cmd.Context(), params)
			if err != nil {
				fmt.Printf("could not get asset - %s, %v \n", reqUUID, err)
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
