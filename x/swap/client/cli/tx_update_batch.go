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
		Use:   "update-batch [batch-id] [status] [network]",
		Short: "Broadcast message update-batch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argBatchId := args[0]
			argStatus := args[1]
			argNetwork := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			batchID, _ := strconv.ParseUint(argBatchId, 10, 64)
			msg := types.NewMsgUpdateBatch(
				clientCtx.GetFromAddress().String(),
				batchID,
				argNetwork,
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
