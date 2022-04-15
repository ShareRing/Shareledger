package cli

import (
	"github.com/spf13/pflag"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSearch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [status]",
		Short: "search swaps, status is required for searching. Supported status: pending, approved, rejected, processing, done",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqStatus := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := parseFilterFlag(cmd.Flags())
			params.Status = reqStatus

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			if err := params.BasicValidate(); err != nil {
				return err
			}

			res, err := queryClient.Search(cmd.Context(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	addFilterFlagsToCmd(cmd)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

const (
	flagSearchID          = "id"
	flagSearchSrcAddr     = "src-addr"
	flagSearchDestAddr    = "dest-addr"
	flagSearchSrcNetwork  = "src-network"
	flagSearchDestNetwork = "dest-network"
)

func addFilterFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().Uint64(flagSearchID, 0, "search swap request by id")
	cmd.Flags().String(flagSearchSrcAddr, "", "search swap request by src address")
	cmd.Flags().String(flagSearchDestAddr, "", "search swap request by dest address")
	cmd.Flags().String(flagSearchSrcNetwork, "", "search swap request by src network")
	cmd.Flags().String(flagSearchDestNetwork, "", "search swap request by dest address")
}

func parseFilterFlag(flagSet *pflag.FlagSet) (filter types.QuerySearchRequest) {
	filter.Id, _ = flagSet.GetUint64(flagSearchID)
	filter.SrcAddr, _ = flagSet.GetString(flagSearchSrcAddr)
	filter.DestAddr, _ = flagSet.GetString(flagSearchDestAddr)
	filter.SrcNetwork, _ = flagSet.GetString(flagSearchSrcNetwork)
	filter.DestNetwork, _ = flagSet.GetString(flagSearchDestNetwork)

	return filter
}
