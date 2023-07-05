package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/swap/types"
)

func CmdShowSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema [network]",
		Short: "get a schema information by network name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNetwork := args[0]

			params := &types.QuerySchemaRequest{
				Network: argNetwork,
			}

			res, err := queryClient.Schema(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schemas",
		Short: "get a list of available network schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySchemasRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Schemas(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
