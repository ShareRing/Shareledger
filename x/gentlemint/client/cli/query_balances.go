package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/gentlemint/types"
	"github.com/spf13/cobra"
)

func CmdBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances",
		Short: "Query balances from [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return err
			}

			params := &types.QueryBalancesRequest{
				Address: args[0],
			}

			res, err := queryClient.Balances(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
