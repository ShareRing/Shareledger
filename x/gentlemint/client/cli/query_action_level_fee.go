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
		Use:   "list-action-level-fee",
		Short: "list all action-level-fee",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllActionLevelFeeRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.ActionLevelFeeAll(context.Background(), params)
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

func CmdShowActionLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-action-level-fee [action]",
		Short: "shows a action-level-fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argAction := args[0]
            
            params := &types.QueryGetActionLevelFeeRequest{
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
