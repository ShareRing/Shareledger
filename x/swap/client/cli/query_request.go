package cli

import (
	"github.com/spf13/pflag"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requests [status]",
		Short: "search swap request, status is required for searching. Supported status: pending, approved",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqStatus := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params, err := parseFilterFlag(cmd.Flags())
			if err != nil {
				return err
			}
			params.Status = reqStatus

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			if err := params.BasicValidate(); err != nil {
				return err
			}

			res, err := queryClient.Swap(cmd.Context(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	addFilterFlagsToCmd(cmd)
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "requests")

	return cmd
}

const (
	flagSearchIDs         = "ids"
	flagSearchSrcAddr     = "src_addr"
	flagSearchDestAddr    = "dest_addr"
	flagSearchSrcNetwork  = "src_network"
	flagSearchDestNetwork = "dest_network"
	flagSearchStatus      = "status"
)

func addFilterFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(flagSearchIDs, "", "search swap request by ids. Supported format: <id0>,<id1>,.. ")
	cmd.Flags().String(flagSearchSrcAddr, "", "search swap request by src address")
	cmd.Flags().String(flagSearchDestAddr, "", "search swap request by dest address")
	cmd.Flags().String(flagSearchSrcNetwork, "", "search swap request by src network")
	cmd.Flags().String(flagSearchDestNetwork, "", "search swap request by dest address")
}

func parseFilterFlag(flagSet *pflag.FlagSet) (filter types.QuerySwapRequest, err error) {
	sIdsV, _ := flagSet.GetString(flagSearchIDs)
	sIds := strings.Split(sIdsV, ",")
	if sIdsV != "" && len(sIds) > 0 {
		filter.Ids = make([]uint64, 0, len(sIds))
		for _, v := range sIds {
			iv, err := strconv.Atoi(v)
			if err != nil {
				return filter, err
			}
			filter.Ids = append(filter.Ids, uint64(iv))
		}
	}

	filter.SrcAddr, _ = flagSet.GetString(flagSearchSrcAddr)
	filter.DestAddr, _ = flagSet.GetString(flagSearchDestAddr)
	filter.SrcNetwork, _ = flagSet.GetString(flagSearchSrcNetwork)
	filter.DestNetwork, _ = flagSet.GetString(flagSearchDestNetwork)

	return filter, nil
}
