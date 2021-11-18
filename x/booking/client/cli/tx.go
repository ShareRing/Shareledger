package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ShareRing/Shareledger/x/booking/types"
	myutils "github.com/ShareRing/Shareledger/x/utils"
)

const (
	bookFee     = 0.05
	completeFee = 0.03
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdBook(),
		GetCmdComplete(),
	)

	return cmd
}

// TODO: implement fee
func GetCmdBook() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "book [uuid] [duration]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// keySeed := viper.GetString(myutils.FlagKeySeed)
			// seed, err := myutils.GetKeeySeedFromFile(keySeed)
			// if err != nil {
			// 	return err
			// }

			// cliCtx, txBldr, err := myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// if err != nil {
			// 	return err
			// }

			// txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, bookFee)
			// if err != nil {
			// 	return err
			// }
			// txBldr = txBldr.WithFees(txFee)
			// uuid := args[0]
			// duration, err := strconv.Atoi(args[1])
			// if err != nil {
			// 	return err
			// }
			// msg := types.NewMsgBook(cliCtx.GetFromAddress(), uuid, int64(duration))
			// err = msg.ValidateBasic()
			// if err != nil {
			// 	return err
			// }

			// return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// seed implementation
			keySeed := viper.GetString(myutils.FlagKeySeed)
			if keySeed != "" {
				clientCtx, err = myutils.CreateContextFromSeed(keySeed, clientCtx)
				if err != nil {
					return err
				}
			}

			uuid := args[0]
			duration, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgBook(clientCtx.GetFromAddress().String(), uuid, int64(duration))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)

		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")

	return cmd
}

func GetCmdComplete() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "complete [bookID]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// keySeed := viper.GetString(myutils.FlagKeySeed)
			// seed, err := myutils.GetKeeySeedFromFile(keySeed)
			// if err != nil {
			// 	return err
			// }

			// cliCtx, txBldr, err := myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// if err != nil {
			// 	return err
			// }

			// txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, completeFee)
			// if err != nil {
			// 	return err
			// }
			// txBldr = txBldr.WithFees(txFee)
			// bookID := args[0]
			// msg := types.NewMsgComplete(cliCtx.GetFromAddress(), bookID)
			// err = msg.ValidateBasic()
			// if err != nil {
			// 	return err
			// }

			// return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// seed implementation
			keySeed := viper.GetString(myutils.FlagKeySeed)
			if keySeed != "" {
				clientCtx, err = myutils.CreateContextFromSeed(keySeed, clientCtx)
				if err != nil {
					return err
				}
			}

			bookID := args[0]
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgComplete(clientCtx.GetFromAddress().String(), bookID)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")

	return cmd
}
