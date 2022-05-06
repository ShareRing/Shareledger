package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
)

func CmdIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in [src_address] [dest_address] [src_network] [amount] [fee]",
		Short: "Broadcast message in",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSrcAddress := args[0]
			argDesAddress := args[1]
			argSrcNetwork := args[2]
			argAmount := args[3]
			argFee := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sAmount, err := sdk.ParseDecCoin(argAmount)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap amount %s err %", argAmount, err)
			}
			sFee, err := sdk.ParseDecCoin(argFee)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap amount %s err %", argFee, err)
			}

			msg := types.NewMsgSwapIn(
				clientCtx.GetFromAddress().String(),
				argSrcAddress,
				argDesAddress,
				argSrcNetwork,
				sAmount,
				sFee,
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
