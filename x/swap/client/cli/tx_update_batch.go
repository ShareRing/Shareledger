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

func CmdUpdateBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-batch [batch-id] [status] [nonce]",
		Short: "Broadcast message update-batch",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBatchId := args[0]
			argStatus := args[1]
			var argNonce string
			if len(args) > 2 {
				argNonce = args[2]
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			batchID, err := strconv.ParseUint(argBatchId, 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdateBatch(
				clientCtx.GetFromAddress().String(),
				batchID,
				argStatus,
			)
			var n uint64
			if argNonce != "" {
				n, err = strconv.ParseUint(argNonce, 10, 64)
				if err != nil {
					return err
				}
			}
			if n != 0 {
				msg.Nonce = n
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
