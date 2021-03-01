package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/sharering/shareledger/x/myutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	minFeeShr = "1shr"
)

func GetCmdCreateValidator(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "create new validator initialized with a self-delegation to it",
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
			txBldr, msg, err := BuildCreateValidatorMsg(cliCtx, txBldr)
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(cli.FlagAmount, "", "Amount of coins to bond")
	cmd.Flags().String(cli.FlagPubKey, "", "The bech32 encoded PubKey of the validator")
	cmd.Flags().String(cli.FlagMoniker, "monica", "The validator's name")
	cmd.Flags().String(cli.FlagIdentity, "", "The optional identity signature (ex. UPort or Keybase)")
	cmd.Flags().String(cli.FlagWebsite, "", "The validator's (optional) website")
	cmd.Flags().String(cli.FlagSecurityContact, "", "The validator's (optional) security contact email")
	cmd.Flags().String(cli.FlagDetails, "", "The validator's (optional) details")
	cmd.Flags().String(cli.FlagCommissionRate, "", "The new commission rate percentage")
	cmd.Flags().String(cli.FlagCommissionMaxRate, "", "The maximum commission rate percentage")
	cmd.Flags().String(cli.FlagCommissionMaxChangeRate, "", "The maximum commission change rate percentage (per day)")
	cmd.Flags().String(cli.FlagMinSelfDelegation, "", "The minimum self delegation required on the validator")
	cmd.Flags().String(cli.FlagNodeID, "", "NodeID")
	cmd.Flags().String(cli.FlagIP, "", "Node ip address")
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd.MarkFlagRequired(cli.FlagAmount)
	cmd.MarkFlagRequired(cli.FlagPubKey)
	cmd.MarkFlagRequired(cli.FlagMoniker)
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(cliCtx context.CLIContext, txBldr auth.TxBuilder) (auth.TxBuilder, sdk.Msg, error) {

	amounstStr := viper.GetString(cli.FlagAmount)
	amount, err := sdk.ParseCoin(amounstStr)
	if err != nil {
		return txBldr, nil, err
	}

	valAddr := cliCtx.GetFromAddress()
	pkStr := viper.GetString(cli.FlagPubKey)

	pk, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, pkStr)
	if err != nil {
		return txBldr, nil, err
	}

	description := types.NewDescription(
		viper.GetString(cli.FlagMoniker),
		viper.GetString(cli.FlagIdentity),
		viper.GetString(cli.FlagWebsite),
		viper.GetString(cli.FlagSecurityContact),
		viper.GetString(cli.FlagDetails),
	)

	// get the initial validator commission parameters
	rateStr := viper.GetString(cli.FlagCommissionRate)
	maxRateStr := viper.GetString(cli.FlagCommissionMaxRate)
	maxChangeRateStr := viper.GetString(cli.FlagCommissionMaxChangeRate)
	commissionRates, err := buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr)
	if err != nil {
		return txBldr, nil, err
	}

	// get the initial validator min self delegation
	msbStr := viper.GetString(cli.FlagMinSelfDelegation)
	minSelfDelegation, ok := sdk.NewIntFromString(msbStr)
	if !ok {
		return txBldr, nil, types.ErrMinSelfDelegationInvalid
	}

	msg := types.NewMsgCreateValidator(
		sdk.ValAddress(valAddr), pk, amount, description, commissionRates, minSelfDelegation,
	)

	if viper.GetBool(flags.FlagGenerateOnly) {
		ip := viper.GetString(cli.FlagIP)
		nodeID := viper.GetString(cli.FlagNodeID)
		if nodeID != "" && ip != "" {
			txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
		}
	}

	return txBldr, msg, nil
}

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission types.CommissionRates, err error) {
	if rateStr == "" || maxRateStr == "" || maxChangeRateStr == "" {
		return commission, errors.New("must specify all validator commission parameters")
	}

	rate, err := sdk.NewDecFromStr(rateStr)
	if err != nil {
		return commission, err
	}

	maxRate, err := sdk.NewDecFromStr(maxRateStr)
	if err != nil {
		return commission, err
	}

	maxChangeRate, err := sdk.NewDecFromStr(maxChangeRateStr)
	if err != nil {
		return commission, err
	}

	commission = types.NewCommissionRates(rate, maxRate, maxChangeRate)
	return commission, nil
}

func GetCmdEditValidator(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-validator",
		Short: "edit an existing validator account",
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

			valAddr := cliCtx.GetFromAddress()
			description := types.NewDescription(
				viper.GetString(cli.FlagMoniker),
				viper.GetString(cli.FlagIdentity),
				viper.GetString(cli.FlagWebsite),
				viper.GetString(cli.FlagSecurityContact),
				viper.GetString(cli.FlagDetails),
			)

			var newRate *sdk.Dec

			commissionRate := viper.GetString(cli.FlagCommissionRate)
			if commissionRate != "" {
				rate, err := sdk.NewDecFromStr(commissionRate)
				if err != nil {
					return fmt.Errorf("invalid new commission rate: %v", err)
				}

				newRate = &rate
			}

			var newMinSelfDelegation *sdk.Int

			minSelfDelegationString := viper.GetString(cli.FlagMinSelfDelegation)
			if minSelfDelegationString != "" {
				msb, ok := sdk.NewIntFromString(minSelfDelegationString)
				if !ok {
					return types.ErrMinSelfDelegationInvalid
				}

				newMinSelfDelegation = &msb
			}

			msg := types.NewMsgEditValidator(sdk.ValAddress(valAddr), description, newRate, newMinSelfDelegation)

			// build and sign the transaction, then broadcast to Tendermint
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(cli.FlagMoniker, types.DoNotModifyDesc, "The validator's name")
	cmd.Flags().String(cli.FlagIdentity, types.DoNotModifyDesc, "The (optional) identity signature (ex. UPort or Keybase)")
	cmd.Flags().String(cli.FlagWebsite, types.DoNotModifyDesc, "The validator's (optional) website")
	cmd.Flags().String(cli.FlagSecurityContact, types.DoNotModifyDesc, "The validator's (optional) security contact email")
	cmd.Flags().String(cli.FlagDetails, types.DoNotModifyDesc, "The validator's (optional) details")
	cmd.Flags().String(cli.FlagCommissionRate, "", "The new commission rate percentage")
	cmd.Flags().String(cli.FlagMinSelfDelegation, "", "The minimum self delegation required on the validator")
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}

func GetCmdDelegate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate [validator-addr] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Delegate liquid tokens to a validator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Delegate an amount of liquid coins to a validator from your wallet.

Example:
$ %s tx staking delegate cosmosvaloper1l2rsakp388kuv9k8qzq6lrm9taddae7fpx59wm 1000stake --key-seed my_key_seed.json
`,
				version.ClientName,
			),
		),
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

			amount, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			delAddr := cliCtx.GetFromAddress()
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgDelegate(delAddr, valAddr, amount)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}

func GetCmdUnbond(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unbond [validator-addr] [amount]",
		Short: "Unbond shares from a validator",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Unbond an amount of bonded shares from a validator.

Example:
$ %s tx staking unbond cosmosvaloper1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj 100stake --from mykey
`,
				version.ClientName,
			),
		),
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

			amount, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUndelegate(delAddr, valAddr, amount)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}

func GetCmdUnjail(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unjail [validator-addr]",
		Short: "unjail a validator",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Example:
$ %s tx staking unbond cosmosvaloper1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj --key-seed ./my_key_seed.json
`,
				version.ClientName,
			),
		),
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

			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := slashing.NewMsgUnjail(valAddr)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(myutils.FlagKeySeed, "", "path to key_seed.json file")
	cmd = flags.PostCommands(cmd)[0]
	return cmd
}
