package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var _ = strconv.Itoa(0)

func CmdLoadShr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-pshr [address] [amount]",
		Short: "Mint new [amount] pshr and send [amount] to address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLoadShr(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argAmount,
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
