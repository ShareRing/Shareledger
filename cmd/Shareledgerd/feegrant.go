package main

import (
	"fmt"
	"strings"
	"time"

	myutils "github.com/ShareRing/Shareledger/x/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantcli "github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmdFeeGrant returns a CLI command handler for creating a MsgGrantAllowance transaction.
func NewCmdFeeGrant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [grantee] --key-seed granter_key_seed.json",
		Short: "Grant Fee allowance to an address",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				`Grant authorization to pay fees from your address. Note, the'--from' flag is
				ignored as it is implied from [granter].

Examples:
%s %s grant shareledger1skjw... --key-seed granter_key_seed.json --spend-limit 100stake --expiration 2022-01-30T15:04:05Z or
%s %s grant shareledger1skjw... --key-seed granter_key_seed.json --spend-limit 100stake --period 3600 --period-limit 10stake --expiration 36000 or
%s %s grant shareledger1skjw... --key-seed granter_key_seed.json --spend-limit 100stake --expiration 2022-01-30T15:04:05Z 
	--allowed-messages "/cosmos.gov.v1beta1.MsgSubmitProposal,/cosmos.gov.v1beta1.MsgVote"
				`, version.AppName, feegrant.ModuleName, version.AppName, feegrant.ModuleName, version.AppName, feegrant.ModuleName,
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

			grantee, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			granter := clientCtx.GetFromAddress()
			sl, err := cmd.Flags().GetString(feegrantcli.FlagSpendLimit)
			if err != nil {
				return err
			}

			// if `FlagSpendLimit` isn't set, limit will be nil
			limit, err := sdk.ParseCoinsNormalized(sl)
			if err != nil {
				return err
			}

			exp, err := cmd.Flags().GetString(feegrantcli.FlagExpiration)
			if err != nil {
				return err
			}

			basic := feegrant.BasicAllowance{
				SpendLimit: limit,
			}

			var expiresAtTime time.Time
			if exp != "" {
				expiresAtTime, err = time.Parse(time.RFC3339, exp)
				if err != nil {
					return err
				}
				basic.Expiration = &expiresAtTime
			}

			var grant feegrant.FeeAllowanceI
			grant = &basic

			periodClock, err := cmd.Flags().GetInt64(feegrantcli.FlagPeriod)
			if err != nil {
				return err
			}

			periodLimitVal, err := cmd.Flags().GetString(feegrantcli.FlagPeriodLimit)
			if err != nil {
				return err
			}

			// Check any of period or periodLimit flags set, If set consider it as periodic fee allowance.
			if periodClock > 0 || periodLimitVal != "" {
				periodLimit, err := sdk.ParseCoinsNormalized(periodLimitVal)
				if err != nil {
					return err
				}

				if periodClock > 0 && periodLimit != nil {
					periodReset := getPeriodReset(periodClock)
					if exp != "" && periodReset.Sub(expiresAtTime) > 0 {
						return fmt.Errorf("period(%d) cannot reset after expiration(%v)", periodClock, exp)
					}

					periodic := feegrant.PeriodicAllowance{
						Basic:            basic,
						Period:           getPeriod(periodClock),
						PeriodReset:      getPeriodReset(periodClock),
						PeriodSpendLimit: periodLimit,
						PeriodCanSpend:   periodLimit,
					}

					grant = &periodic

				} else {
					return fmt.Errorf("invalid number of args %d", len(args))
				}
			}

			allowedMsgs, err := cmd.Flags().GetStringSlice(feegrantcli.FlagAllowedMsgs)
			if err != nil {
				return err
			}

			if len(allowedMsgs) > 0 {
				grant, err = feegrant.NewAllowedMsgAllowance(grant, allowedMsgs)
				if err != nil {
					return err
				}
			}

			msg, err := feegrant.NewMsgGrantAllowance(grant, granter, grantee)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().StringSlice(feegrantcli.FlagAllowedMsgs, []string{}, "Set of allowed messages for fee allowance")
	cmd.Flags().String(feegrantcli.FlagExpiration, "", "The RFC 3339 timestamp after which the grant expires for the user")
	cmd.Flags().String(feegrantcli.FlagSpendLimit, "", "Spend limit specifies the max limit can be used, if not mentioned there is no limit")
	cmd.Flags().Int64(feegrantcli.FlagPeriod, 0, "period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset")
	cmd.Flags().String(feegrantcli.FlagPeriodLimit, "", "period limit specifies the maximum number of coins that can be spent in the period")
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}

// NewCmdRevokeFeegrant returns a CLI command handler for creating a MsgRevokeAllowance transaction.
func NewCmdRevokeFeegrant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [grantee] --key-seed granter_key_seed.json",
		Short: "revoke fee-grant",
		Long: strings.TrimSpace(
			fmt.Sprintf(`revoke fee grant from a granter to a grantee. Note, the'--from' flag is
			ignored as it is implied from '--key-seed' flag.

Example:
 $ %s %s revoke shareledger1skj.. --key-seed granter_key_seed.json
			`, version.AppName, feegrant.ModuleName),
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

			grantee, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := feegrant.NewMsgRevokeAllowance(clientCtx.GetFromAddress(), grantee)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(myutils.FlagKeySeed, "", myutils.KeySeedUsage)

	return cmd
}
