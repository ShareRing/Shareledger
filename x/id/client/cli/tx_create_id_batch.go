package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/ShareRing/Shareledger/x/id/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdCreateIdBatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-id-batch [backup-address] [extra-data] [id] [owner-address]",
		Short: "Broadcast message CreateIdBatch",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBackupAddress := args[0]
			argExtraData := args[1]
			argId := args[2]
			argOwnerAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIdBatch(
				clientCtx.GetFromAddress().String(),
				argBackupAddress,
				argExtraData,
				argId,
				argOwnerAddress,
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
