package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/id/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdIdById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "id-by-id [id]",
		Short: "Query IdById",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
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
