package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/electoral/types"
)

var _ = strconv.Itoa(0)

func CmdEnrollSwapManagers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-swap-managers [addresses]",
		Short: "Broadcast message enroll-swap-manager address_1 address_2",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnrollSwapManagers(
				clientCtx.GetFromAddress().String(),
				args[:],
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
