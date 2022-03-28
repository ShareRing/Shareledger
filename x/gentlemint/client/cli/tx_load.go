package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/sharering/shareledger/x/gentlemint/types"
)

var _ = strconv.Itoa(0)

func CmdLoad() *cobra.Command {
	cmd := &cobra.Command{
		Use: "load [address] [coins]",
		Short: "load [coins] into [address]." +
			"coins: Expected format: {amount0}{denomination},...,{amountN}{denominationN}",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			decCoins, err := sdk.ParseDecCoins(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgLoad(
				clientCtx.GetFromAddress().String(),
				args[0],
				decCoins,
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
