package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/identity/types"

	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sharering/shareledger/x/myutils"
)

const (
	minShrFee = "1shr"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	identityTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "identity transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	identityTxCmd.AddCommand(flags.PostCommands(
		GetCmdEnrollIdSigners(cdc),
		GetCmdEnrollIdSignersFromFile(cdc),
		GetCmdRevokeIdSigners(cdc),
		GetCmdCreateId(cdc),
		GetCmdUpdateId(cdc),
		GetCmdDeleteId(cdc),
	)...)

	return identityTxCmd
}

func GetCmdEnrollIdSigners(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-signer [address1] [address2] [address3]",
		Short: "enroll id signers",
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

			txBldr = txBldr.WithFees(minShrFee)
			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			var signerAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				signerAddresses = append(signerAddresses, addr)
			}
			msg := types.NewMsgEnrollIDSigners(cliCtx.GetFromAddress(), signerAddresses)
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

func GetCmdEnrollIdSignersFromFile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-signers-from-file [filepath]",
		Short: "enroll id signers",
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

			txBldr = txBldr.WithFees(minShrFee)
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

				msg := types.NewMsgEnrollIDSigners(cliCtx.GetFromAddress(), loaderAddresses)
				err = msg.ValidateBasic()
				if err != nil {
					return err
				}
				if err := utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}); err != nil {
					return err
				}
				fmt.Printf("================ ENROLLED %d SIGNERS", n)
			}
			fmt.Printf("\nDONE ENROLL %d ID SIGNERS\n", listLength)
			return nil
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdRevokeIdSigners(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-signer [address1] [address2] [address3]",
		Short: "revoke id signers",
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

			txBldr = txBldr.WithFees(minShrFee)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			var signerAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				signerAddresses = append(signerAddresses, addr)
			}
			msg := types.NewMsgRevokeIDSigners(cliCtx.GetFromAddress(), signerAddresses)
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

func GetCmdCreateId(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [owner] [hash]",
		Short: "create id",
		Args:  cobra.ExactArgs(2),
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

			txBldr = txBldr.WithFees(minShrFee)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateId(cliCtx.GetFromAddress(), addr, args[1])
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

func GetCmdUpdateId(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [owner] [hash]",
		Short: "update id",
		Args:  cobra.ExactArgs(2),
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

			txBldr = txBldr.WithFees(minShrFee)
			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdateId(cliCtx.GetFromAddress(), addr, args[1])
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

func GetCmdDeleteId(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [owner]",
		Short: "delete id",
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

			txBldr = txBldr.WithFees(minShrFee)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty signer address list")
			}
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgDeleteId(cliCtx.GetFromAddress(), addr)
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
