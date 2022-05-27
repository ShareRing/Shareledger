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
		Use:   "in [src_address] [dest_address] [src_network] [txHash] [amount] [fee]",
		Short: "Broadcast message in, to create the swap in request",
		Long: `
			[dest_address] should be shareledger address in shareledger
			[txHash] transaction hash from src network.
		`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSrcAddress := args[0]
			argDesAddress := args[1]
			argSrcNetwork := args[2]
			argHash := args[3]
			argAmount := args[4]
			argFee := args[5]

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

			msg := types.NewMsgRequestIn(
				clientCtx.GetFromAddress().String(),
				argSrcAddress,
				argDesAddress,
				argSrcNetwork,
				argHash,
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
