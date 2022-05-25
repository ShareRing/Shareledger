package cli

import (
    "context"
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/sharering/shareledger/x/swap/types"
)

func CmdListRequestedIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-requested-in",
		Short: "list all requestedIn",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllRequestedInRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.RequestedInAll(context.Background(), params)
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

func CmdShowRequestedIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-requested-in [address]",
		Short: "shows a requestedIn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argAddress := args[0]
            
            params := &types.QueryGetRequestedInRequest{
                Address: argAddress,
                
            }

            res, err := queryClient.RequestedIn(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
