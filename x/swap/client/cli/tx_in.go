package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"github.com/sharering/shareledger/x/swap/types"
)

func CmdIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in [dest_address] [src_network] [txHashes] [amount] [fee]",
		Short: "Broadcast message in, to create the swap in request",
		Long: `
			[dest_address] should be shareledger address in shareledger
			[txHashes] <hash1>:<sender>:<logIndex>,<hash2>:<sender>:<logIndex>.... : tx Hashes list is required.
			[amount] the total of all external transactions' amount minus swap fee.
			[fee] the fee for swap in which is configured in schema data.
		`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDesAddress := args[0]
			argSrcNetwork := args[1]
			hashesLog := strings.Split(args[2], ",")
			txHashes := make([]*types.TxEvent, 0, len(hashesLog))

			for i := range hashesLog {
				hl := strings.TrimSpace(hashesLog[i])

				newHL := strings.Split(hl, ":")

				if len(newHL) > 0 {
					logIndex, err := strconv.ParseUint(newHL[2], 10, 64)
					if err != nil {
						return err
					}
					txHashes = append(txHashes, &types.TxEvent{
						TxHash:   newHL[0],
						Sender:   newHL[1],
						LogIndex: logIndex,
					})
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
