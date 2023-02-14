package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/spf13/cobra"
)

func CmdShowMinimumGasPrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "minimum-gas-prices",
		Short:   "Show minimum gas prices",
		Long:    "Show all minimum gas prices",
		Aliases: []string{"min"},
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.MinimumGasPrices(cmd.Context(), &types.QueryMinimumGasPricesRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
