package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var _ = strconv.Itoa(0)

func CmdSetExchange() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange [rate]",
		Short: "set exchange [rate] shrp to shr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRate := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetExchange(
				clientCtx.GetFromAddress().String(),
				argRate,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			msg.Rate = sdk.MustNewDecFromStr(msg.Rate).Mul(sdk.NewDec(denom.ShrExponent)).String()

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
