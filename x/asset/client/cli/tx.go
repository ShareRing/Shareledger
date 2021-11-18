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

	"github.com/ShareRing/Shareledger/x/asset/types"
	myutils "github.com/ShareRing/Shareledger/x/utils"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

const (
	creationFee = 0.05
	updateFee   = 0.03
	deleteFee   = 0.01
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
		CmdCreateAsset(),
		CmdUpdateAsset(),
		CmdDeleteAsset(),
	)

	return cmd
}

// TODO: implement fee
func CmdCreateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create [hash] [uuid] [status] [fee]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// keySeed := viper.GetString(utils.FlagKeySeed)
			// seed, err := utils.GetKeeySeedFromFile(keySeed)
			// if err != nil {
			// 	return err
			// }

			// cliCtx, txBldr, err := utils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// if err != nil {
			// 	return err
			// }

			// txFee, err := utils.GetFeeFromShrp(cdc, cliCtx, creationFee)
			// if err != nil {
			// 	return err
			// }
			// txBldr = txBldr.WithFees(txFee)

			// hash := []byte(args[0])
			// uuid := args[1]
			// status, err := strconv.ParseBool(args[2])
			// if err != nil {
			// 	return err
			// }
			// fee, err := strconv.Atoi(args[3])
			// if err != nil {
			// 	return err
			// }
			// msg := types.NewMsgCreate(cliCtx.GetFromAddress(), hash, uuid, status, int64(fee))
			// err = msg.ValidateBasic()
			// if err != nil {
			// 	return err
			// }

			// return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})

			//------------------------------
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

			hash := []byte(args[0])
			uuid := args[1]
			status, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			rate, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreate(clientCtx.GetFromAddress().String(), hash, uuid, status, int64(rate))
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

func CmdUpdateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update [hash] [uuid] [status] [fee]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// keySeed := viper.GetString(utils.FlagKeySeed)
			// seed, err := utils.GetKeeySeedFromFile(keySeed)
			// if err != nil {
			// 	return err
			// }

			// cliCtx, txBldr, err := utils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// if err != nil {
			// 	return err
			// }

			// txFee, err := utils.GetFeeFromShrp(cdc, cliCtx, updateFee)

			// if err != nil {
			// 	return err
			// }

			// txBldr = txBldr.WithFees(txFee)

			// hash := []byte(args[0])
			// uuid := args[1]
			// status, err := strconv.ParseBool(args[2])
			// if err != nil {
			// 	return err
			// }
			// fee, err := strconv.Atoi(args[3])
			// if err != nil {
			// 	return err
			// }
			// msg := types.NewMsgUpdate(cliCtx.GetFromAddress(), hash, uuid, status, int64(fee))
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

			hash := []byte(args[0])
			uuid := args[1]
			status, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			rate, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdate(clientCtx.GetFromAddress().String(), hash, uuid, status, int64(rate))
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

func CmdDeleteAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete [uuid]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// inBuf := bufio.NewReader(cmd.InOrStdin())

			// keySeed := viper.GetString(client.FlagKeySeed)
			// seed, err := client.GetKeeySeedFromFile(keySeed)
			// if err != nil {
			// 	return err
			// }

			// cliCtx, txBldr, err := client.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
			// if err != nil {
			// 	return err
			// }

			// txFee, err := client.GetFeeFromShrp(cdc, cliCtx, deleteFee)
			// if err != nil {
			// 	return err
			// }
			// txBldr = txBldr.WithFees(txFee)
			// uuid := args[0]
			// msg := types.NewMsgDelete(cliCtx.GetFromAddress(), uuid)
			// err = msg.ValidateBasic()
			// if err != nil {
			// 	return err
			// }
			// return client.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{&msg})

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
			msg := types.NewMsgDelete(clientCtx.GetFromAddress().String(), uuid)
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
