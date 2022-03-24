package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func CmdListActionLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action-level-fees",
		Short: "List all action-level-fee",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryActionLevelFeesRequest{}

			res, err := queryClient.ActionLevelFees(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowActionLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action-fee [action]",
		Short: "shows a action-level-fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAction := args[0]

			params := &types.QueryActionLevelFeeRequest{
				Action: argAction,
			}

			res, err := queryClient.ActionLevelFee(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
