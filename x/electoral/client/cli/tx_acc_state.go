package cli

import (
	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/electoral/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateAccState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-acc-state [key] [address] [status]",
		Short: "Create a new acc-state",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexKey := args[0]

			// Get value arguments
			argAddress := args[1]
			argStatus := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAccState(
				clientCtx.GetFromAddress().String(),
				indexKey,
				argAddress,
				argStatus,
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

func CmdUpdateAccState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-acc-state [key] [address] [status]",
		Short: "Update a acc-state",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexKey := args[0]

			// Get value arguments
			argAddress := args[1]
			argStatus := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAccState(
				clientCtx.GetFromAddress().String(),
				indexKey,
				argAddress,
				argStatus,
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

func CmdDeleteAccState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-acc-state [key]",
		Short: "Delete a acc-state",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAccState(
				clientCtx.GetFromAddress().String(),
				indexKey,
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
