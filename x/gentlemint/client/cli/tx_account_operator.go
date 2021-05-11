package cli

import (
	"bufio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	shareringUtils "github.com/ShareRing/modules/utils"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sharering/shareledger/x/gentlemint/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetCmdEnrollAccountOperators(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enroll-account-operator [address1] [address2] [address3]",
		Short: "enroll account operator",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

			keySeed := viper.GetString(shareringUtils.FlagKeySeed)
			if len(keySeed) > 0 {
				seed, err := shareringUtils.GetKeeySeedFromFile(keySeed)
				if err != nil {
					return err
				}

				cliCtx, txBldr, err = shareringUtils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
				if err != nil {
					return err
				}
			} else {
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}

			txBldr = txBldr.WithFees(minFeeShr)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty address list")
			}

			var accs []sdk.AccAddress

			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				accs = append(accs, addr)
			}
			msg := types.NewMsgEnrollAccOperators(cliCtx.GetFromAddress(), accs)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(shareringUtils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}

func GetCmdRevokeAccountOperator(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-account-operator [address1] [address2] [address3]",
		Short: "revoke account operator",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())

			var cliCtx context.CLIContext
			var txBldr auth.TxBuilder

			keySeed := viper.GetString(shareringUtils.FlagKeySeed)
			if len(keySeed) > 0 {
				seed, err := shareringUtils.GetKeeySeedFromFile(keySeed)
				if err != nil {
					return err
				}

				cliCtx, txBldr, err = shareringUtils.GetTxBldrAndCtxFromSeed(inBuf, cdc, seed)
				if err != nil {
					return err
				}
			} else {
				txBldr = auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
				cliCtx = context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			}
			txBldr = txBldr.WithFees(minFeeShr)

			if len(args) == 0 {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Empty address list")
			}
			var accs []sdk.AccAddress
			for _, a := range args {
				addr, err := sdk.AccAddressFromBech32(a)
				if err != nil {
					return err
				}
				accs = append(accs, addr)
			}
			msg := types.NewMsgRevokeAccOperators(cliCtx.GetFromAddress(), accs)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(shareringUtils.FlagKeySeed, "", "path to key_seed.json")
	return cmd
}
