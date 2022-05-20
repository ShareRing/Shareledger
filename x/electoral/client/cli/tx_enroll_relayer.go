package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/electoral/types"
	"github.com/spf13/cobra"
	"strconv"
)

var _ = strconv.Itoa(0)

func CmdEnrollRelayers() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enroll-relayer [addresses]",
		Short:   "Broadcast message enroll-relayer",
		Example: "enroll_r-eplayer shareledger1mmaefaaxyfd2rctt9cev52xlda77df49pdyquk shareledger1fja6aazgvw6zfrh59xjc6w0jdpfhdkharz72lr",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnrollRelayers(
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
