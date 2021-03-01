package main

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/sharering/shareledger/x/myutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagCommission = "commission"
)

func GetCmdWithdrawRewards(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-rewards [val-addr]",
		Short: "withdraw rewards for a delegator",
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
			delAddr := cliCtx.GetFromAddress()
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgWithdrawDelegatorReward(delAddr, valAddr)}
			if viper.GetBool(flagCommission) {
				msgs = append(msgs, types.NewMsgWithdrawValidatorCommission(valAddr))
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, msgs)
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd.Flags().Bool(flagCommission, false, "also withdraw validator's commission")
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}
