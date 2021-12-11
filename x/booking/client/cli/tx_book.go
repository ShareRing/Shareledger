package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/booking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "book [uuid] [duration]",
		Short: "Broadcast message Book",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUUID := args[0]
			argDuration := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBook(
				clientCtx.GetFromAddress().String(),
				argUUID,
				argDuration,
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
