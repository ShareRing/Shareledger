package cli

import (
    "context"
	

	
    "github.com/spf13/cobra"
    
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/sharering/shareledger/x/gentlemint/types"
)

func CmdListLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-level-fee",
		Short: "list all level-fee",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllLevelFeeRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.LevelFeeAll(context.Background(), params)
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

func CmdShowLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-level-fee [level]",
		Short: "shows a level-fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argLevel := args[0]
            
            params := &types.QueryGetLevelFeeRequest{
                Level: argLevel,
                
            }

            res, err := queryClient.LevelFee(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
