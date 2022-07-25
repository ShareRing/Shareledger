package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/swap/types"
	"github.com/spf13/cobra"
	"strings"
)

func CmdIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in [dest_address] [src_network] [txHashes] [amount] [fee]",
		Short: "Broadcast message in, to create the swap in request",
		Long: `
			[dest_address] should be shareledger address in shareledger
			[txHashes] <hash1>,<hash2>.... : tx Hashes list is required.
			[amount] the total of all external transactions' amount minus swap fee.
			[fee] the fee for swap in which is configured in schema data.
		`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDesAddress := args[0]
			argSrcNetwork := args[1]
			hashes := strings.Split(args[2], ",")
			txHashes := make([]string, 0, len(hashes))

			for i := range hashes {
				h := strings.TrimSpace(hashes[i])
				if h != "" {
					txHashes = append(txHashes, h)
				}
			}

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

			msg := types.NewMsgRequestIn(
				clientCtx.GetFromAddress().String(),
				argDesAddress,
				argSrcNetwork,
				txHashes,
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
