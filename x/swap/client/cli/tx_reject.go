package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdReject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject [txnIds]",
		Short: "Broadcast message reject the swapping request base on request_id \n [txnIds] format ID1,ID2,ID3",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			argIDS := strings.Split(args[0], ",")
			txIds := make([]uint64, 0, len(argIDS))
			for _, str := range argIDS {
				id, err := strconv.ParseUint(str, 10, 64)
				if err != nil {
					return err
				}
				txIds = append(txIds, id)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReject(
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
