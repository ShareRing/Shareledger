package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/swap/types"
)

var _ = strconv.Itoa(0)

func CmdNextBatchId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-batch-id",
		Short: "Query next batch id that Id will assign to next swap batch",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryNextBatchIdRequest{}

			res, err := queryClient.NextBatchId(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
