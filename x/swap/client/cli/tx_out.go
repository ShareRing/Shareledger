package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out [dest_address] [network] [amount] [fee]",
		Short: "create new request swap out",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDestAddr := args[0]
			argNetwork := args[1]
			argAmount := args[2]
			argFee := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseDecCoin(argAmount)
			if err != nil {
				return err
			}
			fee, err := sdk.ParseDecCoin(argFee)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwapOut(
				clientCtx.GetFromAddress().String(),
				argDestAddr,
				argNetwork,
				amount,
				fee,
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
