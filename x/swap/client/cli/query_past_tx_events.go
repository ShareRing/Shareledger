package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPastTxEvents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "past-tx-events <empty>|<txHash>|<txHash> <logIndex>",
		Short: "get requested swaps in information",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 0 {
				// PastTxEvents
				pageReq, err := client.ReadPageRequest(cmd.Flags())
				if err != nil {
					return err
				}
				params := &types.QueryPastTxEventsRequest{
					Pagination: pageReq,
				}
				res, err := queryClient.PastTxEvents(cmd.Context(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			} else if len(args) == 1 {
				// PastTxEventsByTxHash
				params := &types.QueryPastTxEventsByTxHashRequest{
					TxHash: args[0],
				}
				res, err := queryClient.PastTxEventsByTxHash(cmd.Context(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			} else {
				// PastTxEvent
				logIndex, err := strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return err
				}
				params := &types.QueryPastTxEventRequest{
					TxHash:   args[0],
					LogIndex: logIndex,
				}

				res, err := queryClient.PastTxEvent(cmd.Context(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "past-tx-events")

	return cmd
}
