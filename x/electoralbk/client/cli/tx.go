package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ShareRing/Shareledger/x/electoralbk/types"
	myutils "github.com/ShareRing/Shareledger/x/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

// var (
// 	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
// )

// const (
// 	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
// )

// const (
// 	minFeeShr = "1shr"
// )

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
		GetCmdEnroll(),
		GetCmdRevoke(),
	)

	return cmd
}

// TODO: fee
func GetCmdEnroll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll [address]",
		Short: "enroll a voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// var cliCtx context.CLIContext
			// var txBldr auth.TxBuilder

			// // Get key from key seed
			// keySeed := viper.GetString(myutils.FlagKeySeed)
			// if len(keySeed) > 0 {
			// 	seed, err := myutils.GetKeeySeedFromFile(keySeed)
			// 	if err != nil {
			// 		return err
			// 	}

			// 	cliCtx, txBldr, err = myutils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// 	if err != nil {
			// 		return err
			// 	}
			// } else {
			// 	// Get key from keychain
			// 	txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			// 	cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			// }

			// txBldr = txBldr.WithFees(minFeeShr)

			// voter, err := sdk.AccAddressFromBech32(args[0])
			// if err != nil {
			// 	return err
			// }
			// msg := types.NewMsgEnrollVoter(cliCtx.GetFromAddress(), voter)
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

			voter := args[0]

			msg := types.NewMsgEnrollVoter(clientCtx.GetFromAddress().String(), voter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}

func GetCmdRevoke() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [address]",
		Short: "revoke a voter",
		Args:  cobra.ExactArgs(1),
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

			// txBldr = txBldr.WithFees(minFeeShr)

			// voter, err := sdk.AccAddressFromBech32(args[0])
			// if err != nil {
			// 	return err
			// }
			// msg := types.NewMsgRevokeVoter(cliCtx.GetFromAddress(), voter)
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

			voter := args[0]

			msg := types.NewMsgRevokeVoter(clientCtx.GetFromAddress().String(), voter)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}
