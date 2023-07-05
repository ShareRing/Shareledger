package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/swap/types"
)

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [amount]",
		Short: "Broadcast deposit SHR to swapping module, Note: just accept shr for now",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]

			decCoins, err := sdk.ParseDecCoin(argAmount)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(
				clientCtx.GetFromAddress().String(),
				&decCoins,
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
