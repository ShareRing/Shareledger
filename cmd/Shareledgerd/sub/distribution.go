package sub

import (
	"fmt"
	"strings"

	myutils "github.com/sharering/shareledger/x/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	distributioncli "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewWithdrawRewardsCmd() *cobra.Command {
	bech32PrefixValAddr := sdk.GetConfig().GetBech32ValidatorAddrPrefix()

	cmd := &cobra.Command{
		Use:   "withdraw-rewards [validator-addr]",
		Short: "Withdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Withdraw rewards from a given delegation address,
and optionally withdraw validator commission if the delegation address given is a validator operator.

Example:
$ %s distribution withdraw-rewards %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj --key-seed mynode_key_seed.json
$ %s distribution withdraw-rewards %s1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj --key-seed mynode_key_seed.json --commission
`,
				version.AppName, bech32PrefixValAddr, version.AppName, bech32PrefixValAddr,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			keySeed := viper.GetString(myutils.FlagKeySeed)
			if keySeed != "" {
				clientCtx, err = myutils.CreateContextFromSeed(keySeed, clientCtx)
				if err != nil {
					return err
				}
			}

			delAddr := clientCtx.GetFromAddress()
			valAddr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msgs := []sdk.Msg{types.NewMsgWithdrawDelegatorReward(delAddr, valAddr)}

			if commission, _ := cmd.Flags().GetBool(distributioncli.FlagCommission); commission {
				msgs = append(msgs, types.NewMsgWithdrawValidatorCommission(valAddr))
			}

			for _, msg := range msgs {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	cmd.Flags().Bool(distributioncli.FlagCommission, false, "Withdraw the validator's commission in addition to the rewards")
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewWithdrawAllRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-all-rewards",
		Short: "withdraw all delegations rewards for a delegator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Withdraw all rewards for a single delegator.
Note that if you use this command with --%[2]s=%[3]s or --%[2]s=%[4]s, the %[5]s flag will automatically be set to 0.

Example:
$ %[1]s distribution withdraw-all-rewards --key-seed mynode_key_seed.json
`,
				version.AppName, flags.FlagBroadcastMode, flags.BroadcastSync, flags.BroadcastAsync, distributioncli.FlagMaxMessagesPerTx,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			keySeed := viper.GetString(myutils.FlagKeySeed)
			if keySeed != "" {
				clientCtx, err = myutils.CreateContextFromSeed(keySeed, clientCtx)
				if err != nil {
					return err
				}
			}

			delAddr := clientCtx.GetFromAddress()

			// The transaction cannot be generated offline since it requires a query
			// to get all the validators.
			if clientCtx.Offline {
				return fmt.Errorf("cannot generate tx in offline mode")
			}

			queryClient := types.NewQueryClient(clientCtx)
			delValsRes, err := queryClient.DelegatorValidators(cmd.Context(), &types.QueryDelegatorValidatorsRequest{DelegatorAddress: delAddr.String()})
			if err != nil {
				return err
			}

			validators := delValsRes.Validators
			// build multi-message transaction
			msgs := make([]sdk.Msg, 0, len(validators))
			for _, valAddr := range validators {
				val, err := sdk.ValAddressFromBech32(valAddr)
				if err != nil {
					return err
				}

				msg := types.NewMsgWithdrawDelegatorReward(delAddr, val)
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
				msgs = append(msgs, msg)
			}

			chunkSize, _ := cmd.Flags().GetInt(distributioncli.FlagMaxMessagesPerTx)
			if clientCtx.BroadcastMode != flags.BroadcastBlock && chunkSize > 0 {
				return fmt.Errorf("cannot use broadcast mode %[1]s with %[2]s != 0",
					clientCtx.BroadcastMode, distributioncli.FlagMaxMessagesPerTx)
			}

			return newSplitAndApply(tx.GenerateOrBroadcastTxCLI, clientCtx, cmd.Flags(), msgs, chunkSize)
		},
	}

	cmd.Flags().Int(distributioncli.FlagMaxMessagesPerTx, distributioncli.MaxMessagesPerTxDefault, "Limit the number of messages per tx (0 for unlimited)")
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
