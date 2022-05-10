package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSearchBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search_batch [status]",
		Short: "Query search-batch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySearchBatchesRequest{}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			filterNetwork, err := cmd.Flags().GetString(flagSearchDestNetwork)
			if err != nil {
				return err
			}

			params.Pagination = pageReq
			params.Status = args[0]
			params.Network = filterNetwork

			res, err := queryClient.SearchBatches(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flagSearchDestNetwork, "", "the destination network you want to get")

	return cmd
}
