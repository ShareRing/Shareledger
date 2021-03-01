package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/booking/types"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/myutils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const (
	bookFee     = 0.05
	completeFee = 0.03
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	bookCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "booking transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	bookCmd.AddCommand(flags.PostCommands(
		GetCmdBook(cdc),
		GetCmdComplete(cdc),
	)...)

	return bookCmd
}

func GetCmdBook(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "book [uuid] [duration]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			keySeed := viper.GetString(myutils.FlagKeySeed)
			seed, err := myutils.GetKeeySeedFromFile(keySeed)
			if err != nil {
				return err
			}

			cliCtx, txBldr, err := myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			if err != nil {
				return err
			}

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, bookFee)
			if err != nil {
				return err
			}
			txBldr = txBldr.WithFees(txFee)
			uuid := args[0]
			duration, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgBook(cliCtx.GetFromAddress(), uuid, int64(duration))
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdComplete(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "complete [bookID]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			keySeed := viper.GetString(myutils.FlagKeySeed)
			seed, err := myutils.GetKeeySeedFromFile(keySeed)
			if err != nil {
				return err
			}

			cliCtx, txBldr, err := myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			if err != nil {
				return err
			}

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, completeFee)
			if err != nil {
				return err
			}
			txBldr = txBldr.WithFees(txFee)
			bookID := args[0]
			msg := types.NewMsgComplete(cliCtx.GetFromAddress(), bookID)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}
