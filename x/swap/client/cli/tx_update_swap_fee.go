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

func CmdUpdateSwapFee() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-swap-fee [network]",
		Short:   "update swap in/out's fee for given schema's [network]. Only support shr/nshr",
		Example: `update-swap-fee eth --fee-in 17shr --fee-out 10shr`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNetwork := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feeIn, feeOut, err := parseInOutFeeFromCmd(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdateSwapFee(
				clientCtx.GetFromAddress().String(),
				argNetwork,
				feeIn,
				feeOut,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(FlagsFeeIn, "", "swapping fee in")
	cmd.Flags().String(FlagsFeeOut, "", "swapping fee out")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
