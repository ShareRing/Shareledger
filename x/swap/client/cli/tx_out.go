package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out [dest-addr] [network] [amount] [fee]",
		Short: "create new request swap out",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDestAddr := args[0]
			argNetwork := args[1]
			argAmount := args[2]
			argFee := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOut(
				clientCtx.GetFromAddress().String(),
				argDestAddr,
				argNetwork,
				argAmount,
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
