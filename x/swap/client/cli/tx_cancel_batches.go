package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/swap/types"
)

var _ = strconv.Itoa(0)

func CmdCancelBatches() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [ids]",
		Short: "Broadcast message cancel batches",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argIDS := strings.Split(args[0], ",")
			var txIds = make([]uint64, 0, len(argIDS))
			for _, str := range argIDS {
				id, err := strconv.ParseUint(str, 10, 64)
				if err != nil {
					return err
				}
				txIds = append(txIds, id)
			}
			msg := types.NewMsgCancelBatches(
				clientCtx.GetFromAddress().String(),
				txIds,
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
