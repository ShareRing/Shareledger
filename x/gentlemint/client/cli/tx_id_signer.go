package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"

	myutils "github.com/ShareRing/modules/utils"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetCmdEnrollIdSigners(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-id-signer [address1] [address2] [address3]",
		Short: "enroll id signers",
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

			var signerAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				signerAddresses = append(signerAddresses, addr)
			}
			msg := types.NewMsgEnrollIDSigners(cliCtx.GetFromAddress(), signerAddresses)
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

func GetCmdEnrollIdSignersFromFile(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-id-signers-from-file [filepath]",
		Short: "enroll id signers",
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
		Use:   "revoke-id-signer [address1] [address2] [address3]",
		Short: "revoke id signers",
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
			var signerAddresses []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				signerAddresses = append(signerAddresses, addr)
			}
			msg := types.NewMsgRevokeIDSigners(cliCtx.GetFromAddress(), signerAddresses)
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
