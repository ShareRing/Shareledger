package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	myutils "github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sharering/shareledger/x/asset/types"
)

var (
	creationFee = myutils.HIGHFEE
	updateFee   = myutils.MEDIUMFEE
	deleteFee   = myutils.LOWFEE
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	assetCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "asset transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	assetCmd.AddCommand(flags.PostCommands(
		GetCmdCreateAsset(cdc),
		GetCmdUpdateAsset(cdc),
		GetCmdDeleteAsset(cdc),
	)...)

	return assetCmd
}

func GetCmdCreateAsset(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create [hash] [uuid] [status] [fee]",
		Args: cobra.ExactArgs(4),
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

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, creationFee)
			if err != nil {
				return err
			}
			txBldr = txBldr.WithFees(txFee)
			hash := []byte(args[0])
			uuid := args[1]
			status, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			fee, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			msg := types.NewMsgCreate(cliCtx.GetFromAddress(), hash, uuid, status, int64(fee))
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

func GetCmdUpdateAsset(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update [hash] [uuid] [status] [fee]",
		Args: cobra.ExactArgs(4),
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

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, updateFee)

			if err != nil {
				return err
			}

			txBldr = txBldr.WithFees(txFee)

			hash := []byte(args[0])
			uuid := args[1]
			status, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			fee, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdate(cliCtx.GetFromAddress(), hash, uuid, status, int64(fee))
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

func GetCmdDeleteAsset(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete [hash] [uuid] [status] [fee]",
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

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, deleteFee)
			if err != nil {
				return err
			}
			txBldr = txBldr.WithFees(txFee)
			uuid := args[0]
			msg := types.NewMsgDelete(cliCtx.GetFromAddress(), uuid)
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
