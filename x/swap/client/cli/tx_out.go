package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	denom "github.com/sharering/shareledger/x/utils/demo"
	"github.com/spf13/cobra"
)

func CmdOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out [dest_address] [network] [amount]",
		Short: "create new request swap out. Swap fee will be calculated and included to this transaction automatically",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDestAddr := args[0]
			argNetwork := args[1]
			argAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseDecCoin(argAmount)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			if err != nil {
				return err
			}
			params := &types.QueryGetSignSchemaRequest{
				Network: argNetwork,
			}

			res, err := queryClient.SignSchema(context.Background(), params)
			if err != nil {
				return err
			}
			if res.Schema.Fee == nil || res.Schema.Fee.Out == nil {
				return sdkerrors.Wrapf(sdkerrors.ErrLogic, "fee config was empty")
			}
			networkFee := denom.ToDisplayCoins(sdk.NewCoins(*res.Schema.Fee.Out))

			if err != nil {
				return err
			}
			msg := types.NewMsgRequestOut(
				clientCtx.GetFromAddress().String(),
				argDestAddr,
				argNetwork,
				amount,
				networkFee[0],
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
