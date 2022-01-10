package cli

import (
    

    "github.com/spf13/cobra"
    

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func CmdCreateActionLevelFee() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-action-level-fee [action] [level]",
		Short: "Create a new action-level-fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexAction := args[0]
        
            // Get value arguments
		 argLevel := args[1]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateActionLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexAction,
                argLevel,
			    )
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdUpdateActionLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-action-level-fee [action] [level]",
		Short: "Update a action-level-fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexAction := args[0]
        
            // Get value arguments
		 argLevel := args[1]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateActionLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexAction,
                argLevel,
                )
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}

func CmdDeleteActionLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-action-level-fee [action]",
		Short: "Delete a action-level-fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
             indexAction := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteActionLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexAction,
                )
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}