package cli

import (
	"fmt"
	"github.com/spf13/pflag"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBatches() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "batches",
		Short:   "Query the swapping batches in our blockchain. You must past at least one filter parameter via a flag use flag --help to get all filter flags",
		Example: fmt.Sprintf("batches --%s 1,2,3", flagSearchIDs),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			request := ReadSearchBatchRequest(cmd.Flags())

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			request.Pagination = pageReq

			res, err := queryClient.Batches(cmd.Context(), request)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "batches")
	cmd.Flags().String(flagSearchDestNetwork, "", "the destination network you want to get")
	cmd.Flags().String(flagSearchIDs, "", "the list of batch ids")

	return cmd
}

func ReadSearchBatchRequest(flagSet *pflag.FlagSet) *types.QueryBatchesRequest {
	destNet, _ := flagSet.GetString(flagSearchDestNetwork)
	idsStr, _ := flagSet.GetString(flagSearchIDs)

	idsStrArr := strings.Split(idsStr, ",")
	var ids []uint64

	for _, id := range idsStrArr {
		i, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, i)
	}

	return &types.QueryBatchesRequest{
		Network: destNet,
		Ids:     ids,
	}
}
