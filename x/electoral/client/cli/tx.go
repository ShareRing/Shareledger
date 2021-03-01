package cli

import (
	"bufio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"bitbucket.org/shareringvietnam/shareledger-fix/x/electoral/types"
	"bitbucket.org/shareringvietnam/shareledger-fix/x/myutils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

const (
	minFeeShr = "1shr"
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
		GetCmdEnroll(cdc),
		GetCmdRevoke(cdc),
	)...)

	return bookCmd
}

func GetCmdEnroll(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll [address]",
		Short: "enroll a voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

			// Get key from key seed
			keySeed := viper.GetString(myutils.FlagKeySeed)
			if len(keySeed) > 0 {
				seed, err := myutils.GetKeeySeedFromFile(keySeed)
				if err != nil {
					return err
				}

				cliCtx, txBldr, err = myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
				if err != nil {
					return err
				}
			} else {
				// Get key from keychain
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			voter, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgEnrollVoter(cliCtx.GetFromAddress(), voter)
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

func GetCmdRevoke(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [address]",
		Short: "revoke a voter",
		Args:  cobra.ExactArgs(1),
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

			txBldr = txBldr.WithFees(minFeeShr)

			voter, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgRevokeVoter(cliCtx.GetFromAddress(), voter)
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
