package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	myutils "github.com/ShareRing/modules/utils"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

var (
	minFeeShr = myutils.MINFEE.String() + "shr"
	sendFee   = myutils.LOWFEE
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	gentlemintTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "gentlemint transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	gentlemintTxCmd.AddCommand(flags.PostCommands(
		GetCmdLoadSHR(cdc),
		GetCmdBuySHR(cdc),
		GetCmdSetExchange(cdc),
		GetCmdEnrollSHRPLoader(cdc),
		GetCmdEnrollSHRPLoaderFromFile(cdc),
		GetCmdRevokeSHRPLoader(cdc),
		GetCmdRevokeSHRPLoaderFromFile(cdc),
		GetCmdLoadSHRP(cdc),
		GetCmdSendSHRP(cdc),
		GetCmdSendSHR(cdc),
		GetCmdBurnSHRP(cdc),
		GetCmdBurnSHR(cdc),
		GetCmdEnrollIdSigners(cdc),
		GetCmdEnrollIdSignersFromFile(cdc),
		GetCmdRevokeIdSigners(cdc),
		GetCmdEnrollAccountOperators(cdc),
		GetCmdRevokeAccountOperator(cdc),
		GetCmdEnrollDocIssuer(cdc),
		GetCmdRevokeDocIssuer(cdc),
	)...)

	return gentlemintTxCmd
}

func GetCmdLoadSHR(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-shr [address] [amount]",
		Short: "bid for existing name or claim new name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			// amt, err := strconv.Atoi(args[1])
			// if err != nil {
			// 	return err
			// }
			amt := args[1]
			msg := types.NewMsgLoadSHR(cliCtx.GetFromAddress(), to, amt)
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

func GetCmdLoadSHRP(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load-shrp [address] [amount]",
		Short: "bid for existing name or claim new name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			amt := args[1]
			msg := types.NewMsgLoadSHRP(cliCtx.GetFromAddress(), to, amt)
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

func GetCmdSendSHRP(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-shrp [address] [amount]",
		Short: "send shrp to another address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, sendFee)

			txBldr = txBldr.WithFees(txFee)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			amt := args[1]
			msg := types.NewMsgSendSHRP(cliCtx.GetFromAddress(), to, amt)
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

func GetCmdSendSHR(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-shr [address] [amount]",
		Short: "send shr to another address",
		Long:  "buy shr using shrp balance if the sender doesn't have enough shr",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txFee, err := myutils.GetFeeFromShrp(cdc, cliCtx, sendFee)
			if err != nil {
				return err
			}

			txBldr = txBldr.WithFees(txFee)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			// amt, err := strconv.Atoi(args[1])
			// if err != nil {
			// 	return err
			// }
			amt := args[1]
			msg := types.NewMsgSendSHR(cliCtx.GetFromAddress(), to, amt)
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

func GetCmdBuySHR(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-shr [amount]",
		Short: "buy shr from shrp",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			amt := args[0]
			msg := types.NewMsgBuySHR(cliCtx.GetFromAddress(), amt)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdSetExchange(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange [rate]",
		Short: "set exchange rate of shrp to shr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			msg := types.NewMsgSetExchange(cliCtx.GetFromAddress(), args[0])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdBurnSHRP(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-shrp [amount]",
		Short: "bid for existing name or claim new name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			amt := args[0]
			msg := types.NewMsgBurnSHRP(cliCtx.GetFromAddress(), amt)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdBurnSHR(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-shr [amount]",
		Short: "burn shr from treasurer account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			// amt, err := strconv.Atoi(args[0])
			// if err != nil {
			// 	return err
			// }
			amt := args[0]
			msg := types.NewMsgBurnSHR(cliCtx.GetFromAddress(), amt)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdEnrollSHRPLoader(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-loaders [address 1] [address 2] [address n]",
		Short: "bid for existing name or claim new name",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			var loaderAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				loaderAddresses = append(loaderAddresses, addr)
			}

			msg := types.NewMsgEnrollSHRPLoaders(cliCtx.GetFromAddress(), loaderAddresses)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdEnrollSHRPLoaderFromFile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-loaders-from-file [filepath]",
		Short: "bid for existing name or claim new name",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			addrList, err := myutils.GetAddressFromFile(args[0])
			if err != nil {
				return err
			}

			listLength := len(addrList)
			nloop := listLength / 5
			if nloop*5 > listLength {
				nloop = nloop + 1
			}
			for i := 0; i < nloop; i++ {
				var loaderAddresses []sdk.AccAddress
				m := 5 * i
				n := 5*i + 5
				if n > listLength {
					n = listLength
				}
				for _, a := range addrList[m:n] {
					addr, err := sdk.AccAddressFromBech32(a)
					if err != nil {
						return err
					}
					loaderAddresses = append(loaderAddresses, addr)
				}

				msg := types.NewMsgEnrollSHRPLoaders(cliCtx.GetFromAddress(), loaderAddresses)
				err = msg.ValidateBasic()
				if err != nil {
					return err
				}
				if err := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}); err != nil {
					return err
				}
				fmt.Printf("================ ENROLLED %d LOADERS", n)
			}
			fmt.Printf("\nDONE ENROLL %d SHRP LOADERS\n", listLength)
			return nil
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdRevokeSHRPLoader(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-loaders [address 1] [address 2] [address n]",
		Short: "revoke shrp loaders",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			var loaderAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				loaderAddresses = append(loaderAddresses, addr)
			}

			msg := types.NewMsgEnrollSHRPLoaders(cliCtx.GetFromAddress(), loaderAddresses)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdRevokeSHRPLoaderFromFile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "revoke-loaders-from-file [filepath]",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

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
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			addrList, err := myutils.GetAddressFromFile(args[0])
			if err != nil {
				return err
			}

			listLength := len(addrList)
			nloop := listLength / 5
			if nloop*5 > listLength {
				nloop = nloop + 1
			}
			for i := 0; i < nloop; i++ {
				var loaderAddresses []sdk.AccAddress
				m := 5 * i
				n := 5*i + 5
				if n > listLength {
					n = listLength
				}
				for _, a := range addrList[m:n] {
					addr, err := sdk.AccAddressFromBech32(a)
					if err != nil {
						return err
					}
					loaderAddresses = append(loaderAddresses, addr)
				}

				msg := types.NewMsgRevokeSHRPLoaders(cliCtx.GetFromAddress(), loaderAddresses)
				err = msg.ValidateBasic()
				if err != nil {
					return err
				}
				if err := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}); err != nil {
					return err
				}
				fmt.Printf("================ REVOKED %d LOADERS", n)
			}
			fmt.Printf("\nDONE REVOKE %d SHRP LOADERS\n", listLength)
			return nil
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}
