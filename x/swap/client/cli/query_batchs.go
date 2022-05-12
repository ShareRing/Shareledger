package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdShowBatches() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show_batches [ids]",
		Short:   "shows a batches base on list of IDs",
		Example: "show_batches 1 2 3 ",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			idsArgs := args[:]
			ids := make([]uint64, 0, len(idsArgs))
			for i := range idsArgs {
				id, err := strconv.ParseUint(idsArgs[i], 10, 64)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}

			params := &types.QueryBatchesRequest{
				Ids: ids,
			}

			res, err := queryClient.Batches(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
