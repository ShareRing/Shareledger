package cli

import (
    

    "github.com/spf13/cobra"
    

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

func CmdCreateLevelFee() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-level-fee [level] [fee]",
		Short: "Create a new level-fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexLevel := args[0]
        
            // Get value arguments
		 argFee := args[1]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexLevel,
                argFee,
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

func CmdUpdateLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-level-fee [level] [fee]",
		Short: "Update a level-fee",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexLevel := args[0]
        
            // Get value arguments
		 argFee := args[1]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexLevel,
                argFee,
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

func CmdDeleteLevelFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-level-fee [level]",
		Short: "Delete a level-fee",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
             indexLevel := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLevelFee(
			    clientCtx.GetFromAddress().String(),
			    indexLevel,
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