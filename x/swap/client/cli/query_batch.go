package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdShowBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show_batch [ids]",
		Short: "shows a batch base on list of IDs = 1,2,3,4",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			idsArgs := args[0]
			idsStr := strings.Split(idsArgs, ",")
			ids := make([]uint64, 0, len(idsStr))
			for i := range idsStr {
				id, err := strconv.ParseUint(idsStr[i], 10, 64)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}

			params := &types.QueryGetBatchRequest{
				Ids: ids,
			}

			res, err := queryClient.Batch(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
