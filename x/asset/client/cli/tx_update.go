package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/asset/types"
)

var _ = strconv.Itoa(0)

func CmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [hash] [uuid] [status] [rate]",
		Short: "Broadcast message UpdateAsset",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHash := []byte(args[0])
			argUUID := args[1]
			argStatus, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			argRate, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdate(
				clientCtx.GetFromAddress().String(),
				argHash,
				argUUID,
				argStatus,
				int64(argRate),
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
